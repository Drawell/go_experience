package main

import (
	"fmt"
	"math"
)

func simpleTypes() {
	var i int
	var f float64 = math.Pi
	const b bool = true
	var s = "there is string"
	fmt.Printf("%v %.3f %v %q\n", i, f, b, s)
}

func typeCasting() {
	f := 3.14
	i := int(f)
	f = float64(i)
	fmt.Println(i, f)

}

func main() {
	simpleTypes()
	typeCasting()
}
