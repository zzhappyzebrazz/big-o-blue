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
		path := []int{}
		for i := 0; i < e; i++ {
			path = append(path, -1)
		}
		visited := make([]bool, e)
		fmt.Scanf("%d\n", &n)
		for i := 0; i < n; i++ {
			fmt.Scanf("%d %d", &a, &b)
			graph[a] = append(graph[a], b)
			graph[b] = append(graph[b], a)
		}
		for key, _ := range graph {
			if !visited[key] {
				stackDSF(key, path, graph, visited)
			}
		}
		count := 0
		for _, i := range path {
			if i == -1 {
				count++
			}
		}
		fmt.Println(count)
	}
}

func stackDSF(s int, path []int, graph map[int][]int, visited []bool) {
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
				path[val] = current
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
