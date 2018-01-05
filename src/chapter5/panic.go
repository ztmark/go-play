package main

import "fmt"

func main() {
    fmt.Println(returnValueWithoutReturn())
}

func returnValueWithoutReturn() (r int) {
    defer func() {
        switch p := recover(); p {
        case nil:
        case "something wrong":
            r = 1
        default:
            panic(p)
        }
    }()
    panic("something wrong")
}
