package main

import (
    "fmt"
    "sync/atomic"
    "time"
)

func main() {


    done := make(chan struct{})
    var counter uint64 = 0

    for i := 0; i < 1000; i++ {
        go func() {
            //counter++
            atomic.AddUint64(&counter, 1)
            done<- struct{}{}
        }()
    }


    time.Sleep(time.Microsecond)
    fmt.Println(atomic.LoadUint64(&counter))

    for i := 0; i < 1000; i++ {
        <-done
    }

    fmt.Println(counter)


}
