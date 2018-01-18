package main

import (
    "bufio"
    "os"
    "fmt"
    "strings"
)

func main() {

    scanner := bufio.NewScanner(os.Stdin)

    for scanner.Scan() {
        text := scanner.Text()
        fmt.Println(strings.ToUpper(text))
    }

    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "error: ", err)
        os.Exit(1)
    }

}
