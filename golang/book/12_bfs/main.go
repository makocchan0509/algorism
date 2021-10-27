package main

import (
	"fmt"
	"strings"
)

func main() {

	var r, c int
	var sy, sx int
	var gy, gx int

	fmt.Scan(&r, &c)
	fmt.Scan(&sy, &sx)
	fmt.Scan(&gy, &gx)

	sy -= 1
	sx -= 1
	gy -= 1
	gx -= 1

	mz := make([][]string, r)

	var line string

	for i := 0; i < r; i++ {
		fmt.Scan(&line)
		mz[i] = strings.Split(line, "")
	}

	//fmt.Println(mz)
	dist := make([][]int, r)
	for i := 0; i < r; i++ {
		dy := make([]int, c)
		for z := 0; z < c; z++ {
			dy[z] = -1
		}
		dist[i] = dy
	}

	q := make([][]int, 0)
	q = append(q, []int{sy, sx})

	dist[sx][sy] = 0

	for len(q) > 0 {
		//m := 2
		//for m > 0 {
		t := q
		n := len(t)
		p := t[0]
		q = t[1:n]
		//	fmt.Print("start:")
		//	fmt.Println(q)
		//	fmt.Print("point:")
		//	fmt.Println(p)

		for _, a := range [][]int{[]int{p[0] - 1, p[1]}, []int{p[0] + 1, p[1]}, []int{p[0], p[1] - 1}, []int{p[0], p[1] + 1}} {
			//fmt.Println(a[0], a[1])
			if 0 > a[0] || a[0] > r || 0 > a[1] || a[1] > c {
				//fmt.Print("out:")
				//fmt.Println(a[0], a[1])
				continue
			} else if mz[a[0]][a[1]] == "#" {
				//fmt.Print("wall:")
				//fmt.Println(a[0], a[1])
				continue
			} else if dist[a[0]][a[1]] == -1 {
				//fmt.Println(a[0], a[1])
				dist[a[0]][a[1]] = dist[p[0]][p[1]] + 1
				q = append(q, []int{a[0], a[1]})
				//fmt.Println(q)
			}
		}
		//m -= 1
	}
	//fmt.Println(dist)
	fmt.Println(dist[gy][gx])
}
