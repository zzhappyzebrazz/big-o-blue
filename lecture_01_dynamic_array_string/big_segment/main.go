package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, x, y int
	input := make(map[int][2]int)
	var arrX, arrY []int
	fmt.Scanf("%d\n", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d", &x, &y)
		arr := [2]int{x, y}
		arrX = append(arrX, x)
		arrY = append(arrY, y)
		input[i] = arr
	}
	sort.Ints(arrX)
	sort.Ints(arrY)
	for i := 0; i < n; i++ {
		val := input[i]
		if val[0] == arrX[0] && val[1] == arrY[n-1] {
			fmt.Println(i + 1)
			return
		}
	}
	fmt.Println(-1)
}
