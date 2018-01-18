package main

import (
    "net/url"
    "fmt"
    "net"
)

func main() {
    s := "postgres://user:pass@host.com:5432/path?k=v#f"

    u, e := url.Parse(s)
    if e != nil {
        panic(e)
    }

    fmt.Println(u.Scheme)

    fmt.Println(u.User)
    fmt.Println(u.User.Username())
    fmt.Println(u.User.Password())

    fmt.Println(u.Host)
    fmt.Println(net.SplitHostPort(u.Host))

    fmt.Println(u.Path)
    fmt.Println(u.Fragment)
    fmt.Println(u.Query())
    fmt.Println(u.RawPath)
    fmt.Println(u.RawQuery)
    fmt.Println(url.ParseQuery(u.RawQuery))

}
