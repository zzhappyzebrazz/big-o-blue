package main

import "fmt"

// print out the number of largest connect
// the node in that connected cluster

/*
4 3
1 2
2 4
1 3 ---> 4

4 3
1 2
2 1
2 3 ---> 3

keep count the connected path and update to the max value
I think it is suitable more for the stack approach
at the end of the DSF call we can get the longest path
*/

func main() {
	var n, m, a, b int
	max := -1
	fmt.Scanf("%d %d\n", &n, &m)
	graph := make(map[int][]int)

	for i := 0; i < n; i++ {
		graph[i+1] = []int{}
	}
	for i := 0; i < m; i++ {
		fmt.Scanf("%d %d\n", &a, &b)
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}
	for i, _ := range graph {
		visited := make([]bool, n+1)
		if !visited[i] {
			count := DSF(i, graph, visited)
			if count > max {
				max = count
			}
		}
	}
	fmt.Println(max)
}

func DSF(s int, graph map[int][]int, visited []bool) int {
	count := 1
	stack := []int{s}
	visited[s] = true
	for {
		if len(stack) == 0 {
			break
		}
		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		for _, val := range graph[n] {
			if !visited[val] {
				stack = append(stack, val)
				visited[val] = true
				count++
			}
		}
	}
	return count
}
