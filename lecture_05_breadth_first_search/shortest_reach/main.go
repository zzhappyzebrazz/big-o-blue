package main

import (
	"fmt"
)

// TODO Should implement bufio scanner to read input faster, more efficient
// this struct combine the node, graph and the path in one
type Graph struct {
	Cost  int
	Edges []int
}

func main() {
	// read input, number of test case
	q := 0
	fmt.Scanf("%d\n", &q)
	for i := 0; i < q; i++ {
		n, m, s := 0, 0, 0
		fmt.Scanf("%d %d\n", &n, &m)
		graph := map[int]Graph{}
		for j := 1; j < n+1; j++ {
			graph[j] = Graph{-1, []int{}}
		}
		for j := 0; j < m; j++ {
			u, v := 0, 0
			fmt.Scanf("%d %d\n", &u, &v)
			// update the graph
			value := graph[u]
			value.Edges = append(value.Edges, v)
			graph[u] = value
			value = graph[v]
			value.Edges = append(value.Edges, u)
			graph[v] = value
		}
		fmt.Scanf("%d\n", &s)
		// path := bsf(graph, n, s)
		// printResult(path, n, s)
		findCost(&graph, s)
		for j := 1; j < n+1; j++ {
			path := graph[j]
			if path.Cost != 0 {
				fmt.Printf("%d ", path.Cost)
			}
		}
		fmt.Println()
	}

}

func findCost(graph *map[int]Graph, s int) {
	queue := []int{s}
	weight := 6
	current_cost := 0
	for {
		if len(queue) < 1 {
			break
		}
		for _, node := range queue {
			path := (*graph)[node]
			if path.Cost == -1 {
				path.Cost = current_cost
				queue = append(queue, path.Edges...)
				(*graph)[node] = path
			}
		}
		queue = queue[1:]
		current_cost += weight
	}
}

// naive approach, got time limited exceed
func bsf(graph map[int][]int, n, s int) map[int]int {
	path := map[int]int{}
	queue := []int{}
	queue = append(queue, s)
	for i := 1; i < n+1; i++ {
		path[i] = -1
	}
	for {
		if len(queue) < 1 {
			break
		}
		next := graph[queue[0]]
		for _, val := range next {
			if path[val] == -1 && val != s {
				path[val] = queue[0]
				queue = append(queue, val)
			}
		}
		queue = queue[1:]
	}
	return path
}

// naive approach, got time limited exceed
func printResult(path map[int]int, n, s int) {
	// don't print the node s
	b := []int{}
	for i := 1; i < n+1; i++ {
		if i == s {
			continue
		}
		// count the edge it take to all other nodes, and print it out multiple by 6
		if path[i] == -1 {
			b = append(b, -1)
			continue
		}
		c := 0
		f := i
		for {
			f = path[f]
			c++
			if f == s {
				break
			}
		}
		b = append(b, c)
	}
	for _, val := range b {
		if val == -1 {
			fmt.Printf("%d ", -1)
			continue
		}
		fmt.Printf("%d ", val*6)
	}
	fmt.Println()
}
