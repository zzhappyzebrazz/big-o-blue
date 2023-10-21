package main

import "fmt"

// TODO think clearly when doing algorithm

// question brief
// the game is 90
// if there is a 15 minute of boring time, the TV is off -> calculate the total watch time
// give the number of interesting time and the minutes of that time

// pseudo code
/*
read n

read arr[n] = t1, t2, ... tn

if t1 > 15 -> return t1

for i = 1; i < n; i++
 if i == n -1
    return 90
 if arr[i+1] - arr[i] > 15
    if arr[i] +15 > 90
	 return 90

    return arr[i] + 15
*/

func main() {
	var n, t int
	var arr []int
	_, err := fmt.Scanf("%d\n", &n)
	if err != nil {
		return
	}
	for i := 0; i < n; i++ {
		_, err := fmt.Scanf("%d", &t)
		if err != nil {
			return
		}
		if i == 0 && t > 15 {
			fmt.Println(15)
			return
		}
		arr = append(arr, t)
	}
	for i := 0; i < n; i++ {
		if i == n-1 {
			if arr[i]+15 > 90 {
				fmt.Println(90)
				return
			}
			fmt.Println(arr[i] + 15)
			return
		}
		if arr[i+1]-arr[i] > 15 {
			fmt.Println(arr[i] + 15)
			return
		}
	}
}
