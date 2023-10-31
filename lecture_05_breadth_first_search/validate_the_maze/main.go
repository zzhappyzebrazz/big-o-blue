package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X, Y int
}

func main() {
	input := bufio.NewReader(os.Stdin)
	t := readInt(input)
	for {
		if t == 0 {
			break
		}
		t--
		var (
			maze    [20][20]string
			visited [20][20]bool
			n, m    int
			sx      int = -1
			sy      int = -1
			ex      int = -1
			ey      int = -1
		)

		temp := readArrInt(input)
		n = temp[0]
		m = temp[1]
		open := 0
		for i := 0; i < n; i++ {
			line := readStringLine(input)
			for j := 0; j < m; j++ {
				maze[i][j] = line[j]
				if (i == 0 || i == n-1 || j == 0 || j == m-1) && line[j] == "." {
					open++
					if sx == -1 && sy == -1 {
						sx = i
						sy = j
					} else {
						ex = i
						ey = j
					}
				}
			}
		}
		if open != 2 {
			fmt.Println("invalid")
			continue
		}
		bsf(maze, &visited, sx, sy, n, m)
		if visited[ex][ey] {
			fmt.Println("valid")
			continue
		}
		// reset the visited table
		fmt.Println("invalid")
		continue
	}

}

func bsf(maze [20][20]string, visited *[20][20]bool, sx, sy, n, m int) {
	// create the queue
	q := []Point{
		{sx, sy},
	}
	for {
		if len(q) == 0 {
			break
		}
		// pop the first element out of the queue
		x, y := q[0].X, q[0].Y
		q = q[1:]
		// create the array with all the four points connected to the current node
		v := []Point{
			{x + 1, y},
			{x - 1, y},
			{x, y + 1},
			{x, y - 1},
		}
		for _, val := range v {
			// check the boundary of the connected point
			if val.X >= 0 && val.X < n && val.Y >= 0 && val.Y < m && maze[val.X][val.Y] == "." && !visited[val.X][val.Y] {
				q = append(q, val)
				visited[val.X][val.Y] = true
			}
		}
	}
}

func readInt(in *bufio.Reader) int {
	nStr, _ := in.ReadString('\n')
	nStr = strings.ReplaceAll(nStr, "\r", "")
	nStr = strings.ReplaceAll(nStr, "\n", "")
	n, _ := strconv.Atoi(nStr)
	return n
}

func readLine(in *bufio.Reader) []string {
	line, _ := in.ReadString('\n')
	line = strings.ReplaceAll(line, "\r", "")
	line = strings.ReplaceAll(line, "\n", "")
	numbs := strings.Split(line, " ")
	return numbs
}

func readStringLine(in *bufio.Reader) []string {
	line, _ := in.ReadString('\n')
	line = strings.ReplaceAll(line, "\r", "")
	line = strings.ReplaceAll(line, "\n", "")
	var result []string
	for _, char := range line {
		result = append(result, string(char))
	}
	return result
}

func readArrInt(in *bufio.Reader) []int {
	numbs := readLine(in)
	arr := make([]int, len(numbs))
	for i, n := range numbs {
		val, _ := strconv.Atoi(n)
		arr[i] = val
	}
	return arr
}
