package main

import "fmt"

type pair struct {
	char string
	val  int
}

type Stack []pair

func (s Stack) Peak() string {
	if len(s) == 0 {
		return ""
	}
	return s[len(s)-1].char
}

func (s *Stack) Pop() (pair, bool) {
	if len(*s) == 0 {
		return pair{}, false
	}
	index := len(*s) - 1
	ret := (*s)[index]
	*s = (*s)[:index-1]
	return ret, true
}

func main() {
	var input string
	fmt.Scanf("%s", &input)

}
