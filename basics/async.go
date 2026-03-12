package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Add(a, b int, out chan int) {
	time.Sleep(500 * time.Millisecond)
	out <- a + b
}

func SimpleChan() {
	out := make(chan int)
	a := []int{1, 2, 3, 4, 5}
	b := []int{10, 20, 30, 40, 50}
	for i := 0; i < len(a); i++ {
		go Add(a[i], b[i], out)
		v := <-out
		fmt.Printf("Value: %v\n", v)
	}
}

func BuffChan() {
	out := make(chan int, 5)
	a := []int{1, 2, 3, 4, 5}
	b := []int{10, 20, 30, 40, 50}
	for i := 0; i < len(a); i++ {
		go Add(a[i], b[i], out)
	}
	for i := 0; i < len(a); i++ {
		v := <-out
		fmt.Printf("Value: %v\n", v)
	}
}

type Pair[T any] struct {
	a, b T
}

func Zip[T any](a []T, b []T, out chan Pair[T]) {
	for i := 0; i < min(len(a), len(b)); i++ {
		out <- Pair[T]{a[i], b[i]}
	}
	close(out)
}

func ZipTest() {
	a := []int{1, 2, 3, 4, 5, 6}
	b := []int{10, 20, 30, 40, 50}
	c := make(chan Pair[int])

	go Zip(a, b, c)
	for p := range c {
		fmt.Println(p)
	}

}

type Person struct {
	answers map[string]string
}

func Boring(out chan string) {
	time.Sleep(2000 * time.Millisecond)
	out <- "Its get bored"
}

func Chat(p Person, out chan string, input chan string) {
	boring := make(chan string)
	go Boring(boring)
	for {
		select {
		case inputPhrase := <-input:
			if inputPhrase == "I am out" {
				out <- "bye..."
				close(out)
				return
			}
			output, ok := p.answers[inputPhrase]
			if ok == true {
				out <- output
			} else {
				out <- "what?"
			}
		case b := <-boring:
			out <- b
		default:
			time.Sleep(300 * time.Millisecond)
			out <- "..."
		}
	}
}

func Ask(out chan string) {
	phrases := []string{"Hello", "How are you", "What are u doing", "Go to walk", "I am out"}
	for {
		time.Sleep(1000 * time.Millisecond)
		idx := rand.Int()
		p := phrases[idx%len(phrases)]
		fmt.Println("*", p)
		out <- p
	}

}

func ChatTest() {
	m := map[string]string{"Hello": "hi", "How are you": "ok", "What are u doing": "nothing"}
	p := Person{m}

	asks := make(chan string)
	answers := make(chan string)
	go Chat(p, answers, asks)
	go Ask(asks)
	for answer := range answers {
		fmt.Println("-", answer)
	}

}

func main() {
	//SimpleChan()
	//BuffChan()
	//ZipTest()
	ChatTest()
}
