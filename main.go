package main

import (
	"fmt"
	"github.com/Drawell/simple_calc/calc"
)

func main() {
	text := "2 * (1 + 2 * 3)"
	value, err := calc.Evaluate(text)
	fmt.Println(value, err)
}
