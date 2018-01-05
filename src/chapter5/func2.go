package main

import (
    "fmt"
    "strings"
)

func square(n int) int {
    return n * n
}

func negative(n int) int {
    return -n
}

func product(n, m int) int {
    return n * m
}

func add1(r rune) rune {
    return r + 1
}

func main() {
    f := square
    fmt.Println(f(3))
    f = negative
    fmt.Println(f(4))
    fmt.Printf("%T\n", f)
    //f = product // error


    fmt.Println(strings.Map(add1, "HAL-9000"))



}
