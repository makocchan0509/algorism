package main

import (
	"fmt"
	"strconv"
)

func main() {

	var n int
	fmt.Scan(&n)
	z := 0
	for i := 1; i <= 555555; i++ {

		result := true

		s := strconv.Itoa(i)
		var t string

		for i, c := range s {

			if i == 0 {
				t = string(c)
				continue
			}
			if string(c) != t {
				result = false
			}
		}

		if result {
			z++
		}
		if z == n {
			fmt.Println(i)
			break
		}
	}
}
