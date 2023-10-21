package main

import (
	"fmt"
	"sort"
)

/*
given the set of 8 points
check if all the point is construct of x1 < x2 < x3 and y1 < y2 < y3
if there is a point of x2, y2 -> ugly
else respectable

*/

func main() {
	var arr [8][2]int
	for i := 0; i < 8; i++ {
		fmt.Scanf("%d %d\n", &arr[i][0], &arr[i][1])
	}
	// check all the x and y
	x, y := make(map[int]int), make(map[int]int)
	for _, val := range arr {
		x[val[0]]++
		y[val[1]]++
	}
	if len(x) != 3 || len(y) != 3 {
		fmt.Println("ugly")
		return
	}
	sort.Ints(x)
	sort.Ints(y)
	for _, val := range arr {
		if val[0] == x[1] && val[1] == y[1] {
			fmt.Println("ugly")
			return
		}
	}

	fmt.Println("respectable")
}
