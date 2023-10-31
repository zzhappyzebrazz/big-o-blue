package main

import "fmt"

func main() {
	var c, count int
	var p, x uint64
	var n string
	for {
		fmt.Scanf("%d %d", &p, &c)
		if p == 0 && c == 0 {
			break
		}
		var q []uint64
		for i := 1; i < int(min(p, uint64(c)))+1; i++ {
			q = append(q, uint64(i))
		}
		count += 1
		fmt.Printf("Case %d:\n", count)
		for i := 0; i < c; i++ {
			fmt.Scanf("%s", &n)
			if n == "N" {
				fmt.Printf("%d\n", q[0])
				q = append(q, q[0])
				q = q[1:]
			} else {
				fmt.Scanf("%d", &x)
				var temp []uint64
				temp = append(temp, x)
				for _, val := range q {
					if val != x {
						temp = append(temp, val)
					}
				}
				q = temp
			}
		}
	}
}

func min(p, c uint64) uint64 {
	if p < c {
		return p
	}
	return c
}
