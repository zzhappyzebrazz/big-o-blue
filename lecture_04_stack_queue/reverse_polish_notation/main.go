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
	var n int
	var str string
	fmt.Scanf("%d\n", &n)
	var input []string
	for i := 0; i < n; i++ {
		fmt.Scanf("%s\n", &str)
		input = append(input, str)
	}
	for i := 0; i < n; i++ {
		rpn(input[i])
	}
}

func rpn(input string) {
	var in, out, stack Stack
	in = Stack(input)
	for i := 0; i < len(in); i++ {
		if in[i] >= 'a' && in[i] <= 'z' {
			out.Push(in[i])
		} else if in[i] == ')' {
			char, _ := stack.Pop() // get the operator charactor
			out.Push(char)
			stack.Pop() // pop out the ( charator
		} else {
			stack.Push(in[i])
		}
	}
	fmt.Println(string(out))
}
