package main

import (
	"fmt"
	"strings"
)

func main() {
	var h, w int
	fmt.Scan(&h, &w)
	matrix := make([][]string, h)
	line := ""
	black := "#"

	for i := 0; i < h; i++ {
		fmt.Scan(&line)
		matrix[i] = strings.Split(line, "")
	}

	corner := 0
	for i := 0; i < h-1; i++ {
		for j := 0; j < w-1; j++ {
			count := 0

			if matrix[i][j] == black {
				count++
			}
			if matrix[i+1][j] == black {
				count++
			}
			if matrix[i][j+1] == black {
				count++
			}
			if matrix[i+1][j+1] == black {
				count++
			}
			if count == 3 || count == 1 {
				corner++
			}
		}
	}
	fmt.Println(corner)
}
