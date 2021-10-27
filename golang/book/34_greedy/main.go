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

	X := make([][]int, N)

	for i := 0; i < N; i++ {
		sc.Scan()
		a, _ := strconv.Atoi(sc.Text())
		sc.Scan()
		b, _ := strconv.Atoi(sc.Text())

		X[a-1] = append(X[a-1], b)
	}

	//fmt.Println(X)
	//[[3] [2 4] []]

	cnt := makeIntSlice(101, 0)
	ans := 0

	for d := 0; d < N; d++ {
		for _, b := range X[d] {
			cnt[b] += 1
		}
		//fmt.Println(cnt)
		for bb := 100; bb > -1; bb-- {
			if cnt[bb] > 0 {
				ans += bb
				cnt[bb] -= 1
				break
			}
		}
		fmt.Println(ans)
	}
}

/*
	Create int slice and set default value
    x: slice num
    y: default value
*/
func makeIntSlice(x int, y int) []int {

	s := make([]int, x)

	for i := 0; i < x; i++ {
		s[i] = 0
	}

	return s
}
