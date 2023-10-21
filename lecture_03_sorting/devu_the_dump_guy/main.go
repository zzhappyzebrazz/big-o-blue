package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, x, temp, result int
	var arr []int
	_, err := fmt.Scanf("%d %d\n", &n, &x)
	if err != nil {
		panic(err)
	}
	for i := 0; i < n; i++ {
		_, err := fmt.Scanf("%d", &temp)
		if err != nil {
			panic(err)
		}
		arr = append(arr, temp)
	}
	sort.Ints(arr)
	for i := 0; i < len(arr); i++ {
		result += arr[i] * x
		if x > 1 {
			x--
		}
	}
	fmt.Println(result)
}
