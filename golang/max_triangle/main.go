package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	var l int
	for i := 0; i < n; i++ {
		fmt.Scan(&l)
		a[i] = l
	}
	ans := 0

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for x := j + 1; x < n; x++ {
				len := a[i] + a[j] + a[x]
				max := int(math.Max(float64(a[i]), math.Max(float64(a[j]), float64(a[x]))))
				rest := len - max

				if rest > max {
					ans = int(math.Max(float64(len), float64(ans)))
				}
			}
		}
	}
	fmt.Println(ans)
}
