package main

import (
    "fmt"
)

func main() {
    var arr [20]int
    fmt.Println(arr)

    arr[2] = 12
    fmt.Println(arr)

    arr1 := [10]int{}
    fmt.Println(arr1)
    arr1[1] = 23
    fmt.Println(arr1)

    fmt.Println("===========")
    arr1Copy := plus1(arr1)
    fmt.Println(arr1)
    fmt.Println(arr1Copy)
    fmt.Println("===========")


    fmt.Println("===========")
    arr1Copy1 := plus2(&arr1)
    fmt.Println(arr1)
    fmt.Println(arr1Copy1)
    fmt.Println("===========")

    arr3 := [...]int{1,2,3,4,5,6}
    fmt.Println(len(arr3))
    fmt.Println(arr3)
    fmt.Printf("%T\n", arr3)


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

func plus1(arr [10]int) [10]int {
    for i := 0; i < 10; i++ {
        arr[i] += 1
    }
    return arr
}

func plus2(arr *[10]int) [10]int {
    for i := 0; i < 10; i++ {
        (*arr)[i] += 2
    }
    return *arr
}