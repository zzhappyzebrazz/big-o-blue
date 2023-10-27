package main

import "fmt"

// using reference method to modify the object, not copy it out
type Stack []byte

// IsEmpty: check if stack is empty
func (s Stack) IsEmpty() bool {
	return len(s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(char byte) {
	*s = append(*s, char) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (byte, bool) {
	if s.IsEmpty() {
		return 0, false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

func main() {
	var n, i int
	var str string
	var input []string
	fmt.Scanf("%d\n", &n)
	for i < n {
		fmt.Scanf("%s\n", &str)
		input = append(input, str)
		i++
	}
	for _, in := range input {
		solution(in)
	}
}

func solution(in string) {
	chars := []byte(in)
	if chars[0] == '>' {
		fmt.Println(0)
		return
	}
	var s Stack
	var result int
	for _, c := range chars {
		if c == '<' {
			s.Push(c)
		}
		if c == '>' && len(s) > 0 {
			s.Pop()
			result += 2
		}
		
	}
	if len(s) > 0 {
		fmt.Println(0)
		return
	}
	fmt.Println(result)
}
