package main

import "fmt"

var n, k, aa int
var a []int

func main() {
	fmt.Scan(&n)
	a = make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&aa)
		a[i] = aa
	}
	fmt.Scan(&k)

	if dfs(0, 0) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func dfs(i int, sum int) bool {
	if i == n {
		return sum == k
	}

	if dfs(i+1, sum) {
		return true
	}
	if dfs(i+1, sum+a[i]) {
		return true
	}
	return false
}
