package main

import "fmt"

// detect loop in a non-directed graph
// be careful with the not fully connected graph
// cycle detect
/*
3

2 1
1 2 ---> NO

2 2
1 2
2 1 ----> YES

4 4
2 3
3 4
4 2
1 3 ----> YES

my solution is to check if the node in the visited table is being checked again for the second time
this means there is a loop in the graph
*/

func main() {
	var t int
	fmt.Scanf("%d\n", &t)
	for {
		if t < 1 {
			break
		}
		t--
		var n, e, a, b int
		fmt.Scanf("%d %d\n", &n, &e)
		// visited := make([]bool, n+1)
		graph := make(map[int][]int)
		/*
			the inPath has 3 state:
			 inPath[u] = 0 -> the node is not visited, not in path -> initial state
			 inPath[u] = 1 -> the node is visited, inPath -> to check the cycle in the path
			 inPath[u] = 2 -> the node is done DSF
		*/
		inPath := make([]int, n+1)
		for i := 1; i <= n; i++ {
			graph[i] = []int{}
			inPath[i] = 0
		}
		for i := 0; i < e; i++ {
			fmt.Scanf("%d %d\n", &a, &b)
			// only add the connected node to one side to avoid the checking again the parent of a node
			// 0->1 only if also assign 1->0 this automatically is loop, if not we have to check the condition of parent node
			graph[a] = append(graph[a], b)
		}
		// run DSF and if the node is true in visited[] then out put YES
		result := "NO"
		// check all the node in the graph in case not fully-connected graph
		for key, _ := range graph {
			if inPath[key] == 0 {
				if DSF(key, graph, inPath) {
					result = "YES"
					break
				}
			}
		}
		fmt.Println(result)
	}
}

// this is a recursive DSF
func DSF(s int, graph map[int][]int, inPath []int) bool {
	// the node is visited and in path
	inPath[s] = 1
	for _, node := range graph[s] {
		// need condition to check the parent node
		if inPath[node] == 0 {
			if DSF(node, graph, inPath) {
				return true
			}
		} else if inPath[node] == 1 {
			return true
		}
	}
	// the node is done DSF set the state to 2
	inPath[s] = 2
	return false
}
