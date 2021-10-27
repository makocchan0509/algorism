package main

import "fmt"

var edges [][]int
var length []int
var done []bool

func main() {

	var N, M int
	fmt.Scan(&N, &M)
	edges = make([][]int, N)
	indeg := make([]int, N)

	for i, _ := range indeg {
		indeg[i] = 0
	}

	for i := 0; i < M; i++ {
		var x, y int
		fmt.Scan(&x, &y)
		edges[x-1] = append(edges[x-1], y-1)
		indeg[y-1] += 1
	}

	length = make([]int, N)
	for i, _ := range length {
		length[i] = 0
	}

	done = make([]bool, N)
	for i, _ := range done {
		done[i] = false
	}
	//fmt.Println(edges)
	//fmt.Println(indeg)

	for i, v := range indeg {
		if v == 0 {
			rec(i)
		}
	}
	//fmt.Println(length)
	fmt.Println(maxInSlice(length))
}

func rec(i int) int {

	if done[i] {
		return length[i]
	} else {
		length[i] = 0

		for _, j := range edges[i] {
			length[i] = max(length[i], rec(j)+1)
		}
		done[i] = true
		return length[i]
	}
}

func max(l int, r int) int {

	if l < r {
		return r
	} else {
		return l
	}
}

func maxInSlice(x []int) int {
	m := -100000000
	for _, v := range x {
		m = max(m, v)
	}
	return m
}
