package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner((os.Stdin))

func main() {
	sc.Split(bufio.ScanWords)

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	h := make([]float64, n)

	for i := 0; i < n; i++ {
		sc.Scan()
		v, _ := strconv.Atoi(sc.Text())
		h[i] = float64(v)
	}

	cost := make([]int, n)
	for i, _ := range cost {
		cost[i] = 0
	}

	cost[0] = 0
	cost[1] = int(math.Abs(h[0] - h[1]))

	for i := 2; i < n; i++ {
		cost[i] = min(cost[i-1]+int(math.Abs(h[i-1]-h[i])), cost[i-2]+int(math.Abs(h[i-2]-h[i])))
	}

	fmt.Println(cost[n-1])

}

func min(l int, r int) int {

	if l > r {
		return r
	} else {
		return l
	}
}
