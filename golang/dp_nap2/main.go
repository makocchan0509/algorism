package main

import "fmt"

var n, W int
var w, v int
var ca [][]int

func main() {
	fmt.Scan(&n)

	ca = make([][]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&w, &v)
		ca[i] = []int{w, v}
	}
	fmt.Scan(&W)

}
