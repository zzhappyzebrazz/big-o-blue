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

	queue := []int{}
	path := map[int]int{
		0: -1,
		1: -1,
		2: -1,
		3: -1,
		4: -1,
		5: -1,
	}
	s, f := 0, 5
	// breadth-first search
	queue = append(queue, s)
	for {
		if len(queue) < 1 {
			break
		}
		next := graph[queue[0]]
		for _, val := range next {
			if path[val] == -1 && val != s {
				queue = append(queue, val)
				path[val] = queue[0]
			}
		}
		queue = queue[1:]
	}
	// print out the path
	if s == f {
		fmt.Println(s)
		return
	}
	if path[f] == -1 {
		fmt.Println(-1)
		return
	}
	b := []int{}
	for {
		b = append(b, f)
		f = path[f]
		if f == s {
			b = append(b, f)
			break
		}
	}
	for i := len(b) - 1; i >= 0; i-- {
		fmt.Printf("%d ", b[i])
	}
}
