package main

import (
	"fmt"
)

var ws []int
var vs []int
var ans int
var W int

func main() {

	var N int
	fmt.Scan(&N, &W)

	ws = make([]int, N)
	vs = make([]int, N)

	for i := 0; i < N; i++ {
		var sw, sv int
		fmt.Scan(&sw, &sv)
		ws[i] = sw
		vs[i] = sv
	}

	ans = 0

	fmt.Println(dfs(0, 0, 0))
}

func dfs(i int, ans int, weight int) int {
	//	fmt.Println(i, ans, weight)

	if i == len(ws) {
		return ans
	}
	if weight+ws[i] > W {
		return dfs(i+1, ans, weight)
	}
	return max(dfs(i+1, ans+vs[i], weight+ws[i]), dfs(i+1, ans, weight))
}

func max(l int, r int) int {
	if l > r {
		return l
	} else {
		return r
	}

}
