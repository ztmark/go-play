package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	reverse(a[:])
	fmt.Println(a)
	reverse(a[:2])
	reverse(a[2:])
	reverse(a)
	fmt.Println(a)
	reverse(a)
	reverse(a[2:])
	reverse(a[:2])
	fmt.Println(a)
	a = append(a, 6)
	a = append(a, 7)
	a = append(a, 8)
	fmt.Println(a, len(a), cap(a))

	data := []string{"abcd", "", "what a day"}
	fmt.Printf("%q\n", data)
	fmt.Printf("%q\n", nonempty(data))

	d := []int{1, 2, 3, 4}
	reverse1(&d)
	fmt.Println(d)
}

func reverse(s []int) {
	for i, j := 0, len(s) - 1; i < j; i, j = i + 1, j - 1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverse1(s *[]int)  {
	for i, j := 0, len(*s) - 1; i < j; i, j = i + 1, j - 1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}