package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var X, Y, R float64
	fmt.Scan(&X, &Y, &R)

	var x, y, r int
	// x = int64(X * 1e4)
	// y = int64(Y * 1e4)
	// r = int64(R * 1e4)

	x = int(math.Round(X * 10000.0))
	y = int(math.Round(Y * 10000.0))
	r = int(math.Round(R * 10000.0))

	count := 0
	ll := int(math.Round(math.Ceil(float64(x-r)/1e4) * 1e4))
	rr := int(math.Round(math.Floor(float64(x+r)/1e4) * 1e4))

	sqrt := func(v int) int {
		return sort.Search(3e9, func(x int) bool {
			return v < x*x
		})
	}

	for i := ll; i <= rr; i += 10000 {
		t := sqrt((r*r)-((x-i)*(x-i))) - 1
		y1 := int(math.Ceil(float64(y-t) / 1e4))
		y2 := int(math.Floor(float64(y+t) / 1e4))
		count += y2 - y1 + 1
	}

	fmt.Println(count)

}
