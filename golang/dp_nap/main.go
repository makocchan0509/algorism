package main

import (
	"fmt"
	"math"
)

var n, W int
var w, v int
var ca [][]int
var dp [][]int

func main() {
	fmt.Scan(&n)
	ca = make([][]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&w, &v)
		ca[i] = []int{w, v}
	}
	dp = make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = []int{-1, -1, -1, -1, -1, -1}
	}
	fmt.Scan(&W)
	fmt.Println(rec(0, W))
}

func rec(i, j int) int {
	if dp[i][j] >= 0 {
		return dp[i][j]
	}

	var res int
	if i == n {
		res = 0
	} else if j < ca[i][0] {
		res = rec(i+1, j)
	} else {
		res = int(math.Max(float64(rec(i+1, j)), float64(rec(i+1, j-ca[i][0])+ca[i][1])))
	}
	dp[i][j] = res
	return res
}
