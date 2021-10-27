package main

import "fmt"

var N int
var total int

func main() {

	fmt.Scan(&N)

	dfs(0, false, false, false)

	print(total)

}

func dfs(n int, use3 bool, use5 bool, use7 bool) {

	if n > N {
		return
	}

	if use3 && use5 && use7 {
		total += 1
	}

	dfs(10*n+3, true, use5, use7)
	dfs(10*n+5, use3, true, use7)
	dfs(10*n+7, use3, use5, true)
}
