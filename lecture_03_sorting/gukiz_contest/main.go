package main

import (
	"fmt"
	"sort"
)

// 1 3 3 -> 3 1 1
// 1 -> 1
// 3 5 3 4 5 -> 4 1 4 3 1

/*
5 5 4 3 3
1 1 3 4 4
2 3 4 5 6
*/

func main() {
	var n, temp int
	var arr []int
	fmt.Scanf("%d\n", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &temp)
		arr = append(arr, temp)
	}
	arr1 := make([]int, len(arr))
	copy(arr1, arr)
	sort.Ints(arr1)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		arr1[i], arr1[j] = arr1[j], arr1[i]
	}
	myMap := make(map[int]int)
	myMap[arr1[0]] = 1
	count := 2
	for i := 1; i < n; i++ {
		if arr1[i] == arr1[i-1] {
			myMap[arr1[i]] = myMap[arr1[i-1]]
		} else {
			myMap[arr1[i]] = count
		}
		count++
	}
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", myMap[arr[i]])
	}
}
