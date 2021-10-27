package main

import "fmt"

func main() {
	var n int
	var s string
	fmt.Scan(&n, &s)

	a := 0
	b := n - 1
	c := ""

	for a <= b {
		left := false
		for i := 0; a+i <= b; i++ {
			if s[a+i] < s[b-i] {
				left = true
				break
			} else if s[a+i] > s[b-i] {
				left = false
				break
			}
		}
		if left {
			a++
			c = c + s[a-1:a]
		} else {
			b--
			c = c + s[b+1:b+2]
		}
	}
	fmt.Println(c)
}
