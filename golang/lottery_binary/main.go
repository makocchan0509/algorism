package main

import (
	"fmt"
	"sort"
	"time"
)

var n, m, kk int
var k = make([]int, n)

func binary_search(x int) bool {
	l := 0
	r := n
	var i int
	for r-l >= 1 {
		i = (l + r) / 2
		if k[i] == x {
			return true
		} else if k[i] < x {
			l = i + 1
		} else {
			r = i
		}
	}
	return false
}

func main() {

	fmt.Scan(&n, &m)

	for i := 0; i < n; i++ {
		fmt.Scan(&kk)
		k[i] = kk
	}
	fmt.Println(k)
	sort.Sort(sort.IntSlice(k))
	fmt.Println(k)
	start := time.Now()
	ans := false
	for a := 0; a < n; a++ {
		for b := 0; b < n; b++ {
			for c := 0; c < n; c++ {
				if binary_search(m - k[a] - k[b] - k[c]) {
					ans = true
				}
			}
		}
	}
	fmt.Println(ans)
	end := time.Now()
	fmt.Printf("%fsec\n", (end.Sub(start)).Seconds())
}
