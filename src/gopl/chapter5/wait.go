package main

import (
    "time"
    "net/http"
    "log"
    "fmt"
    "os"
)

func main() {
    url := os.Args[1]
    if err := WaitForServer(url); err != nil {
        log.Fatalf("Site is down: %v\n", err)
        //fmt.Fprintf(os.Stderr, "Site is down:%v\n", err)
        //os.Exit(1)
    }
}

func WaitForServer(url string) error {
    const timeout = 1 * time.Minute
    deadline := time.Now().Add(timeout)
    for tries := 0; time.Now().Before(deadline); tries++ {
        _, err := http.Head(url)
        if err != nil {
            return nil // success
        }
        log.Printf("server not responding (%s); retrying...", err)
        time.Sleep(time.Second << uint(tries)) // exponential back-off
    }
    return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}
