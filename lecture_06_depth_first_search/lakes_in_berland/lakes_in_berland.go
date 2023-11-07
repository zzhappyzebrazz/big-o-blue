package main

// count all the connected lake '.' that is not at the edge of the matrix, map
// output the number of least connected lake, smallest side
// print out the map after the lake being filled with land '#'

/*
5 4 1
****
*..*
****
**.*
..**
-----
1
****
*..*
****
****
..**

the conner lake is at the edge so it not counted as a lake
an isOcean bool to make the lake that connected to the ocean -> in the edge
have a struct that count the size of the lake and the begin point of the lake
sort the lake struct and begin to fill the lake to land
we have 3 lake and the question ask for 2 lake
fill the first k - s lake
the output number is the size++ after each lake got fill
There are 2 DSF function, one to get all the lake
the second is to fill the lake


*/
import (
	"fmt"
	"sort"
)

type lake struct {
	Size, X, Y int
}

type Lakes []lake

func (e Lakes) Len() int {
	return len(e)
}

func (e Lakes) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func (e Lakes) Less(i, j int) bool {
	return e[i].Size < e[j].Size
}

var (
	X = []int{-1, +1, 0, 0}
	Y = []int{0, 0, -1, +1}
)

func main() {
	var n, m, k int
	fmt.Scanf("%d %d %d\n", &n, &m, &k)
	graph := make([][]string, n)
	visited := make([][]bool, n)
	for i := range graph {
		graph[i] = make([]string, m)
		visited[i] = make([]bool, m)
	}
	for i := 0; i < n; i++ {
		s := ""
		fmt.Scanf("%s\n", &s)
		for j, c := range s {
			graph[i][j] = string(c)
		}
	}
	lakes := []lake{}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if graph[i][j] == "." && !visited[i][j] {
				isOcean := false
				count := 0
				DSF(i, j, n, m, &count, &isOcean, graph, visited)
				if !isOcean {
					lakes = append(lakes, lake{Size: count, X: i, Y: j})
				}
			}
		}
	}
	// sort the lakes array
	sort.Sort(Lakes(lakes))
	count := 0
	for i := 0; i < len(lakes)-k; i++ {
		count += lakes[i].Size
		DSFfill(lakes[i].X, lakes[i].Y, n, m, graph)
	}
	fmt.Println(count)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Print(graph[i][j])
		}
		fmt.Println()
	}
}

func DSF(i, j, n, m int, count *int, isOcean *bool, graph [][]string, visited [][]bool) {
	visited[i][j] = true
	*count++
	if i == 0 || i == n-1 || j == 0 || j == m-1 {
		*isOcean = true
	}
	for k := 0; k < 4; k++ {
		x := i + X[k]
		y := j + Y[k]
		if x > -1 && x < n && y > -1 && y < m && graph[x][y] == "." && !visited[x][y] {
			DSF(x, y, n, m, count, isOcean, graph, visited)
		}
	}
}

func DSFfill(i, j, n, m int, graph [][]string) {
	graph[i][j] = "#"
	for k := 0; k < 4; k++ {
		x := i + X[k]
		y := j + Y[k]
		if x > -1 && x < n && y > -1 && y < m && graph[x][y] == "." {
			DSFfill(x, y, n, m, graph)
		}
	}
}
