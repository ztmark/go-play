package main

import (
	"os"
	"fmt"
	"strings"
)

func main() {
	//var s, sep string
	//for i := 0; i < len(os.Args); i++ {
	//	s += sep + os.Args[i]
	//	sep = " "
	//}
	//fmt.Println(s)

	for _, v := range os.Args[1:] {
		fmt.Println(v)
	}

	fmt.Println(strings.Join(os.Args, " "))
	fmt.Println(os.Args)

}
