package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {

	var n, m, q int
	fmt.Scan(&n, &m, &q)
	graph := make([][]int, n)
	//fmt.Println(graph)
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Scan(&u, &v)

		u = u - 1
		v = v - 1
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	sc.Split(bufio.ScanWords)

	c := makeIntArr(n)

	//fmt.Println(c)

	query := make([][]int, q)

	for i := 0; i < q; i++ {
		query[i] = makeIntArrExt(3)
	}

	for i := 0; i < len(query); i++ {

		fmt.Println(c[query[i][1]])

		if query[i][0] == 2 {
			c[query[i][1]] = query[i][2]
		} else {
			for _, v := range graph[query[i][1]] {
				c[v] = c[query[i][1]]
			}
		}
	}
}

func makeIntArr(x int) []int {

	arrInt := make([]int, x)
	for i := 0; i < x; i++ {
		sc.Scan()
		v, _ := strconv.Atoi(sc.Text())
		arrInt[i] = v
	}
	return arrInt
}

func makeIntArrExt(x int) []int {

	arrInt := make([]int, x)

	sc.Scan()
	s1, _ := strconv.Atoi(sc.Text())

	sc.Scan()
	s2, _ := strconv.Atoi(sc.Text())
	s2 = s2 - 1

	arrInt[0] = s1
	arrInt[1] = s2

	if s1 == 2 {
		sc.Scan()
		s3, _ := strconv.Atoi(sc.Text())
		arrInt[2] = s3
	}

	return arrInt
}
