package main

import (
	"fmt"
	"strings"
)

var N, M int
var pond [][]string

func main() {
	fmt.Scan(&N, &M)
	pond = make([][]string, N)

	var p string
	var ans int

	for i := 0; i < N; i++ {
		fmt.Scan(&p)
		pond[i] = strings.Split(p, "")
	}

	for y := 0; y < N; y++ {
		for x := 0; x < M; x++ {
			if pond[y][x] == "W" {
				ans++
				dfs_pond(y, x)
			}
		}
	}
	fmt.Println(ans)
}

func dfs_pond(y, x int) {
	pond[y][x] = "."
	desc_pond()
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			ddx := x + dx
			ddy := y + dy
			if ddx < 0 || ddy < 0 || ddx >= M || ddy >= N {
				continue
			}

			if pond[ddy][ddx] == "W" {
				dfs_pond(ddy, ddx)
			}
		}
	}
	return
}

func desc_pond() {
	for i := 0; i < N; i++ {
		fmt.Println(strings.Join(pond[i], ""))
	}
	fmt.Println("")
}
