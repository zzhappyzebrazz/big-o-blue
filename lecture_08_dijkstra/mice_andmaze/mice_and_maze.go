package main

import (
	"container/heap"
	"fmt"
	"math"
)

/*
the question ask to find the number of mice get out of the maze in a limit of time
the graph is directed, the cells are nodes, the obstacles is the weight
the time for the mice to escape is the cost in dist table
input
N <- number of cell N < 100
E <- number of exit cell
T <- deadline
M <- connections in maze

output the number of node that has the dist value that <= T
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

	var n, t, e, m, u, v, w int
	fmt.Scanf("%d\n%d\n%d\n%d\n", &n, &e, &t, &m)
	graph := make(map[int][]pair)
	for i := 0; i < m; i++ {
		fmt.Scanf("%d %d %d\n", &u, &v, &w)
		// this is a directed graph
		graph[u] = append(graph[u], pair{v: v, w: w})
	}
	// start the count = 1 is the default value that the mice at the exist cell always get out 
	count := 1
	for u, _ := range graph {
		if u == e {
			continue
		}
		dist := make([]int, n+1)
		for i := range dist {
			dist[i] = math.MaxInt
		}
		Dijkstra(u, graph, dist)
		if dist[e] <= t {
			count++
		}
	}
	fmt.Println(count)
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
