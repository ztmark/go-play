package main

import (
    "math"
    "fmt"
    "image/color"
)

type Point struct {
    X, Y float64
}

func Distance(p, q Point) float64 {
    return math.Hypot(q.X - p.X, q.Y - p.Y)
}

func (p Point) Distance(q Point) float64 {
    return math.Hypot(p.X - q.X, p.Y - q.Y)
}

func (p *Point) ScaleBy(factor float64) {
    p.X *= factor
    p.Y *= factor
}

type Path []Point

func (path Path) Distance() float64 {
    sum := 0.0
    for i := range path {
        if i > 0 {
            sum += path[i].Distance(path[i - 1])
        }
    }
    return sum
}

func add1(p Point) Point {
    p.X += 1
    p.Y += 1
    return p
}

type ColorPoint struct {
    Point
    Path
    color color.RGBA
}

func main() {
    p := Point{1, 3}
    q := Point{3, 8}
    fmt.Println(Distance(p, q))
    fmt.Println(p.Distance(q))

    p.ScaleBy(2)
    fmt.Println(p)

    ptr := &p

    ptr.ScaleBy(2)
    fmt.Println(p)

    fmt.Println(ptr.Distance(q))

    fmt.Println(p)
    p1 := add1(p)
    fmt.Println(p)
    fmt.Println(p1)

    fmt.Println("===================")
    perim := Path{
        {1, 1},
        {5, 1},
        {5, 4},
        {1, 1},
    }
    fmt.Println(perim.Distance())


    fmt.Println("===================")
    red := color.RGBA{255, 0, 0, 255}
    blue := color.RGBA{0, 0, 255, 255}
    cp := ColorPoint{Point{1, 2}, Path{}, red}
    cq := ColorPoint{Point{4, 5}, Path{}, blue}
    fmt.Println(cp.Point.Distance(cq.Point))
    p.ScaleBy(2)
    q.ScaleBy(2)
    fmt.Println(cp.Point.Distance(cq.Point))

}
