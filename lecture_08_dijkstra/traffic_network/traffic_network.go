package main

import (
	"container/heap"
	"fmt"
	"math"
)

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
	var T, n, m, k, s, t, u, v, w int
	fmt.Scanf("%d\n", &T)
	for {
		if T--; T < 0 {
			break
		}
		fmt.Scanf("%d %d %d %d %d\n", &n, &m, &k, &s, &t)
		graph := make(map[int][]pair)
		for i := 1; i < m+1; i++ {
			fmt.Scanf("%d %d %d\n", &u, &v, &w)
			graph[u] = append(graph[u], pair{v: v, w: w})
		}
		result := math.MaxInt
		for i := 0; i < k; i++ {
			copiedGraph := make(map[int][]pair)
			for k, v := range graph {
				cpValue := make([]pair, len(v))
				copy(cpValue, v)
				copiedGraph[k] = cpValue
			}
			dist := make([]int, n+1)
			for i := 1; i < n+1; i++ {
				dist[i] = math.MaxInt
			}
			fmt.Scanf("%d %d %d\n", &u, &v, &w)
			copiedGraph[u] = append(copiedGraph[u], pair{v: v, w: w})
			copiedGraph[v] = append(copiedGraph[v], pair{v: u, w: w})
			Dijkstra(s, copiedGraph, dist)
			if dist[t] < result {
				result = dist[t]
			}
		}
		if result == math.MaxInt {
			fmt.Println(-1)
			continue
		}
		fmt.Println(result)
	}
}

func Dijkstra(s int, graph map[int][]pair, dist []int) {
	dist[s] = 0
	h := minHeap{}
	heap.Push(&h, pair{v: s, w: 0})
	for {
		if len(h) == 0 {
			break
		}
		current := heap.Pop(&h).(pair)
		// this is an improvement on the original algorithm
		// this will check if the current top of the heap travel cost to the
		// value its stored in the dist table
		// this mean the node is not efficient and don't need to check it
		// due to the first condition is fail, the node travel cost is higher
		// so don't need to check the weight it hold
		if current.w != dist[current.v] {
			continue
		}
		for _, next := range graph[current.v] {
			if current.w+next.w < dist[next.v] {
				dist[next.v] = current.w + next.w
				heap.Push(&h, pair{v: next.v, w: dist[next.v]})
			}
		}
	}
}
