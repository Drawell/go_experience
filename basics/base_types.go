package main

import (
	"fmt"
	"math"
)

func SimpleTypes() {
	var i int
	var f float64 = math.Pi
	const b bool = true
	var s = "there is string"
	fmt.Printf("%v %.3f %v %q\n", i, f, b, s)
}

func TypeCasting() {
	f := 3.14
	i := int(f)
	f = float64(i)
	fmt.Println(i, f)

}

func Pointers() {
	a := 1
	PlusPlus(&a)
	fmt.Println(a)
}

func PlusPlus(ptr *int) {
	*ptr++
}

type Point struct {
	X int
	Y int
}

func AddToPoint(p Point, x, y int) Point {
	p.X += x
	p.Y += y
	return p
}

func MovePoint(p *Point, to Point) {
	p.X += to.X
	p.Y += to.Y
}

func Structs() {
	p1 := Point{1, 2}
	p2 := AddToPoint(p1, 1, 2)
	p3 := p1
	MovePoint(&p3, Point{X: 2})

	fmt.Println(p1, p2, p3)
}

func ArraysAndSlices() {
	var arr [2]string
	arr[0] = "Hello"
	arr[1] = "World"

	for s := 0; s < len(arr); s++ {
		for i := 0; i < len(arr[s]); i++ {
			fmt.Print(string(arr[s][i]))
		}
		fmt.Print(" ")
	}
	fmt.Println()

	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr2)
	arr2[1] = 42

	var slice []int = arr2[1:4]
	fmt.Println(slice)

	slice2 := arr2[0:len(arr2)]
	fmt.Println(slice2)

	phantomArraySlice := []int{5, 4, 3, 2, 1}
	fmt.Println(phantomArraySlice, "len=", len(phantomArraySlice), "cap=", cap(phantomArraySlice))

	phantomArraySlice = phantomArraySlice[1:2]
	fmt.Println(phantomArraySlice, "len=", len(phantomArraySlice), "cap=", cap(phantomArraySlice))

	phantomArraySlice = phantomArraySlice[:4]
	fmt.Println(phantomArraySlice, "len=", len(phantomArraySlice), "cap=", cap(phantomArraySlice))

	dynamic := make([]int, 5)
	fmt.Println(dynamic, "len=", len(dynamic), "cap=", cap(dynamic))

	dynamic2 := make([]int, 1, 5)
	fmt.Println(dynamic2, "len=", len(dynamic2), "cap=", cap(dynamic2))

	dynamic2 = dynamic2[:cap(dynamic2)]
	fmt.Println(dynamic2, "len=", len(dynamic2), "cap=", cap(dynamic2))

	var app []int
	app = append(app, 1)
	fmt.Println(app, "len=", len(app), "cap=", cap(app))

	app = append(app, 2, 3, 4)
	fmt.Println(app, "len=", len(app), "cap=", cap(app))
}

func Maps() {
	figures := map[int]string{
		1: "circle",
		2: "rectangle",
	}
	var root = make(map[string]map[int]string)
	root["figures"] = figures
	root["colors"] = map[int]string{
		255000000: "red",
		000255000: "green",
		000000255: "blue",
	}

	fmt.Println(root)

	for key, value := range root {
		fmt.Println(key)
		for idx, inner := range value {
			fmt.Printf("\t %d: %s\n", idx, inner)
		}
	}

	var value, ok = root["colors"]
	fmt.Println(value, ok)

	delete(root, "colors")

	value, ok = root["colors"]
	fmt.Println(value, ok)
}

func CompareWith(value int) func(int) int {
	return func(other int) int {
		if other > value {
			return 1
		} else if other < value {
			return -1
		} else {
			return 0
		}
	}
}

func Closure() {
	a, b, c := 10, 15, 20
	comparator := CompareWith(15)
	fmt.Println(comparator(a))
	fmt.Println(comparator(b))
	fmt.Println(comparator(c))
}

func main() {
	SimpleTypes()
	TypeCasting()
	Pointers()
	Structs()
	ArraysAndSlices()
	Maps()
	Closure()
}
