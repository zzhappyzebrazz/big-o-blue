package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// give an map of character
// travel that map and find is there a connect path that give up the phrase ALLIZZWELL
/*
5
3 6
AWE.QX
LLL.EO
IZZWLL  ---> ok

1 10
ALLIZZWELL ----> ok


2 9
A.L.Z.E..
.L.I.W.L. ----> NO


3 3
AEL
LWZ
LIZ ----> NO

1 10
LLEWZZILLA ----> YES case reverse

*/

var (
	arr = []string{
		"A", "L", "L", "I", "Z", "Z", "W", "E", "L", "L",
	}
	dr = []int{
		-1, 0, 1, -1, 1, -1, 0, 1,
	}
	dc = []int{
		1, 1, 1, 0, 0, -1, -1, -1,
	}
	T, R, C int
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Scanf("%d\n", &T)
	for t := 0; t < T; t++ {
		scanner.Scan()
		testCaseInfo := strings.Fields(scanner.Text())
		R, _ = strconv.Atoi(testCaseInfo[0])
		C, _ = strconv.Atoi(testCaseInfo[1])
		graph := make([][]string, R)
		for i := range graph {
			graph[i] = make([]string, C)
		}
		// visited := [][]bool{}
		for i := 0; i < R; i++ {
			scanner.Scan()
			s := scanner.Text()
			for j, c := range s {
				graph[i][j] = string(c)
				// visited[i][j] = false
			}
		}
		result := "NO"
		for i := 0; i < R; i++ {
			for j := 0; j < C; j++ {
				if graph[i][j] == "A" {
					if DSF(i, j, 0, graph) {
						result = "YES"
						break
					}
				}
			}
		}
		fmt.Println(result)
		scanner.Scan() // this line just to satisfied the SPOJ kine of question with multiple testcase
	}

}

func DSF(i, j, index int, graph [][]string) bool {
	index++
	for k := 0; k < 8; k++ {
		ir := i + dr[k]
		jc := j + dc[k]
		if ir > -1 && ir < R && jc > -1 && jc < C && graph[ir][jc] == arr[index] {
			if index == 9 {
				return true
			}
			check := DSF(ir, jc, index, graph)
			if check {
				return true
			}
			if !check && k == 7 {
				return false
			}
		}
	}
	return false
}
