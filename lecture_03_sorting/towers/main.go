package main

import "fmt"

/*
6 5 6 7 -> sort: 5 6 6 7 ->
map int int
key height value number of that length
return max value, number of key
*/

func main() {
	var n, max, temp int
	myMap := make(map[int]int)
	fmt.Scanf("%d\n", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &temp)
		value, exist := myMap[temp]
		if exist {
			value++
			myMap[temp] = value
		} else {
			myMap[temp] = 1
		}
	}
	for _, value := range myMap {
		if value >= max {
			max = value
		}
	}
	fmt.Println(max, len(myMap))
}
