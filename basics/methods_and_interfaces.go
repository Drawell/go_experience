package main

import "fmt"

type Walkable interface {
	Walk(toX, toY int)
}

type Animal struct {
	x, y int
}

func (a *Animal) Walk(toX, toY int) {
	a.x += toX
	a.y += toY
}

func (a *Animal) Where() {
	fmt.Printf("I am on x=%d, y=%d\n", a.x, a.y)
}

type Robot struct {
	x, y int
	name string
}

func (r *Robot) Walk(toX, toY int) {
	r.x += toY
	r.y += toX
}

func (r *Robot) Where() {
	fmt.Printf("X=%d, X=%d\n", r.x, r.y)
}

func (r Robot) String() string {
	return fmt.Sprintf("I AM %s! I MUST DESTROY CAPITALISM!", r.name)
}

func GoX(w Walkable) {
	w.Walk(10, 0)
}

func Walkers() {
	barsik := Animal{0, 0}
	barsik.Walk(10, 4)
	barsik.Where()

	var walker Walkable
	iosif := Robot{1, 1, "IOSIF"}
	walker = &iosif
	GoX(walker)
	walker.Walk(10, 4)
	iosif.Where()
	fmt.Println(iosif)
}

func TypeAssertion() {
	var whatIsIt interface{} = "hello"
	fmt.Println(whatIsIt)

	str, ok := whatIsIt.(string)
	fmt.Println(str, ok)

	f, ok := whatIsIt.(float64)
	fmt.Println(f, ok)
}

func ConvertDate(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("convert from timestamp\n")
	case string:
		fmt.Printf("convert from string\n")
	default:
		fmt.Printf("i dont know how to convert %v, %T\n", v, i)
	}
}

func TypeSwitch() {
	ConvertDate(1772755200)
	ConvertDate("06.03.2026")
	ConvertDate(1.0)
}

func main() {
	Walkers()
	TypeAssertion()
	TypeSwitch()
}
