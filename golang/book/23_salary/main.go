package main

import "fmt"

var emp [][]int
var N int

func main() {

	fmt.Scan(&N)

	emp = make([][]int, N)
	for i := 1; i < N; i++ {
		var b int
		fmt.Scan(&b)
		emp[b-1] = append(emp[b-1], i)
	}

	fmt.Println(dfs(0))

}

func dfs(i int) int {

	if len(emp[i]) == 0 {
		return 1
	} else {
		values := []int{}
		for _, v := range emp[i] {
			values = append(values, dfs(v))
		}
		return maxInSlice(values) + minInSlice(values) + 1
	}
}

func minInSlice(x []int) int {
	m := 100000000
	for _, v := range x {
		m = min(m, v)
	}
	return m
}

func maxInSlice(x []int) int {
	m := -100000000
	for _, v := range x {
		m = max(m, v)
	}
	return m
}

func min(l int, r int) int {

	if l > r {
		return r
	} else {
		return l
	}
}

func max(l int, r int) int {

	if l < r {
		return r
	} else {
		return l
	}
}
