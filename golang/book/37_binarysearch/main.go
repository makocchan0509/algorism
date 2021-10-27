package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {

	sc.Split(bufio.ScanWords)
	sc.Scan()

	N,_ := strconv.Atoi(sc.Text())

	sc.Scan()
	K,_ := strconv.Atoi(sc.Text())

	A := make([]int,N)

	for i:=0;i<N;i++{
		sc.Scan()
		a,_ := strconv.Atoi(sc.Text())
		A[i] = a
	}

	ok := N
	ng := -1

	for int(math.Abs(float64(ok-ng))) > 1{
		mid := (ok+ng) / 2

		if A[mid] >= K{
			ok = mid
		}else{
			ng = mid
		}

		fmt.Println("ok",ok)
		fmt.Println("ng",ng)
		fmt.Println("mid",mid)
	}

	if ok == N{
		fmt.Println(-1)
	}else{
		fmt.Println(ok)
	}

}
