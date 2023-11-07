package main

// nha hao tam
// path travel with weight on the edges
// get the longest connected path and calculate the all the weight on that path
// output the max of those path's weight

/*
1
6
1 2 3
2 3 4
2 6 2
6 4 6
6 5 5

-> this is a tree and the question is asking to find the diameter of the tree
-> diameter of a tree is the longest distance between 2 leaf node -> find this 2 leaf node

*/

import "fmt"

type weight struct {
	v, w int
}

var (
	leaf    = 0
	maxDist = 0
)

func main() {
	var T, n, u, v, w int
	fmt.Scanf("%d\n", &T)
	for t := 0; t < T; t++ {
		fmt.Scanf("%d\n", &n)
		graph := make(map[int][]weight, n+1)
		dist := []int{}
		for i := 0; i < n-1; i++ {
			fmt.Scanf("%d %d %d\n", &u, &v, &w)
			graph[u] = append(graph[u], weight{v, w})
		}
		for i := 1; i < n+1; i++ {
			dist = append(dist, -1)
		}
		// the first DSF to find the start node of graph diameter
		// choose a node and start DSF from it, return the leaf node at which the distance from start node is the longest
		DSF(1, graph, dist)
		
		for i := 1; i < len(dist); i++ {
			if array[i] > maxValue {
				maxValue = array[i]
				maxIndex = i
			}
		}
	}
}

func DSF(s int, graph map[int][]weight, dist []int) {
	maxDist = dist[s]
	if 
}

// def DFS(src):
// 	//
//     global leaf, max_dist
// 	// an array with len V+1 all -1
//     dist = [-1] * (V + 1)
// 	// the DSF stack
// 	s = [src]
// 	// distance table at the src node

//     dist[src] = 0

//     while len(s):
//         u = s.pop()
// 		// the node and its weight in the graph
//         for v, w in graph[u]:
// 			// the node is not visited
//             if dist[v] == -1:
// 			// update the distance table, at node v the distance is dist u + w
//                 dist[v] = dist[u] + w
// 				// update the value max_dist to the current distance of node v
//                 max_dist = max(max_dist, dist[v])
//                 s.append(v)
//     // return the leaf value is the index of the distance table which has the most value
//     leaf = dist.index(max(dist))

// t = int(input())

// for _ in range(t):
//     V = int(input())
//     E = V - 1
//     graph = [[] for _ in range(V + 1)]

//     for i in range(E):
//         u, v, w = map(int, input().split())

//         graph[u].append((v, w))
//         graph[v].append((u, w))
//     // save leaf node
//     leaf = 0
// 	// track the max dist in tree
//     max_dist = 0
// 	// this problem is to find the 2 leaf node that has the longest distance between
// 	// run DSF for a random node
//     DFS(1) // this first run is to find the first node
//     DFS(leaf) // the second run is to find the other node that has the longest distance

//     print(max_dist)
