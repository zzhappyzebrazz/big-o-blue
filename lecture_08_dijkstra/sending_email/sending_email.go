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
	var N, n, m, s, t, u, v, w int
	testcase := 0
	fmt.Scanf("%d\n", &N)
	for {
		if N--; N < 0 {
			break
		}
		testcase++
		fmt.Scanf("%d %d %d %d\n", &n, &m, &s, &t)
		dist := make([]int, n+1)
		for i := range dist {
			dist[i] = math.MaxInt
		}
		graph := make(map[int][]pair)
		for i := 0; i < m; i++ {
			fmt.Scanf("%d %d %d\n", &u, &v, &w)
			graph[u] = append(graph[u], pair{v: v, w: w})
			graph[v] = append(graph[v], pair{v: u, w: w})
		}
		Dijkstra(s, graph, dist)
		if dist[t] == math.MaxInt {
			fmt.Printf("Case #%d: unreachable\n", testcase)
			continue
		}
		fmt.Printf("Case #%d: %d\n", testcase, dist[t])
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
		if current.w != dist[current.v] {
			continue
		}
		for _, next := range graph[current.v] {
			if next.w+current.w < dist[next.v] {
				dist[next.v] = next.w + current.w
				heap.Push(&h, pair{v: next.v, w: dist[next.v]})
			}
		}
	}
}
