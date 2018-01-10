package memotest

import (
    "net/http"
    "io/ioutil"
    "time"
    "log"
    "fmt"
    "testing"
    "sync"
)

func httpGetBody(url string) (interface{}, error) {
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    return ioutil.ReadAll(resp.Body)
}

var HTTPGetBody = httpGetBody

func incomingURLs() <-chan string {
    ch := make(chan string)
    go func() {
        for _, url := range []string{
            "https://www.bilibili.com/",
            "https://v.qq.com/",
            "https://www.v2ex.com/",
            "http://gopl.io",
            "https://www.bilibili.com/",
            "https://v.qq.com/",
            "https://www.v2ex.com/",
            "http://gopl.io",
        } {
            ch <- url
        }
        close(ch)
    }()
    return ch
}

type M interface {
    Get(key string) (interface{}, error)
}

func Sequential(t *testing.T, m M) {
    //!+seq
    for url := range incomingURLs() {
        start := time.Now()
        value, err := m.Get(url)
        if err != nil {
            log.Print(err)
            continue
        }
        fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
    }
    //!-seq
}

func Concurrent(t *testing.T, m M) {
    var n sync.WaitGroup
    for url := range incomingURLs() {
        n.Add(1)
        go func(url string) {
            defer n.Done()
            start := time.Now()
            value, err := m.Get(url)
            if err != nil {
                log.Print(err)
            }
            fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
        }(url)
    }
    n.Wait()
}

