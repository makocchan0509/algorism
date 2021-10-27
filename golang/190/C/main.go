package main

import "fmt"

func main() {
	var N, M, K int
	fmt.Scan(&N, &M)

	A := make([]int, M)
	B := make([]int, M)

	for i := 0; i < M; i++ {
		fmt.Scan(&A[i], &B[i])
	}

	fmt.Scan(&K)
	C := make([]int, K)
	D := make([]int, K)
	for i := 0; i < K; i++ {
		fmt.Scan(&C[i], &D[i])
	}

	ans := 0
	for i := 0; i < 1<<K; i++ {
		dish := make([]bool, N)
		for k := 0; k < K; k++ {
			//fmt.Println(i & (1 << k))
			if i&(1<<k) > 0 {
				dish[C[k]-1] = true
			} else {
				dish[D[k]-1] = true
			}
		}
		sum := 0
		for m := 0; m < M; m++ {
			if dish[A[m]-1] && dish[B[i]-1] {
				sum++
			}
		}
		if ans < sum {
			ans = sum
		}
	}
	fmt.Println(ans)

}
