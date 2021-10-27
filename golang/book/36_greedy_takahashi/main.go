package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type cuisine struct {
	sum int
	tk  int
	ak  int
}

func main() {

	sc.Split(bufio.ScanWords)
	sc.Scan()

	N, _ := strconv.Atoi(sc.Text())

	C := make([]cuisine, N)

	for i := 0; i < N; i++ {

		sc.Scan()
		a, _ := strconv.Atoi(sc.Text())
		sc.Scan()
		b, _ := strconv.Atoi(sc.Text())

		var c cuisine
		c.sum = a + b
		c.tk = a
		c.ak = b

		C[i] = c
	}

	sort.Slice(C, func(i, j int) bool {
		return C[i].sum > C[j].sum
	})

	//fmt.Println(C)
	ans := 0

	for i := 0; i < N; i++ {
		c := C[i]
		if i%2 == 0 {
			ans += c.tk
		} else {
			ans -= c.ak
		}
	}
	fmt.Println(ans)
}
