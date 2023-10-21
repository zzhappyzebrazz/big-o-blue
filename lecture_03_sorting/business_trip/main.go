package main

import (
	"fmt"
	"sort"
)

/*
5
1 1 1 1 2 2 3 2 2 1 1 1
1 1 1 1 1 1 1 2 2 2 2 3
*/
func main() {
	var n, temp, count int
	var arr []int
	fmt.Scanf("%d\n", &n)
	for i := 0; i < 12; i++ {
		fmt.Scanf("%d", &temp)
		arr = append(arr, temp)
	}
	if n == 0 {
		fmt.Println(0)
		return
	}
	sort.Ints(arr)
	for i := 12 - 1; i >= 0; i-- {
		if n <= 0 {
			fmt.Println(count)
			return
		}
		n = n - arr[i]
		count++
	}
	if n <= 0 {
		fmt.Println(count)
		return
	}
	fmt.Println(-1)
}
