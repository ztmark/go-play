package main

import (
    "log"
    "net/http"
    "fmt"
)

type dollars float32

type database map[string]dollars

func (d dollars) String() string {
    return fmt.Sprintf("$%.2f", d)
}

func (db database) list(w http.ResponseWriter, r *http.Request) {
    for item, price := range db {
        fmt.Fprintf(w, "%s: %s\n", item, price)
    }
}

func (db database) price(w http.ResponseWriter, r *http.Request) {
    item := r.URL.Query().Get("item")
    price, ok := db[item]
    if !ok {
        w.WriteHeader(http.StatusNotFound)
        fmt.Fprintf(w, "no such item: %q\n", item)
        return
    }
    fmt.Fprintf(w, "%s\n", price)
}

//func (db database) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//
//    switch r.URL.Path {
//    case "/list":
//        for item, price := range db {
//            fmt.Fprintf(w, "%s: %s\n", item, price)
//        }
//    case "/price":
//        item := r.URL.Query().Get("item")
//        price, ok := db[item]
//        if !ok {
//            w.WriteHeader(http.StatusNotFound)
//            fmt.Fprintf(w, "no such item: %q\n", item)
//            return
//        }
//        fmt.Fprintf(w, "%s\n", price)
//    default:
//        w.WriteHeader(http.StatusNotFound)
//        fmt.Fprintf(w, "no such page: %s\n", r.URL)
//    }
//
//
//}

func main() {
    db := database{"shoes": 50, "socks": 5}

    mux := http.NewServeMux()
    mux.Handle("/list", http.HandlerFunc(db.list))
    mux.Handle("/price", http.HandlerFunc(db.price))

    // use default server mux
    go func() {
        http.HandleFunc("/list", db.list)
        http.HandleFunc("/price", db.price)
        log.Fatal(http.ListenAndServe("localhost:8680", nil))
    }()

    log.Fatal(http.ListenAndServe("localhost:8080", mux))


}
