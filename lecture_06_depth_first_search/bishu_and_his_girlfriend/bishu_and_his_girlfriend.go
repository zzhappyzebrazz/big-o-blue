package main

import "fmt"

/*
Note: if a graph with N nodes and N-1 edges and all fully connected, no loop
-> this is a tree and in this scenario, the output of BSF and DSF is the same
it's recommended that to use recursion DSF in tree traversal for a better approach
be familiar with recursive DSF so that it will be more natural when learn about further
algorithm about graph, tree and trie

*/

func main() {
	var n, q, u, v, x int
	fmt.Scanf("%d\n", &n)
	minDist, minId := n, n
	graph := map[int][]int{}
	dist := make([]int, n+1)
	for i := 0; i < n-1; i++ {
		fmt.Scanf("%d %d", &u, &v)
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}
	for i := 0; i < n; i++ {
		dist[i+1] = -1
	}
	girl := []int{}
	fmt.Scanf("%d\n", &q)
	for i := 0; i < q; i++ {
		fmt.Scanf("%d", &x)
		girl = append(girl, x)
	}
	dist[1] = 0
	DSF(graph, dist, girl, 1, &minDist, &minId)
	fmt.Println(minId)
}

func DSF(graph map[int][]int, dist, girl []int, s int, minDist, minId *int) {
	for _, next := range graph[s] {
		if dist[next] == -1 {
			dist[next] = dist[s] + 1
			if isExist(girl, next) {
				if dist[next] < *minDist {
					*minDist = dist[next]
					*minId = next
				}
				if dist[next] == *minDist && next < *minId {
					*minId = next
				}
			}
			DSF(graph, dist, girl, next, minDist, minId)
		}
	}
}

func isExist(arr []int, s int) bool {
	for _, i := range arr {
		if i == s {
			return true
		}
	}
	return false
}
