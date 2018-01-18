package main

import (
    bs64 "encoding/base64"
    "fmt"
)

func main() {

    s := "this is a string"

    es := bs64.StdEncoding.EncodeToString([]byte(s))
    fmt.Println(es)
    ss, err := bs64.StdEncoding.DecodeString(es)
    if err != nil {
        panic(err)
    }
    fmt.Println(string(ss))

    fmt.Println("======")

    es = bs64.RawStdEncoding.EncodeToString([]byte(s))
    fmt.Println(es)
    ss, err = bs64.RawStdEncoding.DecodeString(es)
    if err != nil {
        panic(err)
    }
    fmt.Println(string(ss))

    fmt.Println("======")

    es = bs64.URLEncoding.EncodeToString([]byte(s))
    fmt.Println(es)
    ss, err = bs64.URLEncoding.DecodeString(es)
    if err != nil {
        panic(err)
    }
    fmt.Println(string(ss))

    fmt.Println("======")

    es = bs64.RawURLEncoding.EncodeToString([]byte(s))
    fmt.Println(es)
    ss, err = bs64.RawURLEncoding.DecodeString(es)
    if err != nil {
        panic(err)
    }
    fmt.Println(string(ss))

}
