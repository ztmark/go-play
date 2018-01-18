package main

import (
    "fmt"
    "strings"
)

func main() {
    var strs = []string{"peach", "apple", "pear", "plum"}
    fmt.Println(Index(strs, "pear"))
    fmt.Println(Include(strs, "grape"))
    fmt.Println(Any(strs, func(v string) bool {
        return strings.HasPrefix(v, "p")
    }))
    fmt.Println(All(strs, func(v string) bool {
        return strings.HasPrefix(v, "p")
    }))
    fmt.Println(Filter(strs, func(v string) bool {
        return strings.Contains(v, "e")
    }))

    fmt.Println(Map(strs, strings.ToUpper))
}


func Index(ss []string, s string) int {
    for index, value := range ss {
        if value == s {
            return index
        }
    }
    return -1
}

func Include(ss []string, s string) bool {
    return Index(ss, s) != -1
}

func Any(ss []string, f func(string) bool) bool {
    for _, s := range ss {
        if f(s) {
            return true
        }
    }
    return false
}

func All(ss []string, f func(string) bool) bool {
    for _, s := range ss {
        if !f(s) {
            return false
        }
    }
    return true
}

func Filter(ss []string, f func(string) bool) []string {
    var res []string
    for _, s := range ss {
        if f(s) {
            res = append(res, s)
        }
    }
    return res
}

func Map(ss []string, f func(string) string) []string {
    var res []string
    for _, s := range ss {
        res = append(res, f(s))
    }
    return res
}
