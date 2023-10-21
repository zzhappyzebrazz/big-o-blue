// Books
/*
n number of book
t total free minutes
a[] the time it take to read the a[i] book
you can start at any given a[i]
how many book can he read
ex: 4 book, 5 free mins
a[] = 3 1 2 1
start at a[3]
-> a[3] + a[0] + a[1] = 5


*/

package main

import (
	"fmt"
)

func main() {
	// get input
	var n, t, temp int
	var arr []int
	_, err := fmt.Scanf("%d %d\n", &n, &t)
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
	if n == 1 {
		if arr[1] > t {
			fmt.Println(0)
		} else {
			fmt.Println(1)
		}
	}

	// Big-O example
	// create two pointer p_start and p_end
	// let make the first run with an example of array [ps, pe]
	// increase ps but no need to set the pe back to start
	// until the new sum of [ps, pe] satisfied the condition, increase the pe
	// keep running until the pe read the end

	// a[] 3 1 2 1
	// loop ps = pe = 0; sum = 3 < 5; max = 1
	// ps = 0, pe = 1; sum = 4 < 5; max = 2
	// ps = 0 -> 1, pe = 2; sum = 6 > 5 -> sum = 6 - arr[0] = 3; max = 2
	// ps = 1, pe = 3; sum = 3 + 1 < 5; max = 3 -> end

	var max, sum, ps int
	for pe := 0; pe < n; pe++ {
		sum += arr[pe]
		if sum > t {
			sum -= arr[ps]
			ps++
			continue
		}
		if max <= pe-ps+1 {
			max = pe - ps + 1
		}
	}
	fmt.Println(max)
}

func myTwoPointerApproach(n, t int, arr []int) {
	// two pointer approach
	// the first pointer start at the begin, and the second one stay at the end
	// calculate all the time to read the books
	// minus the pointer p1 run
	// a[] 3 2 1 1 -> sum = 3 + 2 + 1 + 1 = 7
	// sum -= a[0]
	// if sum < n
	// return n - i - 1
	// not working
	var max, sum int
	for i := 0; i < n; i++ {
		sum += arr[i]
	}
	if sum <= t {
		fmt.Println(n)
		return
	}
	for i := 0; i < n; i++ {
		sum -= arr[i]
		if sum <= t {
			max = n - i - 1
			fmt.Println(max)
			return
		}
	}
}

func bruteFore(n, t int, arr []int) {
	// brute force
	var max, sum int
	for i := 0; i < n; i++ {
		sum = arr[i]
		if sum > t {
			continue
		}
		c := 1
		arr1 := append(arr[i:], arr[0:i]...)
		for j := 1; j < n; j++ {
			sum += arr1[j]
			if sum <= t {
				c++

			} else {
				break
			}
		}
		if c > max {
			max = c
		}
	}
	fmt.Println(max)
}
