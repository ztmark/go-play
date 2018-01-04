package main

import (
	"fmt"
	"os"
	"strings"
)

func main()  {
	// fmt.Println(os.Args[0:])

	s, sep := "", ""
	for _, arg := range os.Args[0:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)

    fmt.Println(strings.Join(os.Args[0:], " "))
}
