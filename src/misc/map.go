package main

import "fmt"

func main() {
    var m = make(map[string]int, 2)
    fmt.Println(m)
    m["a"] = 1
    fmt.Println(m)
    m["b"] = 2
    fmt.Println(m)
    m["c"] = 2
    fmt.Println(m)

    fmt.Println(m["a"])

    v, p := m["b"]
    fmt.Println(v, p)

    v, p = m["d"]
    fmt.Println(v, p)
    fmt.Println(len(m))
}
