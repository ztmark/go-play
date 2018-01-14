package main

func main() {

}

type tree struct {
	value int
	left, right *tree
}

func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValue(values[:0], root)
}
func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
		// equivalent
		//return &tree{value: value}
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func appendValue(values []int, t *tree) []int {
	if t != nil {
		values = appendValue(values, t.left)
		values = append(values, t.value)
		values = appendValue(values, t.right)
	}
	return values
}

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func rectangle() {
	var w Wheel
	w.Circle.X = 1
	w.X = 1
	w.Circle.Y = 2
	w.Y = 2
	w.Circle.Radius = 3
	w.Radius = 3
	w.Spokes = 10

	w = Wheel{Circle{Point{1, 2}, 3}, 10}

	w = Wheel{
		Circle: Circle{
			Point: Point{X:1, Y:1},
			Radius:3,
		},
		Spokes:10,
	}
}

