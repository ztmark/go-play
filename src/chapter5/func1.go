package main

import "fmt"

func main() {
    f(1, 2, 3, "a", "b")
    fmt.Println(f1(1, 2, 3))
    fmt.Println(f2(1, 2, 3))
    fmt.Println(f3(1, 2, 3))
    fmt.Println(f4(1, 2, 3))
    fmt.Println(f5(1, 2, 3))
    fmt.Println(f6(1, 2, 3))
    fmt.Println(f7(1, 2, 3))
    fmt.Println(f8())
    fmt.Println(f9())
    fmt.Println(f10())
    fmt.Println(f11())
}

func f(i, j, k int, s, t string) {
    fmt.Println(i, j, k, s, t)
}

func f1(x int, y int, z int) int {
    return x + y + z
}

func f2(x int, y int, z int) (r int) {
    return x + y + z
}

func f3(x int, y int, z int) (r int) {
    r = x + y + z
    return
}

func f4(x int, y int, _ int) int {
    return x + y
}

func f5(x int, _ int, y int) int {
    return x + y
}

func f6(x int, _ int, _ int) int {
    return x
}

func f7(int, int, int) int {
    return 0
}

func f8() (int, int) {
    return 0,1
}

func f9() (x int, y int) {
    return -1, 0
}

func f10() (x, y int) {
    return -2, -1
}

func f11() (x, y int) {
    x, y = -3, -2
    return
}