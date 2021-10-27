package main

import (
	"fmt"
	"math"
)

func main() {

	var A, R, N int

	fmt.Scan(&A, &R, &N)
	ans := solve(A, R, N)

	if ans == 0 {
		fmt.Println("large")
	} else {
		fmt.Println(ans)
	}
}

func solve(A int, R int, N int) int {

	if R == 1 {
		return A
	}
	for i := 0; i < N-1; i++ {
		A *= R

		if A > int(math.Pow(10, 9)) {
			return 0
		}
	}
	return A
}
