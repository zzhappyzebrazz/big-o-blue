package main

import "fmt"

/*
Heap là cấu trúc cây nhị phân hoàn chỉnh
cách đọc cây heap từ 1 mảng
node bên trái của index hiện tại: 2n+1, node bên phải 2n+2

Heap dùng để cài đặt và giải quyết bài toán liên quan đến hàng đợi ưu tiên: priority queue
Dùng để tối ưu các thuật toán Dijkstra, Prim,...
Thuật toán sắp xếp heapsort
Xây dựng cây Heap từ màng O(n)
tìm max, min O(1) -> đầu heap hoặc cuối heap
thêm, xóa một phần từ vào heap O(logn) -> chia cây để thêm phần tử

thuật toàn heapify -> xây dựng cây heap từ mảng

*/

func main() {
	arr := []int{
		7, 12, 6, 10, 17, 15, 2, 4,
	}
	for i := len(arr)/2 - 1; i >= 0; i-- {
		minHeapify(i, arr)
	}
	fmt.Println(arr)
}

// min heap
func minHeapify(index int, arr []int) {
	// node index n/2-1
	n := len(arr)
	left := index*2 + 1
	right := index*2 + 2
	smaller := index

	if left < n && arr[smaller] > arr[left] {
		smaller = left
	}
	if right < n && arr[smaller] > arr[right] {
		smaller = right
	}
	if smaller != index {
		arr[index], arr[smaller] = arr[smaller], arr[index]
		minHeapify(smaller, arr)
	}
}

// pop -> delete the first element of the heap
func pop(arr []int) {
	// swap the last element of the heap to the top
	// perform heapify to re-order the heap
	if len(arr) == 0 {
		return
	}
	arr[0] = arr[len(arr)-1]
	arr = arr[:len(arr)-1]
	for i := len(arr)/2 - 1; i >= 0; i-- {
		minHeapify(i, arr)
	}
}

// push -> add a new node to the heap
func push(n int, arr []int) {
	// add the new node the the heap
	// recursive check the parent of the new node is satisfied
	arr = append(arr, n)
	for i := len(arr) - 1; arr[(i-1)/2] > arr[i]; i = (i - 1) / 2 {
		arr[i], arr[(i-1)/2] = arr[(i-1)/2], arr[i]
	}
}
