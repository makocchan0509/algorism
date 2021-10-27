package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n, m, q int

	fmt.Scan(&n, &m, &q)

	graph := make([][]bool, n)

	for i := 0; i < n; i++ {
		g := make([]bool, n)
		for y := 0; y < n; y++ {
			g[y] = false
		}
		graph[i] = g
	}
	//fmt.Println(graph)

	for i := 0; i < m; i++ {
		var u, v int
		fmt.Scan(&u, &v)

		u = u - 1
		v = v - 1

		graph[u][v] = true
		graph[v][u] = true
	}

	//fmt.Println(graph)

	var sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	c := make([]int, n)

	for i := 0; i < n; i++ {
		sc.Scan()
		c[i], _ = strconv.Atoi(sc.Text())
	}
	//fmt.Println(c)
	s := make([][]int, q)

	for i := 0; i < q; i++ {
		sl := make([]int, 3)
		sc.Scan()
		s1, _ := strconv.Atoi(sc.Text())
		sc.Scan()
		s2, _ := strconv.Atoi(sc.Text())
		s2 = s2 - 1
		sl[0] = s1
		sl[1] = s2
		if s1 == 2 {
			sc.Scan()
			s3, _ := strconv.Atoi(sc.Text())
			sl[2] = s3
		}
		s[i] = sl
	}
	fmt.Println(s)

	for i := 0; i < len(s); i++ {

		fmt.Println(c[s[i][1]])

		if s[i][0] == 2 {
			c[s[i][1]] = s[i][2]
		} else {
			arr := graph[s[i][1]]
			for z, v := range arr {
				if v {
					c[z] = c[s[i][1]]
				}
			}
		}
	}
}
