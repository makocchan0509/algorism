package main

import (
	"fmt"
)

func main() {

	var N, W int
	fmt.Scan(&N, &W)

	ws := make([]int, N+1)
	vs := make([]int, N+1)

	for i := 0; i < N+1; i++ {
		if i == 0 {
			ws[i] = 0
			vs[i] = 0
			continue
		}

		var sw, sv int
		fmt.Scan(&sw, &sv)
		ws[i] = sw
		vs[i] = sv
	}

	nap := make([][]int, N+1)
	for i := 0; i < N+1; i++ {
		var lin []int
		for y := 0; y < W+1; y++ {
			lin = append(lin, -1000000000)
		}
		nap[i] = lin
	}
	nap[0][0] = 0

	for i := 1; i < N+1; i++ {
		for w := 0; w < W+1; w++ {
			nap[i][w] = max(nap[i][w], nap[i-1][w])

			if w-ws[i] >= 0 {
				nap[i][w] = max(nap[i][w], nap[i-1][w-ws[i]]+vs[i])
			}

		}
	}
	ans := 0
	for i, arr := range nap {
		if i == N {
			for _, v := range arr {
				ans = max(ans, v)
			}
		}
	}
	fmt.Println(ans)

}

func max(l int, r int) int {
	if l > r {
		return l
	} else {
		return r
	}

}
