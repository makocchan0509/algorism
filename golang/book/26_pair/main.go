package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {

	sc.Split(bufio.ScanWords)

	sc.Scan()
	N, _ := strconv.Atoi(sc.Text())

	A := make([][]int, N-1)

	for i := 0; i < N-1; i++ {

		aa := make([]int, N)

		for y := 0; y < len(aa); y++ {
			aa[y] = 0
		}
		for j := i + 1; j < len(aa); j++ {
			sc.Scan()
			v, _ := strconv.Atoi(sc.Text())
			aa[j] = v
		}
		A[i] = aa
	}

	//fmt.Println(A)

	ALL := 1 << N
	//fmt.Println(ALL)

	happy := make([]int, ALL)
	for i := 0; i < len(happy); i++ {
		happy[i] = 0
	}

	for n := 0; n < ALL; n++ {
		for i := 0; i < N; i++ {
			for j := i + 1; j < N; j++ {
				if has_bit(n, i) && has_bit(n, j) {
					happy[n] += A[i][j]
				}
			}
		}
	}

	fmt.Println(happy)

	ans := -10 * 10000000

	for n1 := 0; n1 < ALL; n1++ {
		for n2 := 0; n2 < ALL; n2++ {
			if n1&n2 > 0 {
				continue
			}
			n3 := ALL - 1 - (n1 | n2)
			ans = max(ans, happy[n1]+happy[n2]+happy[n3])
		}
	}

	fmt.Println(ans)
}

func has_bit(n int, i int) bool {
	//fmt.Printf("n: %v i: %v has_bit: %v \n", n, i, n&(1<<i))
	return (n&(1<<i) > 0)
}

func max(l int, r int) int {

	if l < r {
		return r
	} else {
		return l
	}
}
