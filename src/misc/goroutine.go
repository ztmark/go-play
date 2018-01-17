package main

import (
    "fmt"
    "time"
)

type readOp struct {
    key  string
    resp chan string
}

type writeOp struct {
    key   string
    value string
    resp  chan bool
}

func main() {

    readChannel := make(chan *readOp)
    writeChannel := make(chan *writeOp)

    readCounter := 0
    writeCounter := 0

    go func() {
        data := make(map[string]string)

        for {
            select {
            case r := <-readChannel:
                r.resp <- data[r.key]
                readCounter++
            case w := <-writeChannel:
                data[w.key] = w.value
                w.resp <- true
                writeCounter++
            }
        }

    }()

    for r := 0; r < 50; r++ {
        go func() {
            op := readOp{
                key:  "key" + string(r),
                resp: make(chan string),
            }
            readChannel <- &op
            <-op.resp
        }()
    }

    for w := 0; w < 50; w++ {
        go func() {
            op := writeOp{
                key:   "key" + string(w),
                value: "value" + string(w),
                resp:  make(chan bool),
            }
            writeChannel <- &op
            <-op.resp
        }()
    }

    time.Sleep(time.Second)
    fmt.Println(readCounter)
    fmt.Println(writeCounter)
}
