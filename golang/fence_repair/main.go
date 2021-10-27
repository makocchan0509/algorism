package main

import "fmt"

func main() {

	var N int
	var l int
	fmt.Scan(&N)
	L := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scan(&l)
		L[i] = l
	}

	ans := 0

	for N > 1 {
		mil1 := 0
		mil2 := 1

		if L[mil1] > L[mil2] {
			tmp := mil1
			mil1 = mil2
			mil2 = tmp
		}

		for i := 2; i < N; i++ {
			if L[i] < L[mil1] {
				mil2 = mil1
				mil1 = i
			} else if L[i] < L[mil2] {
				mil2 = i
			}
		}

		t := L[mil1] + L[mil2]
		ans += t
		if mil1 == N-1 {
			tmp := mil1
			mil1 = mil2
			mil2 = tmp
		}
		fmt.Println(L, mil1, mil2)
		L[mil1] = t
		L[mil2] = L[N-1]
		N--
		fmt.Println(L)
	}
	fmt.Println(ans)
}
