package main

import (
    "bytes"
    "fmt"
)

const shiftBits = 32 << uint(^(uint(0)) >> 63)

type IntSet struct {
    words []int
}

func (s *IntSet) Has(x int) bool {
    word, bit := x / shiftBits, uint(x % shiftBits)
    return word < len(s.words) && s.words[word] & (1 << bit) != 0
}

func (s *IntSet) Add(x int) {
    word, bit := x / shiftBits, uint(x % shiftBits)
    for word >= len(s.words) {
        s.words = append(s.words, 0)
    }
    s.words[word] |= 1 << bit
}

func (s *IntSet) UnionWith(t *IntSet) {
    for i, tword := range t.words {
        if i < len(s.words) {
            s.words[i] |= tword
        } else {
            s.words = append(s.words, tword)
        }
    }
}

func (s *IntSet) String() string {
    var buf bytes.Buffer
    buf.WriteByte('{')
    for i, word := range s.words {
        if word == 0 {
            continue
        }
        for j := 0; j < shiftBits; j++ {
            if word & (1 << uint(j)) != 0 {
                if buf.Len() > len("{") {
                    buf.WriteByte(' ')
                }
                fmt.Fprintf(&buf, "%d", shiftBits * i + j)
            }
        }
    }
    buf.WriteByte('}')
    return buf.String()
}

func (s *IntSet) AddAll(ints ...int) {
    for _, n := range ints {
        s.Add(n)
    }
}

func (s *IntSet) IntersectWith(set *IntSet) {
    for i, _ := range s.words {
        if i < len(set.words) {
            s.words[i] &= set.words[i]
        } else {
            s.words[i] &= 0
        }
    }

}

func (s *IntSet) DifferenceWith(set *IntSet) {
    for i, w := range set.words {
        if i < len(s.words) {
            s.words[i] &= ^w
        }
    }
}

func (s *IntSet) SymmetricDifference(set *IntSet) {
    for i, w := range set.words {
        if i < len(s.words) {
            s.words[i] = (s.words[i] | w) &^ (s.words[i] & w)
        }
    }
}

func (s *IntSet) Len() int {
    l := 0
    for _, w := range s.words {
        for i := 1; w > 0; w >>= 1 {
            if w & i == 1 {
                l++
            }
        }
    }
    return l
}

func (s *IntSet) Remove(x int) {
    w, b := x / shiftBits, x % shiftBits
    s.words[w] &^= 1 << uint(b)
}

func (s *IntSet) Clear() {
    for i, _ := range s.words {
        s.words[i] &= 0
    }
}

func (s *IntSet) Copy() *IntSet {
    r := IntSet{}
    for _, w := range s.words {
        r.words = append(r.words, w)
    }
    return &r
}

func (s *IntSet) Elems() []int {
    return nil
}

func main() {
    var x, y IntSet
    x.Add(1)
    x.Add(144)
    x.Add(9)
    fmt.Println(x.String())

    y.Add(9)
    y.Add(43)
    fmt.Println(y.String())

    x1 := x.Copy()
    x1.IntersectWith(&y)
    fmt.Println(x1)
    x1 = x.Copy()
    x1.DifferenceWith(&y)
    fmt.Println(x1)
    x.SymmetricDifference(&y)
    fmt.Println(x.String())
    fmt.Println(x.Len())
    x.Remove(43)
    fmt.Println(x.String())
    x.Clear()
    fmt.Println(x.String())

    fmt.Println("====")
    fmt.Println(shiftBits)
    fmt.Println(32 << (uint(^(uint(0)) >> 63)))
}