package main

import (
	"fmt"
)

func pop(s *[]int) int {
	i := len(*s) - 1
	ret := (*s)[i]
	*s = (*s)[:i]
	return ret
}

func main() {
	var input string
	var s []int
	fmt.Scanf("%s", &input)
	for _, c := range input {
		if c == 'O' {
			s = append(s, 16)
		} else if c == 'C' {
			s = append(s, 12)
		} else if c == 'H' {
			s = append(s, 1)
		} else if c == '(' {
			s = append(s, -1)
		} else if c >= '1' && c <= '9' {
			s[len(s)-1] = s[len(s)-1] * (int(c) - 48)
		} else {
			count := 0
			for {
				if s[len(s)-1] == -1 {
					break
				}
				count += pop(&s)
			}
			pop(&s) // pop out the -1 value
			s = append(s, count)
		}
	}
	answer := 0
	for _, val := range s {
		answer += val
	}
	fmt.Println(answer)
}
