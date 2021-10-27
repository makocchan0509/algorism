package main

import "fmt"

func main() {
	var n, r, x int
	fmt.Scan(&n, &r)

	X := make([]int, n)

	for j := 0; j < n; j++ {
		fmt.Scan(&x)
		X[j] = x
	}
	i := 0
	ans := 0

	for i < n {
		s := X[i]
		i++
		for i < n && X[i] <= s+r {
			i++
		}
		p := X[i-1]

		for i < n && X[i] <= p+r {
			i++
		}
		ans++
	}
	fmt.Println(ans)
}
