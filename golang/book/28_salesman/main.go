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
	A := make([][]int, N)

	for i := 0; i < N; i++ {
		line := make([]int, N)
		for z := 0; z < N; z++ {
			sc.Scan()
			j, _ := strconv.Atoi(sc.Text())
			line[z] = j
		}
		//fmt.Println(line)
		A[i] = line
	}

	//fmt.Println(A)

	ALL := 1 << N
	fmt.Println(ALL)
	cost := make([][]int, ALL)

	for i := 0; i < ALL; i++ {
		line := make([]int, N)
		for z := 0; z < N; z++ {
			line[z] = 10 * 1000000000
		}
		cost[i] = line
	}

	//fmt.Println(cost)
	cost[0][0] = 0

	n := 0
	i := 0
	j := 0
	for n < ALL {
		for i < N {
			for j < N {
				//fmt.Println(n, i, j)
				if has_bit(n, j) || i == j {
					j += 1
					continue
				}
				cost[n|(1<<j)][j] = min(cost[n|(1<<j)][j], cost[n][i]+A[i][j])
				j += 1
			}
			i += 1
			j = 0
		}
		n += 1
		i = 0
		j = 0
	}
	viewSlice(cost)
	fmt.Println(cost[ALL-1][0])
}

func has_bit(n int, i int) bool {
	return (n&(1<<i) > 0)
}

func min(l int, r int) int {

	if l > r {
		return r
	} else {
		return l
	}
}

func viewSlice(target [][]int) {

	for _, v := range target {
		fmt.Println(v)
	}
}
