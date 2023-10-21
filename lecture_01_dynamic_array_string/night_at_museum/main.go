package main

import (
	"fmt"
	"strings"
)

// question:
/*
the alphabet string circle start from a, all letter is lowercase
a b c d ..... x y z -> a ... -> 26 rings
the ring can move clockwise and anti-clockwise, each jump from a letter is counted as 1 step
count how many steps it take to fully construct the input string
example
ares
start at a
offset a = 97, z = 122
input[0] = a -> a = 0
input[1] = a -> r, clockwise = |a - r| = 17, anti-clockwise = 26 - r - a  = 9 -> input[1] = 9
input[2] = r -> e, clockwise = |r - e| = 13, anti-clockwise = 26 - e - r = 13 -> input[2] = 13
input[3] = e -> s, clockwise = |e - s| = 14, anti-clockwise = 26 - 14 = 12 -> input[3] = 12
result = 0 + 9 + 13 + 12 = 34
*/
func main() {
	var input string
	var result int
	_, err := fmt.Scanln(&input)
	if err != nil {
		panic(err)
	}
	input = strings.ToLower(input)
	for i := 0; i < len(input); i++ {
		if i == 0 {
			result += findStep(97, int(input[i]))
		} else {
			result += findStep(int(input[i-1]), int(input[i]))
		}
	}

	fmt.Println(result)
}

func abs(input int) int {
	if input < 0 {
		return -input
	}
	return input
}

func findStep(current, next int) int {
	clockwise := abs(current - next)
	anti_clockwise := 26 - clockwise
	if clockwise < anti_clockwise {
		return clockwise
	}
	return anti_clockwise
}
