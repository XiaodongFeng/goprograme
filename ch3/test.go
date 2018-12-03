package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Hello, 世界"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))
	substr := "世界"
	//"Hello, 世界" 中包含字符串 "世界": true
	fmt.Printf("\"%s\" 中包含字符串 \"%s\": %t\n", s, substr, Contains(s, substr))


	//0       H
	//1       e
	//2       l
	//3       l
	//4       o
	//5       ,
	//6
	//7       世
	//10      界
	for i := 0; i < len(s);  {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}


	//0       'H'     72
	//1       'e'     101
	//2       'l'     108
	//3       'l'     108
	//4       'o'     111
	//5       ','     44
	//6       ' '     32
	//7       '世'    19990
	//10      '界'    30028
	for i, r := range s {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

	n:=0
	for _,_ = range s {
		n++
	}
	//9
	fmt.Println(n)

	n=0
	for range s {
		n++
	}
	//9
	fmt.Println(n)

	fmt.Println(string(1234567)) // "�
}

func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

func Contains(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}
