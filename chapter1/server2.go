package main

import (
	"sync"
	"net/http"
	"log"
	"fmt"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler2)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func counter(writer http.ResponseWriter, request *http.Request) {
	mu.Lock()
	fmt.Fprintf(writer, "Count %d\n", count)
	mu.Unlock()
}

func handler2(writer http.ResponseWriter, request *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(writer, "URL.Path = %q\n", request.URL.Path)
}