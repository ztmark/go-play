package main

import (
    "fmt"
    "time"
    "log"
)


func main() {

    defer trace("main")()

    var fs []func()
    for i := 0; i < 5; i++ {
        i := i // fix the problem
       fs = append(fs, func() {
           time.Sleep(1 * time.Second)
           fmt.Println(i)
       })
    }

    for _, f := range fs {
        f()
    }


}

func trace(msg string) func() {
    start := time.Now()
    log.Printf("enter %s", msg)
    return func() {
        log.Printf("exit %s (%s)", msg, time.Since(start))
    }
}