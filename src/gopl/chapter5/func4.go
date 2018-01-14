package main

import (
    "runtime"
    "os"
    "fmt"
)

func main() {
    defer printStack()
    go func() {
        fmt.Println("in a goroutine")
    }()
    fmt.Println("in main")
}

func printStack() {
    var buf [4096]byte
    n := runtime.Stack(buf[:], true)
    os.Stdout.Write(buf[:n])
}
