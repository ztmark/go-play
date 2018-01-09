package main

import (
    "fmt"
    "time"
    "os"
)

func main() {
    //launch()
    //printlnNumbers()
    printCountdown()

}
func printCountdown() {
    abort := make(chan struct{})
    go func() {
        os.Stdin.Read(make([]byte, 1))
        abort <- struct{}{}
    }()
    fmt.Println("Commencing countdown. Press return to abort.")
    tick := time.Tick(1 * time.Second)
    for countdown := 10; countdown > 0; countdown-- {
        fmt.Println(countdown)
        select {
        case <-tick:
            // Do nothing
        case <-abort:
            fmt.Println("Launch aborted!")
            return
        }
    }
    fmt.Println("Launched!")
}

func launch() {
    abort := make(chan struct{})
    go func() {
        os.Stdin.Read(make([]byte, 1))
        abort <- struct{}{}
    }()
    fmt.Println("Commencing countdown. Press return to abort.")
    select {
    case <-time.After(10 * time.Second):
        // Do nothing.
        fmt.Println("Launched!")
    case <-abort:
        fmt.Println("Launch aborted!")
        return
    }
    fmt.Println("Launched!")
}

func printlnNumbers() {
    ch := make(chan int, 2)
    for i := 0; i < 10; i++ {
        select {
        case x := <-ch:
            fmt.Println(x)
        case ch <- i:
        }
    }
}
