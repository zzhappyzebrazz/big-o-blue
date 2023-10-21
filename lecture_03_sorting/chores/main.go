package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, a, b, temp int
	var arr []int
	_, err := fmt.Scanf("%d %d %d\n", &n, &a, &b)
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
	result := arr[b] - arr[b-1]
	if result < 1 {
		fmt.Println(0)
		return
	}
	fmt.Println(result)
}
