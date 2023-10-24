package main

import (
	"fmt"
	"sort"
)

/*
n*2 friend -> n girls n boys, boys get twice tea as girls
w litters in tea pot
an array of cup that can hold ai litters of tea
out put the total amount of tea have to serve the friend with the give input
the output has to less than the w

in n, w
in arr[2n]
sort arr -> get least amount of tea that cup can hold -> the max amount can serve
the first half cup of tea is served to the girls and the upper half is for the boys
output the min of


ex: 2 4
1 1 1 1
max cup for girls = arr[0] = 1
max cup for girls from boys cup = arr[n/2] = 0.5
the tea must follow the min due to the condition -> girls get 0.5 each boys get 1 each
return 0.5*n + 2*0.5*n = 1 + 2 = 3

ex: 1 5
2 3 -> min(2, 1.5) = 1.5
-> 1.5*1 + 1.5*2*1 = 3

*/

func main() {
	var n, w, t float64
	fmt.Scanf("%f %f\n", &n, &w)
	var arr []float64
	for i := 0.0; i < 2*n; i++ {
		fmt.Scanf("%f", &t)
		arr = append(arr, t)
	}
	sort.Float64s(arr)
	m := min(arr[0], arr[int(n)]/2)
	fmt.Println(min(m*3*n, w))
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
