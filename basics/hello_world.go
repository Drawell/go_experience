package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Hello, world!")

	var name string
	count, err := fmt.Scanln(&name)
	fmt.Println("Hello, ", name, "!")
	log.Println(count, err)
}
