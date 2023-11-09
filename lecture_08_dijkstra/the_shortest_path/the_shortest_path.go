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
	var s, n, p, v, w, r int
	var name string
	fmt.Scanf("%d\n", &s)
	for {
		if s--; s < 0 {
			break
		}
		fmt.Scanf("%d\n", &n)

		cities := make(map[string]int)
		graph := make(map[int][]pair)
		for i := 1; i < n+1; i++ {
			fmt.Scanln(&name)
			cities[name] = i
			fmt.Scanf("%d\n", &p)
			for j := 0; j < p; j++ {
				fmt.Scanf("%d %d\n", &v, &w)
				// directed graph
				graph[i] = append(graph[i], pair{v: v, w: w})
			}
		}
		fmt.Scanf("%d\n", &r)
		for i := 0; i < r; i++ {
			var name1, name2 string
			fmt.Scanf("%s %s\n", &name1, &name2)
			dist := make([]int, n+1)
			for i := range dist {
				dist[i] = math.MaxInt
			}
			Dijkstra(cities[name1], graph, dist)
			fmt.Println(dist[cities[name2]])
		}
		fmt.Scanf("\n")
	}
}

func Dijkstra(s int, graph map[int][]pair, dist []int) {
	// set the distance of the start node = 0
	dist[s] = 0
	// create the heap
	h := minHeap{}
	// push the first element into the heap
	heap.Push(&h, pair{v: s, w: 0})
	for {
		if len(h) == 0 {
			break
		}
		current := heap.Pop(&h).(pair)
		if dist[current.v] != current.w {
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
