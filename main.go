package main

import (
	"fmt"
	"go_experience/calculator"
)

func main() {
	text := "2 * (1 + 2 * 3)"
	value, err := calculator.Evaluate(text)
	fmt.Println(value, err)
}
