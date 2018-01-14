package main

import "fmt"

func main() {
    var s []int
    fmt.Println(s)
    fmt.Println("len :", len(s))
    fmt.Println("cap :", cap(s))
    //s[0] = 23 // panic index out of range

    fmt.Println("===============")

    s = make([]int, 3)
    fmt.Println(s)
    fmt.Println("len :", len(s))
    fmt.Println("cap :", cap(s))

    fmt.Println("===============")
    //s[3] = 3 // panic index out of range
    //fmt.Println(s)
    //fmt.Println("len :", len(s))
    //fmt.Println("cap :", cap(s))

    fmt.Println("===============")
    s1 := append(s, 23, 34)
    fmt.Printf("%T\n", s1)
    fmt.Println(s1)
    fmt.Println("len :", len(s1))
    fmt.Println("cap :", cap(s1))


    fmt.Println("===============")
    s2 := []int{1, 2, 3, 4, 5}
    fmt.Println(s2)
    fmt.Printf("%T\n", s2)
    fmt.Println("len :", len(s2))
    fmt.Println("cap :", cap(s2))

    fmt.Println("===============")
    s3 := s2[1:]
    fmt.Println(s3)
    fmt.Printf("%T\n", s3)
    fmt.Println("len :", len(s3))
    fmt.Println("cap :", cap(s3))
    s3[2] = 2222
    fmt.Println(s2)
    fmt.Println(s3)

    fmt.Println("===============")
    a := [10]int{1,2,3,4,5,6,7}
    s4 := a[3:7]
    fmt.Println(a)
    fmt.Printf("%T\n", a)
    fmt.Println("len :", len(a))
    fmt.Println("cap :", cap(a))
    fmt.Println(s4)
    fmt.Printf("%T\n", s4)
    fmt.Println("len :", len(s4)) // 4
    fmt.Println("cap :", cap(s4)) // 7
    //s4[6] = 16 // panic index out of range

    s4 = s4[:cap(s4)]
    s4[6] = 16
    fmt.Println(s4)
    fmt.Println("len :", len(s4))
    fmt.Println("cap :", cap(s4))

    s4 = append(s4, 14)
    s4 = append(s4, 15)
    s4 = append(s4, 16)
    s4 = append(s4, 17)
    fmt.Println(a)
    fmt.Println("len :", len(a))
    fmt.Println("cap :", cap(a))
    fmt.Println(s4)
    fmt.Println("len :", len(s4))
    fmt.Println("cap :", cap(s4))

    s4[1] = 100
    fmt.Println(a)
    fmt.Println(s4)



}
