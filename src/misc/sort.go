package main

import (
    "sort"
    "fmt"
)

type ByLength []string

func (s ByLength) Len() int {
    return len(s)
}

func (s ByLength) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func (s ByLength) Less(i, j int) bool {
    return len(s[i]) < len(s[j])
}

func main() {
    s := []string{"a", "what", "are", "funny", "lol"}
    fmt.Println(s)
    sort.Sort(ByLength(s))
    fmt.Println(s)
}