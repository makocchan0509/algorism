package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	var s int
	flg := false
	for i := 0; i < a; i++ {
		fmt.Scan(&s)

		if s != b {
			if !flg {
				flg = true
				fmt.Printf("%d", s)
			} else {
				fmt.Printf(" %d", s)
			}
		}
	}
	if !flg {
		fmt.Println()
	}

}
