package main

import "fmt"

func main() {
    var arr [20]int
    fmt.Println(arr)

    arr[2] = 12
    fmt.Println(arr)

    arr1 := [10]int{}
    fmt.Println(arr1)
    arr1[1] = 23
    fmt.Println(arr1)

    arr2 := [5][3]int{}
    fmt.Println(arr2)
    arr2[3][2] = 32
    fmt.Println(arr2)
    for i := 0; i < 5; i++ {
        for j := 0; j < 3; j++ {
            arr2[i][j] = i + j
        }
    }
    fmt.Println(arr2)

}
