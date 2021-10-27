package main

import (
	"fmt"
	"strings"
)

func main() {
	var N int
	var S string
	fmt.Scan(&N)
	fmt.Scan(&S)

	s := strings.Split(S, "")

	min_turn := 300000

	sum_W := []int{0}

	for i := 0; i < N; i++ {

		if s[i] == "W" {
			sum_W = append(sum_W, sum_W[i]+1)
		} else {
			sum_W = append(sum_W, sum_W[i])
		}
	}

	//WEEWW
	//[0 1 1 1 2 3]

	fmt.Println(sum_W)

	for i := 0; i < N; i++ {

		turn := 0
		w := sum_W[i]
		e := (N - 1 - i) - (sum_W[N] - sum_W[i+1])

		turn = w + e

		min_turn = min(min_turn, turn)
	}

	fmt.Println(min_turn)
}

// This is long time process method O(N^2)
/* func main() {

	var N int
	var S string
	fmt.Scan(&N)
	fmt.Scan(&S)

	s := strings.Split(S, "")

	min_turn := 300000

	// i = leader
	for i := 0; i < N; i++ {

		turn := 0

		// p = not leader
		for p := 0; p < N; p++ {

			if p == i {
				continue
			}

			if p < i && s[p] == "W" {
				turn += 1
			}

			if p > i && s[p] == "E" {
				turn += 1
			}

			min_turn = min(min_turn, turn)
		}
	}

	fmt.Println(min_turn)
}
*/

func min(l int, r int) int {

	if l > r {
		return r
	} else {
		return l
	}
}
