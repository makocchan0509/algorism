package main

import (
	"fmt"
	"math"
)

var R, B, x, y int

func main() {

	fmt.Scan(&R, &B)
	fmt.Scan(&x, &y)

	//fmt.Println(R,B,x,y)

	ok := 0
	ng := int(math.Pow(10, 18)) + 1

	for int(math.Abs(float64(ok-ng))) > 1 {
		mid := (ok + ng) / 2
		if check(mid) {
			ok = mid
		} else {
			ng = mid
		}
	}

	fmt.Println(ok)
}

func check(X int) bool {
	r := R - X
	b := B - X

	if r < 0 || b < 0 {
		return false
	}

	num := r/(x-1) + b/(y-1)

	return num >= X

}
