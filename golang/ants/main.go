package main

import (
	"fmt"
	"math"
)

func main() {
	var l, n, a int
	fmt.Scan(&l, &n)

	minT, maxT := 0, 0
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		minT = int(math.Max(float64(minT), math.Min(float64(a), float64(l-a))))
		maxT = int(math.Max(float64(maxT), math.Max(float64(a), float64(l-a))))
		fmt.Println(minT, maxT)
	}
	fmt.Println(minT, maxT)
}
