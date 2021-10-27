package main

import "fmt"

func main() {

	var N int
	fmt.Scan(&N)

	c := make([]int, N+1)
	c[0] = 0
	for i := 1; i < N+1; i++ {
		var p int
		fmt.Scan(&p)
		c[i] = p
	}

	P := sum(c)

	result := make([][]bool, N+1)

	for i := 0; i < N+1; i++ {
		line := make([]bool, P+1)
		for i, _ := range line {
			line[i] = false
		}
		result[i] = line
	}

	result[0][0] = true

	//	fmt.Println(c)
	//	fmt.Println(P)
	//	fmt.Println(result)

	for i := 1; i < N+1; i++ {
		for s := 0; s < P+1; s++ {

			if result[i-1][s] {
				result[i][s] = true
			}
			if s >= c[i] && result[i-1][s-c[i]] {
				result[i][s] = true
			}
		}
	}

	ans := 0

	for s := 0; s < P+1; s++ {
		if result[N][s] {
			ans += 1
		}
	}

	fmt.Println(ans)

}

func sum(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}
