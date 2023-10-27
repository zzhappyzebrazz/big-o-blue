package main

import (
	"fmt"
	"reflect"
	"sort"
)

// implement the queue data structure
type Queue []int
type Stack []int

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

// IsEmpty: check if stack is empty
func (s Stack) IsEmpty() bool {
	return len(s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(in int) {
	*s = append(*s, in) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	index := len(*s) - 1   // Get the index of the top most element.
	element := (*s)[index] // Index into the slice and obtain the element.
	*s = (*s)[:index]      // Remove it from the stack by slicing it off.
	return element, true

}

func (s Stack) Peek() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	return s[len(s)-1], true

}
func main() {
	var n, t int

	var input []Queue
	fmt.Scanf("%d\n", &n)
	for n > 0 {
		var q Queue
		for i := 0; i < n; i++ {
			fmt.Scanf("%d", &t)
			q.Push(t)
		}
		input = append(input, q)
		fmt.Scanf("%d\n", &n)
	}
	for _, val := range input {
		solution(val)
	}
}

func solution(in Queue) {
	var compare, out Queue
	var lobby Stack
	compare = append(compare, in...)
	sort.Ints(compare)
	n := len(in)
	for i := 0; i < n; i++ {
		// check the car in the lobby is equal out + 1, this should always checked, until the top of the car stack is no longer satisfied
		for {
			val, _ := lobby.Peek()
			if len(lobby) > 0 && len(out) > 0 && val == out[len(out)-1]+1 {
				out = append(out, val)
				lobby.Pop()
			} else {
				break
			}
		}

		temp, _ := in.Pop()
		if temp == 1 || (len(out) > 0 && temp == out[len(out)-1]+1) {
			out = append(out, temp)
			continue
		}
		lobby = append(lobby, temp)
	}
	for i := len(lobby) - 1; i >= 0; i-- {
		out = append(out, lobby[i])
	}
	ok := reflect.DeepEqual(out, compare)
	if ok {
		fmt.Println("yes")
		return
	}
	fmt.Println("no")
}
