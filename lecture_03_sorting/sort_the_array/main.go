package main

// given an array of distinct number
// check is it possible to reverse a segment of array a so that the array is sorted
/*
4
2 1 3 4
yes
1 2

4
3 1 2 4
no


*/

func main() {
	// var n, temp int
	// fmt.Scanf("%d\n", &n)
	// for int
}

// TODO: Unsolved
func arrayReverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func isSorted(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}
