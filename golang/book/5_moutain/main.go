package main

import (
	"fmt"
	"strings"
)

func main() {

	var n int
	var s string

	fmt.Scan(&n)

	m := make([][]string, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&s)
		m[i] = strings.Split(s, "")
	}

	i := n - 1

	for i > -1 {
		if i == n-1 {
			i--
			continue
		}
		for j, v := range m[i] {
			result := false
			if v == "#" {
				if j > 0 && m[i+1][j-1] == "X" {
					result = true
				}
				if m[i+1][j] == "X" {
					result = true
				}
				if j < n*2-1 && m[i+1][j+1] == "X" {
					result = true
				}
			}
			if result {
				m[i][j] = "X"
			}
		}
		i--
	}

	for i := 0; i < len(m); i++ {
		fmt.Println(strings.Join(m[i], ""))
	}
}
