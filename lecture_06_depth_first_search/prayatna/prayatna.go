package main

import "fmt"

func main() {
	var t, e, n, a, b int
	fmt.Scanf("%d\n", &t)
	for {
		if t < 1 {
			break
		}
		t--
		fmt.Scanf("%d\n", &e)
		graph := map[int][]int{}
		for i := 0; i < e; i++ {
			graph[i] = []int{}
		}
		visited := make([]bool, e)
		fmt.Scanf("%d\n", &n)
		for i := 0; i < n; i++ {
			fmt.Scanf("%d %d", &a, &b)
			graph[a] = append(graph[a], b)
			graph[b] = append(graph[b], a)
		}
		count := 0
		for key, _ := range graph {
			if !visited[key] {
				count++
			}
			if !visited[key] {
				stackDSF(key, graph, visited)
			}
		}
		fmt.Println(count)
	}
}

func stackDSF(s int, graph map[int][]int, visited []bool) {
	stack := []int{s}
	visited[s] = true
	for {
		if len(stack) == 0 {
			break
		}
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// val := graph[current]
		for _, val := range graph[current] {
			if !visited[val] {
				visited[val] = true
				stack = append(stack, val)
			}
		}
	}
}

// stack call too much ->
func DSF(s int, path []int, graph map[int][]int, visited []bool) {
	visited[s] = true
	nodes := graph[s]
	for _, i := range nodes {
		if !visited[i] {
			path[i] = s
			DSF(i, path, graph, visited)
		}
	}
}
