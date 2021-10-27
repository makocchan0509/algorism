package main

import (
	"fmt"
	"strings"
)

var vstd [][]bool
var ro [][]string
var h, w int

func main() {
	var sy, sx, gy, gx int
	fmt.Scan(&h, &w)

	ro = make([][]string, h)

	for i := 0; i < h; i++ {
		var s string
		fmt.Scan(&s)
		ro[i] = strings.Split(s, "")
	}

	for i, v := range ro {
		for j, vv := range v {
			if vv == "s" {
				sy = i
				sx = j
			} else if vv == "g" {
				gy = i
				gx = j
			}
		}
	}

	vstd = make([][]bool, h)
	for i := 0; i < h; i++ {
		ry := make([]bool, w)
		for j := 0; j < w; j++ {
			ry[j] = false
		}
		vstd[i] = ry
	}

	//fmt.Println(ro)
	//fmt.Println(vstd)
	//fmt.Println(sy, sx)
	//fmt.Println(gy, gx)

	dfs(sy, sx)

	if vstd[gy][gx] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}

func dfs(y int, x int) {

	vstd[y][x] = true

	for _, hr := range [][]int{[]int{y - 1, x}, []int{y + 1, x}, []int{y, x - 1}, []int{y, x + 1}} {
		//fmt.Println(hr[0], hr[1])
		if 0 > hr[0] || hr[0] > h-1 || 0 > hr[1] || hr[1] > w-1 {
			continue
		} else if ro[hr[0]][hr[1]] == "#" {
			continue
		} else if !(vstd[hr[0]][hr[1]]) {
			dfs(hr[0], hr[1])
		}
	}
}
