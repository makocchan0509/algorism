package main

import (
	"fmt"
	"math"
)

func main() {
	var C []int
	var Y = [6]int{1, 5, 10, 50, 100, 500}
	var A, c int

	C = make([]int, 6)

	for i := 0; i < 6; i++ {
		fmt.Scan(&c)
		C[i] = c
	}
	fmt.Scan(&A)
	ans := 0

	for i := 5; 0 <= i; i-- {
		t := int(math.Min(float64(A/Y[i]), float64(C[i])))
		A -= t * Y[i]
		ans += t
	}
	fmt.Println(ans)
}
