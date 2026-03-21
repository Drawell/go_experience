package main

import (
	"fmt"
	calcLib "github.com/Drawell/simple_calc/lib"
)

func main() {
	text := "2 * (1 + 2 * 3)"
	value, err := calcLib.Evaluate(text)
	fmt.Println(value, err)
}
