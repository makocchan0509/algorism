package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var cost []int
var done []bool
var h []float64

var sc = bufio.NewScanner((os.Stdin))

func main() {

	sc.Split(bufio.ScanWords)

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	h = make([]float64, n)

	for i := 0; i < n; i++ {
		sc.Scan()
		v, _ := strconv.Atoi(sc.Text())
		h[i] = float64(v)
	}

	cost = make([]int, n)
	for i, _ := range cost {
		cost[i] = 0
	}

	done = make([]bool, n)
	for i, _ := range done {
		done[i] = false
	}

	fmt.Println(rec(n - 1))

}

func rec(i int) int {

	if done[i] {
		return cost[i]
	}

	if i == 0 {
		cost[i] = 0
	} else if i == 1 {
		cost[i] = rec(0) + int(math.Abs(h[i-1]-h[i]))
	} else {
		cost[i] = min(rec(i-1)+int(math.Abs(h[i-1]-h[i])), rec(i-2)+int(math.Abs(h[i-2]-h[i])))
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
