package main

import "fmt"

type ArgError struct {
    arg int
    msg string
}

func (e *ArgError) Error() string {
    return fmt.Sprintf("%s arg=%d\n", e.msg, e.arg)
}

func f1(arg int) (int, error) {
    if arg == 13 {
        return 0, &ArgError{arg:arg, msg:"not accept argument"}
    }
    return arg* 2, nil
}

func main() {
    r, e := f1(13)
    fmt.Println(r, e)
    r, e = f1(11)
    fmt.Println(r, e)
}