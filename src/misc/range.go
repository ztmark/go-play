package main

import "fmt"

func main() {
    nums := []int{1,2,3,4,5}
    for i, n := range nums {
        fmt.Println(i, n)
    }

    kvs := map[string]int{"a":1, "b":2, "c":4, "d":5}
    for k, v := range kvs {
        fmt.Println(k, v)
    }
    for k := range kvs {
        fmt.Println(k)
        fmt.Println(kvs[k])
    }

    for i, c := range "what" {
        fmt.Println(i, c)
        fmt.Printf("%c\n", c)
    }
}
