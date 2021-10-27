package main

import (
	"fmt"
	"time"
)

func main() {

	var n, m, kk int
	fmt.Scan(&n, &m)
	k := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&kk)
		k[i] = kk
	}
	start := time.Now()
	ans := false
	for a := 0; a < n; a++ {
		for b := 0; b < n; b++ {
			for c := 0; c < n; c++ {
				for d := 0; d < n; d++ {
					if k[a]+k[b]+k[c]+k[d] == m {
						ans = true
					}
				}
			}
		}
	}
	fmt.Println(ans)
	end := time.Now()
	fmt.Printf("%fsec\n", (end.Sub(start)).Seconds())
}
