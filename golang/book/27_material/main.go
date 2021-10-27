package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {

	sc.Split(bufio.ScanWords)

	sc.Scan()
	N, _ := strconv.Atoi(sc.Text())
	N -= 0
	sc.Scan()
	M, _ := strconv.Atoi(sc.Text())

	S := []int{}
	C := []int{}
	S = append(S, 0)
	C = append(C, 0)

	for i := 0; i < M; i++ {
		sc.Scan()
		s := sc.Text()
		sc.Scan()
		c, _ := strconv.Atoi(sc.Text())
		chars := strings.Split(s, "")
		s_val := 0
		for j := 0; j < len(chars); j++ {
			l := N - 1
			if chars[j] == "Y" {
				s_val |= 1 << (l - j)
			}
		}
		S = append(S, s_val)
		C = append(C, c)
	}

	fmt.Println(S)
	fmt.Println(C)

	ALL := 1 << N

	cost := make([][]int, M+1)
	INF := 10000000000

	for i := 0; i < M+1; i++ {
		co := make([]int, ALL)
		for j := 0; j < len(co); j++ {
			co[j] = INF
		}
		cost[i] = co
	}

	cost[0][0] = 0

	for i := 1; i < M+1; i++ {
		for n := 0; n < ALL; n++ {
			cost[i][n] = min(cost[i][n], cost[i-1][n])
			cost[i][n|S[i]] = min(cost[i][n|S[i]], cost[i-1][n]+C[i])
			//fmt.Println(i, n|S[i])
			//viewSlice(cost)
		}
	}

	//viewSlice(cost)

	ans := cost[M][ALL-1]
	if ans == INF {
		ans = -1
	}
	fmt.Println(ans)
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
