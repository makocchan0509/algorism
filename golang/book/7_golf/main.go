package main

import "fmt"

func main() {

	var k int
	var a, b int

	fmt.Scan(&k)
	fmt.Scan(&a, &b)

	result := false

	for a <= b {
		if a%k == 0 {
			result = true
			break
		}
		a++
	}
	if result {
		fmt.Println("OK")
	} else {
		fmt.Println("NG")
	}
}
