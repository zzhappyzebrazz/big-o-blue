package main

import "fmt"

func main() {
	graph := map[int][]int{
		0: {1, 3},
		1: {0, 2, 3, 5},
		2: {1, 5},
		3: {0, 1, 4, 5},
		4: {3},
		5: {1, 2, 3},
	}
	s, f := 0, 5
	dsfWithStack(graph, s, f)
	dsfRecursion(graph, s, f)
}

func dsfWithStack(graph map[int][]int, s, f int) {
	// add start to the stack
	stack := []int{s}
	path := []int{}
	for i := 0; i < len(graph); i++ {
		path = append(path, -1)
	}
	// mark the path to start took 0
	path[s] = 0
	for {
		if len(stack) == 0 {
			break
		}
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// val := graph[current]
		for _, val := range graph[current] {
			if path[val] == -1 {
				stack = append(stack, val)
				path[val] = current
			}
		}
	}
	// print out the path
	result := []int{}
	current := f
	for {
		result = append(result, current)
		if current == s {
			break
		}
		current = path[current]
	}
	for i := len(result) - 1; i >= 0; i-- {
		fmt.Printf("%d ", result[i])
	}
	fmt.Println()
}

func dsfRecursion(graph map[int][]int, s, f int) {
	path := []int{}
	for i := 0; i < len(graph); i++ {
		path = append(path, -1)
	}
	path[s] = 0
	recursion(graph, path, s)
	result := []int{}
	current := f
	for {
		result = append(result, current)
		if current == s {
			break
		}
		current = path[current]
	}
	for i := len(result) - 1; i >= 0; i-- {
		fmt.Printf("%d ", result[i])
	}
	fmt.Println()
}

func recursion(graph map[int][]int, path []int, s int) {
	for _, val := range graph[s] {
		if path[val] == -1 {
			path[val] = s
			recursion(graph, path, val)
		}
	}
}
