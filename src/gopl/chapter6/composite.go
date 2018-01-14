package main

import (
    "sync"
    "fmt"
)

var cache = struct {
    sync.Mutex
    mapping map[string]string
}{
    mapping: make(map[string]string),
}

func Lookup(key string) string {
    cache.Lock()
    v := cache.mapping[key]
    cache.Unlock()
    return v
}

func main() {

    cache.mapping["ww"] = "www"
    fmt.Println(Lookup("ww"))
    fmt.Println(Lookup("www"))

    lockForCache := cache.Lock
    unlockForCache := cache.Unlock
    lockForCache()
    v := cache.mapping["ww"]
    unlockForCache()
    fmt.Println(v)



}