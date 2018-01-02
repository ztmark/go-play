package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	var m map[string]int
	fmt.Printf("%T\n", m)
	fmt.Println(m)
	fmt.Println(m == nil)
	fmt.Println(m["a"])

	if age, ok := m["a"]; !ok {
		fmt.Println("not exist")
	} else {
		fmt.Println(age)
	}

	fmt.Println(mapEqual(map[string]int{"a": 1}, map[string]int{"b": 1}))
	fmt.Println(mapEqual(map[string]int{"a": 1}, map[string]int{"a": 1}))


	fmt.Println(k([]string{"a", "b", "12", "cd"}))
}

func mapEqual(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}

func dup() {
	seen := make(map[string]bool)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}
}

var counter = make(map[string]int)

func k(list []string) string {
	return fmt.Sprintf("%q", list)
}

func Add(list []string) {
	counter[k(list)]++
}

func Count(list []string) int {
	return counter[k(list)]
}

var graph = make(map[string]map[string]bool)

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}