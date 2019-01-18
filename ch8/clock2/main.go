package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listrener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listrener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		// handle
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
