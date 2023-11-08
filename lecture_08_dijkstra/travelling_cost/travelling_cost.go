package main

import (
	"container/heap"
	"fmt"
	"math"
)

/*
Chuẩn bị dữ liệu cho thuật toán dijkstra
chuyể ma trân kề/danh sách cạnh kề về graph
đỉnh              0    1    2	3
đỉnh kề/trọng sơ 1,1  2,5  5,1 0,2
					  3,2
					  5,7

mảng chi phí dist, khởi tạo với giá trị INF

mảng lưu vết path

hàng đợi ưu tiên, priority queue, lưu cặp giá trị đỉnh, chi phí -> min heap với giá trị heapify là chi phí


*/

type pair struct {
	v, w int
}

type minHeap []pair

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].w < h[j].w }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(pair))
}

func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	// the distance table, I think we don't need the path table to record the path, just need to update the dist table
	dist := make([]int, 501)
	for i := range dist {
		dist[i] = math.MaxInt
	}
	var t, s, v, w, b, n int
	fmt.Scanf("%d\n", &t)
	// the graph table
	graph := make(map[int][]pair)
	for i := 0; i < t; i++ {
		fmt.Scanf("%d %d %d\n", &s, &v, &w)
		// this is a undirected graph, a road that mean you can travel back and forth on the road
		graph[s] = append(graph[s], pair{v: v, w: w})
		graph[v] = append(graph[v], pair{v: s, w: w})
	}
	fmt.Scanf("%d\n%d\n", &b, &n)
	Dijkstra(b, graph, dist)
	for i := 0; i < n; i++ {
		temp := 0
		fmt.Scanf("%d\n", &temp)
		if dist[temp] == math.MaxInt {
			fmt.Println("NO PATH")
		} else {
			fmt.Println(dist[temp])
		}
	}
}

// using priority queue -> heap (this case min heap by node distance)
// this will not do the Dijsktra recursion version that cost O(n2)
// time complexity O(Nlogn)
func Dijkstra(s int, graph map[int][]pair, dist []int) {
	h := minHeap{}
	// start node
	heap.Push(&h, pair{v: s, w: 0})
	dist[s] = 0
	for {
		if len(h) == 0 {
			break
		}
		current := heap.Pop(&h).(pair)
		// a condition to by pass the checking of the a pair value saved  in the heap that has the cost
		// travel to it is not equal to the value of its cost in the dist table
		if dist[current.v] != current.w {
			continue
		}
		for _, next := range graph[current.v] {
			// the travel cost from the current node to its next node
			// is less than the value stored in the dist table
			if current.w+next.w < dist[next.v] { // the weight value is from the heap, not stored in the graph
				// update the dist table at the current next node
				// equal to the current node cost and the next node cost
				dist[next.v] = current.w + next.w
				// push the new value to the priority min heap -> mean by the weight value
				heap.Push(&h, pair{v: next.v, w: dist[next.v]})
				// you can update the path table here
				// path[it.v] = n.v
			}
		}
	}
}
