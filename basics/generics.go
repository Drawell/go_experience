package main

import "fmt"

type StopIterationError string

func (e StopIterationError) Error() string {
	return fmt.Sprintf("Iteration stopd: %v", string(e))
}

type List[T comparable] struct {
	next  *List[T]
	value T
}

func (l *List[T]) Get(index int) (T, error) {
	current := 0
	for current < index {
		if l.next == nil {
			return *new(T), StopIterationError("No such index")
		}
		l = l.next
		current += 1
	}

	return l.value, nil

}

func main() {
	l3 := &List[int]{nil, 3}
	l2 := &List[int]{l3, 2}
	l1 := &List[int]{l2, 1}
	l0 := &List[int]{l1, 0}

	value, ok := l0.Get(0)
	fmt.Printf("Value: %v, ok: %v\n", value, ok)

	value, ok = l0.Get(3)
	fmt.Printf("Value: %v, ok: %v\n", value, ok)

	value, ok = l0.Get(5)
	fmt.Printf("Value: %v, ok: %v\n", value, ok)
}
