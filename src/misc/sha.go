package main

import (
    "crypto/sha1"
    "fmt"
    "crypto/sha256"
)

func main() {
    sha1Sum()

    sha256Sum()
}

func sha256Sum() {
    s := "this is a string"
    hash := sha256.New()
    hash.Write([]byte(s))
    sum := hash.Sum(nil)
    fmt.Println(sum)
    fmt.Printf("%x\n", sum)
}

func sha1Sum() {
    s := "this is a string"
    hash := sha1.New()
    hash.Write([]byte(s))
    sum := hash.Sum(nil)
    fmt.Println(sum)
    fmt.Printf("%x\n", sum)
}
