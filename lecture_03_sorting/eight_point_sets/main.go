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
example:
0 0
2 1
1 0
0 2
2 2
1 0
2 1
0 2
-> sort
0 0
0 1
0 2
1 0
1 2
2 0
2 1
2 2
*/

type point struct {
	x int
	y int
}

// the type that hold custom type
// this struct will implement the 2 interface method of sort package
// len() return the len of the array
// swap() the swap function
// less() the comparator, how to sort the array, return bool

type eightPointsSet []point

func (e eightPointsSet) Len() int {
	return len(e)
}

func (e eightPointsSet) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func (e eightPointsSet) Less(i, j int) bool {
	return e[i].x < e[j].x || (e[i].x == e[j].x && e[i].y < e[j].y)
}

func main() {
	var arr []point
	var t point
	for i := 0; i < 8; i++ {
		fmt.Scanf("%d %d\n", &t.x, &t.y)
		arr = append(arr, t)
	}
	sort.Sort(eightPointsSet(arr))
	// if condition to check all the point
	if arr[0].x == arr[1].x && arr[1].x == arr[2].x && arr[2].x < arr[3].x &&
		arr[3].x == arr[4].x && arr[4].x < arr[5].x &&
		arr[5].x == arr[6].x && arr[6].x == arr[7].x &&
		arr[0].y == arr[3].y && arr[3].y == arr[5].y && arr[5].y < arr[1].y &&
		arr[1].y == arr[6].y && arr[6].y < arr[2].y &&
		arr[2].y == arr[4].y && arr[4].y == arr[7].y {
		fmt.Println("respectable")
		return
	}

	// big-o solution
	/*
	read x[] and y[] and sort x, y
	generate a beautiful array of point that is respectable
	sort input arr
	compare the beautiful array and input arr, return the answer
	*/
	fmt.Println("ugly")
}
