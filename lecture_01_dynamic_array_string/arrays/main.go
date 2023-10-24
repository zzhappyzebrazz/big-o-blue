package main

import (
	"fmt"
)

func main() {
	var nA, nB, k, m, t int
	fmt.Scanf("%d %d\n", &nA, &nB)
	fmt.Scanf("%d %d\n", &k, &m)
	var arrA, arrB []int
	for i := 0; i < nA; i++ {
		fmt.Scanf("%d", &t)
		arrA = append(arrA, t)
	}
	for i := 0; i < nB; i++ {
		fmt.Scanf("%d", &t)
		arrB = append(arrB, t)
	}
	if arrA[k-1] < arrB[nB-m] {
		fmt.Println("YES")
		return
	}
	fmt.Println("NO")
}
