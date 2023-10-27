package main

import "fmt"

// implement the queue data structure
type Queue []int

func (s Queue) IsEmpty() bool {
	return len(s) == 0
}

// Push a new value onto the Queue
func (s *Queue) Push(in int) {
	*s = append(*s, in) // Simply append the new value to the end of the Queue
}

// Remove and return top element of Queue. Return false if Queue is empty.
func (s *Queue) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	element := (*s)[0] // Index into the slice and obtain the element.
	*s = (*s)[1:]      // Remove it from the Queue by slicing it off.
	return element, true

}

func main() {
	var n int
	var input []int
	for {
		fmt.Scanf("%d\n", &n)
		if n == 0 {
			break
		}
		input = append(input, n)
	}
	for _, i := range input {
		solution(i)
	}
}

func solution(n int) {
	if n == 1 {
		fmt.Printf("Discarded cards:\nRemaining card: 1\n")
		return
	}
	var q, out Queue
	for i := 1; i <= n; i++ {
		q = append(q, i)
	}
	for {
		val, _ := q.Pop()
		out = append(out, val)
		if len(q) == 1 {
			break
		}
		val, _ = q.Pop()
		q.Push(val)
	}
	fmt.Print("Discarded cards: ")
	for i, num := range out {
		fmt.Print(num)
		// Add a comma after each element except the last one
		if i < len(out)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println()
	fmt.Printf("Remaining card: %d\n", q[0])
}
