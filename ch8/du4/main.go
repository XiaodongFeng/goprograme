package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var done = make(chan struct{})

var verbose = flag.Bool("v", false, "show verbose progress messages")
var sema = make(chan struct{}, 20)

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	go func() {
		a := make([]byte, 1)
		os.Stdin.Read(a)
		fmt.Println(a)
		close(done)
	}()

	fileSizes := make(chan int64)
	var n sync.WaitGroup
	go func() {
		for _, root := range roots {
			n.Add(1)
			go walkDir(&n, root, fileSizes)
		}
	}()

	go func() {
		n.Wait()
		close(fileSizes)
	}()

	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

	var nfiles, nbytes int64
loop:
	for {
		select {
		case <-done:
			// Drain fileSizes to allow existing goroutines to finish.
			for range fileSizes {
				// Do nothing.
			}
			panic(1)
			return
		case size, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes was closed
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}
	printDiskUsage(nfiles, nbytes) // final totals
}

func walkDir(n *sync.WaitGroup, dir string, fileSizes chan<- int64) {
	defer n.Done()
	if cancelled() {
		return
	}
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			n.Add(1)
			walkDir(n, subdir, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func dirents(dir string) []os.FileInfo {
	select {
	case sema <- struct{}{}: // acquire token
	case <-done:
		return nil // cancelled
	}
	defer func() { <-sema }() // release token

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}
