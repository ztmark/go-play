package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var f float32 = 1 << 24
	fmt.Println(f == f + 1)
	s := "Hello, 世界"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))

	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}

	for i, r := range s {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

	msg := "Hello there 😀😬"
	fmt.Println(len(msg))
	fmt.Println(utf8.RuneCountInString(msg))
	fmt.Println(msg[12:])
	fmt.Println(msg[13:])

	var g float64 = 3 + 0i
	g = 2
	g = 1e123
	g = 'a'
	fmt.Printf("%d,%[1]v,%[1]T", g)

}
