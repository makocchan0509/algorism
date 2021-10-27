package main

import "fmt"

func main() {

	var N, X int
	fmt.Scan(&N, &X)

	var A []int
	var B []int

	for i := 0; i < N; i++ {
		var w int
		fmt.Scan(&w)

		if i%2 == 0 {
			A = append(A, w)
		} else {
			B = append(B, w)
		}
	}

	//fmt.Println(A)
	//fmt.Println(B)
	m := map[int]int{}

	for n := 0; n < (1 << len(B)); n++ {
		s := 0
		for i := 0; i < N; i++ {
			if has_bit(n, i) {
				s += B[i]
			}
		}
		m[s] += 1
	}
	//fmt.Println(m)
	//fmt.Println(1 << len(B))

	ans := 0

	for n := 0; n < (1 << len(A)); n++ {
		s := 0
		for i := 0; i < N; i++ {
			if has_bit(n, i) {
				s += A[i]
			}
		}
		ans += m[X-s]
	}
	fmt.Println(ans)
}

func has_bit(n int, i int) bool {
	return (n&(1<<i) > 0)
}
