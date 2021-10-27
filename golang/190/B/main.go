package main

import "fmt"

func main() {

	var N, S, D int
	fmt.Scan(&N, &S, &D)

	var X, Y int
	for i := 0; i < N; i++ {
		fmt.Scan(&X, &Y)
		if X < S && Y > D {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}
