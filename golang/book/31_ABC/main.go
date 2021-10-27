package main

import "fmt"

func main() {
	var N int

	fmt.Scan(&N)

	ans := 0

	for i := 1; i < N; i++ {

		b_count := (N - 1) / i
		fmt.Println(b_count)
		ans += b_count
	}

	fmt.Println(ans)
}
