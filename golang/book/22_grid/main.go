package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

type place struct {
	i int
	j int
}

func main() {
	sc.Split(bufio.ScanWords)

	sc.Scan()
	N, _ := strconv.Atoi(sc.Text())

	sc.Scan()
	M, _ := strconv.Atoi(sc.Text())
	M -= 0
	an := make([][]int, N)

	for i := 0; i < N; i++ {
		sc.Scan()
		nm := sc.Text()

		l := strings.Split(nm, "")

		an[i] = strToInt(l)
	}

	//fmt.Println(an)

	g := make([][]place, 11)
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			n := an[i][j]
			var p place
			p.i = i
			p.j = j
			//nn := []int{i, j}
			g[n] = append(g[n], p)
		}
	}

	fmt.Println(g)

	cost := make([][]int, N)

	for i := 0; i < N; i++ {
		c := make([]int, M)
		for j := 0; j < M; j++ {
			c[j] = int(math.Pow(10, 10))
		}
		cost[i] = c
	}

	//fmt.Println(cost)
	sp := g[0][0]
	//fmt.Println(sp)
	cost[sp.i][sp.j] = 0

	for n := 1; n < len(g); n++ {
		for _, p := range g[n] {
			for _, pp := range g[n-1] {
				cost[p.i][p.j] = min(cost[p.i][p.j], cost[pp.i][pp.j]+int(math.Abs(float64(p.i-pp.i)))+int(math.Abs(float64(p.j-pp.j))))
			}
		}
	}
	fmt.Println(cost)

	gp := g[10][0]
	ans := cost[gp.i][gp.j]
	if ans == int(math.Pow(10, 10)) {
		ans = -1
	}
	fmt.Println(ans)

}

func strToInt(s []string) []int {

	result := make([]int, len(s))
	for i, v := range s {
		n := 0
		if v == "S" {
			n = 0
		} else if v == "G" {
			n = 10
		} else {
			n, _ = strconv.Atoi(v)
		}
		result[i] = n
	}

	return result
}

func min(l int, r int) int {

	if l > r {
		return r
	} else {
		return l
	}
}
