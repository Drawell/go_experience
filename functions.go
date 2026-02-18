package main

import (
	"fmt"
)

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func fib(prev, next int) (prev_, next_ int) {
	prev_, next_ = next, prev+next
	return
}

func main() {
	fmt.Println("add:", add(1, 2))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(fib(1, 1))
}
