package main

// question:
/*
compare the index number of 2 char arrays
s1 = k = 107
s2 = m = 109
-> sum s1 = 107
-> sum s2 = 109
s1 - s2 = 2 > 1
-> exist
s1[n-1] = s1[n-1] + 1
*/

import "fmt"

func main() {
	var s1, s2 string
	// _ := []rune(s2)
	_, err := fmt.Scanln(&s1)
	if err != nil {
		return
	}
	_, err = fmt.Scanln(&s2)
	if err != nil {
		return
	}

	fmt.Println("No such string")

}
