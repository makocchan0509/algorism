package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var h []int
var cost []int
var done []bool
var t1, t2, t3 int

func main() {

	sc.Split(bufio.ScanWords)
	sc.Scan()

	n, _ := strconv.Atoi(sc.Text())

	sc.Scan()
	l, _ := strconv.Atoi(sc.Text())

	h = make([]int, l+1)
	for i, _ := range h {
		h[i] = 0
	}
	for i := 0; i < n; i++ {
		sc.Scan()
		v, _ := strconv.Atoi(sc.Text())

		h[v] = v
	}

	cost = make([]int, l+1)
	for i, _ := range cost {
		cost[i] = 100000 * 10000
	}
	cost[0] = 0

	done = make([]bool, l+1)
	for i, _ := range done {
		done[i] = false
	}

	sc.Scan()
	t1, _ = strconv.Atoi(sc.Text())

	sc.Scan()
	t2, _ = strconv.Atoi(sc.Text())

	sc.Scan()
	t3, _ = strconv.Atoi(sc.Text())

	ans := rec(l)

	for _, v := range []int{l - 3, l - 2, l - 1} {
		if v >= 0 {
			ans = min(ans, cost[v]+t1/2+t2*(2*(l-v)-1)/2)
			//fmt.Println(cost[v] + t1/2 + t2*(2*(l-v)-1)/2)
		}
	}
	fmt.Println(ans)

}

func rec(i int) int {

	if done[i] {
		return cost[i]
	}

	if i == 0 {
		return cost[i]
	}

	cost[i] = min(cost[i], rec(i-1)+t1)

	if i >= 2 {
		cost[i] = min(cost[i], rec(i-2)+t1+t2)
	}

	if i >= 4 {
		cost[i] = min(cost[i], rec(i-4)+t1+(t2*3))
	}

	if h[i] != 0 {
		cost[i] += t3
	}

	done[i] = true

	return cost[i]
}

func min(l int, r int) int {

	if l > r {
		return r
	} else {
		return l
	}
}
