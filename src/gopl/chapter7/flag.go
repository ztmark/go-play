package main

import (
    "fmt"
    "flag"
)

type Length struct {
    value float64
    unit string
}

func (l *Length) String() string {
    return fmt.Sprintf("%f%s", l.value, l.unit)
}

type lengthFlag struct {
    Length
}

func (l *lengthFlag) Set(s string) error {
    var unit string
    var value float64
    fmt.Sscanf(s, "%f%s", &value, &unit)

    switch unit {
    case "m", "M":
        l.value = value
        l.unit = "meter"
        return nil
    case "cm", "CM":
        l.value = value
        l.unit = "centimeter"
        return nil
    }
    return fmt.Errorf("invalid length %q", s)
}

func LengthFlag(name string, value Length, usage string) *Length {
    f := lengthFlag{value}
    flag.CommandLine.Var(&f, name, usage)
    return &f.Length
}

var length = LengthFlag("length", Length{0.0, "m"}, "the length")

func main() {
    flag.Parse()
    fmt.Println(*length)
}