package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

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

func GoX(w Walkable) {
	w.Walk(10, 0)
}

func Walkers() {
	barsik := Animal{0, 0}
	barsik.Walk(10, 4)
	barsik.Where()

	var walker Walkable
	wallie := Robot{1, 1, "IOSIF"}
	walker = &wallie
	GoX(walker)
	walker.Walk(10, 4)
	wallie.Where()
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

func (r *Robot) String() string {
	return fmt.Sprintf("I AM %s! I MUST DESTROY CAPITALISM!", r.name)
}

type MyException struct {
	code   int
	detail string
}

func (r MyException) Error() string {
	return fmt.Sprintf("ERROR %d, detail: %v", r.code, r.detail)
}

func DestroyCapitalism(r *Robot) (string, error) {
	if r == nil {
		return "", MyException{400, "Need some robot"}
	} else {
		return "DESTROYING IN PROGRESS...", nil
	}
}

func Destroyer() {
	iosif := Robot{1, 1, "IOSIF"}
	fmt.Println(&iosif)
	if v, err := DestroyCapitalism(nil); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(v)
	}

	fmt.Println(DestroyCapitalism(&iosif))
}

type Fib struct {
	prev, next int
}

func (f *Fib) Read(data []byte) (int, error) {
	data[0] = byte(f.prev)
	data[1] = byte(f.prev >> 8)
	data[2] = byte(f.prev >> 16)
	data[3] = byte(f.prev >> 24)

	f.prev, f.next = f.next, f.prev+f.next
	return 4, nil
}

func BytesToInt(data []byte) int {
	res := 0
	res += int(data[0])
	res += int(data[1]) << 8
	res += int(data[2]) << 16
	res += int(data[3]) << 24
	return res
}

func FibRead() {
	r := &Fib{1, 1}

	buff := make([]byte, 4)
	for r.prev < 30 {
		_, ok := r.Read(buff)
		if ok == nil {
			fmt.Println(BytesToInt(buff))
		}
	}
}

type rot13Reader struct {
	r io.Reader
}

func (r *rot13Reader) Read(data []byte) (int, error) {
	n, _ := r.r.Read(data)
	for i := 0; i < n; i++ {
		var offset byte = 0
		if data[i] >= byte('a') && data[i] <= byte('z') {
			offset = byte('a')
		} else if data[i] >= byte('A') && data[i] <= byte('Z') {
			offset = byte('A')
		} else {
			continue
		}

		data[i] = offset + (data[i]-offset+13)%26
	}
	return n, nil
}

func Rot13() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

func main() {
	Walkers()
	TypeAssertion()
	TypeSwitch()
	Destroyer()
	FibRead()
	Rot13()
}
