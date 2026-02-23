package main

import (
	"fmt"
	"math/rand"
)

func LoopIf() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	for sum > 0 {
		fmt.Println(sum)
		sum -= 10
	}

	for {
		if i := rand.Int() % 2; i == 0 {
			break
		} else {
			fmt.Println(i)
		}
	}
}

func LoopSlice() {

	arr := []int{1, 2, 3, 4}
	for idx, value := range arr {
		fmt.Println(idx, ":", value)
	}

	for _, value := range arr {
		fmt.Println("only value:", value)
	}

	for idx := range make([]int, 2) {
		fmt.Println("only index:", idx)
	}

}

func SwitchCase(x int) {
	switch {
	case x%2 == 0:
		fmt.Println("Divides by two")
	case x%3 == 0:
		fmt.Println("Divides by three")
	default:
		fmt.Printf("No divides by two or three")
	}
}

func DeferredFunc(x int) {
	fmt.Println("start")

	defer fmt.Println(x / 2)
	defer fmt.Println(x / 3)
	defer fmt.Println(x / 5)
	defer fmt.Println(x / 7)

	fmt.Println("end")
}

func main() {
	LoopIf()
	LoopSlice()
	SwitchCase(6)
	DeferredFunc(11)
}
