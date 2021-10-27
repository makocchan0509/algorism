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
	A,_ := strconv.Atoi(sc.Text())

	sc.Scan()
	B,_ := strconv.Atoi(sc.Text())

	sc.Scan()
	X,_ := strconv.Atoi(sc.Text())

	ok := 0
	ng := int(math.Pow(float64(10),float64(9))) + 1

	for int(math.Abs(float64(ok-ng))) > 1{
		mid := (ok+ng) / 2
		d := len(strconv.Itoa(mid))
		price := A * mid + B * d

		if price <= X {
			ok = mid
		}else{
			ng = mid
		}

		fmt.Println("ok",ok)
		fmt.Println("ng",ng)
		fmt.Println("mid",mid)
	}

	fmt.Println(ok)

}
