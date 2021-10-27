package main

import "fmt"

func main() {

	var c1, c2, c3 int
	grid := make([][]int, 3)

	for i := 0; i < len(grid); i++ {
		fmt.Scan(&c1, &c2, &c3)
		grid[i] = []int{c1, c2, c3}
	}

	result := true

	if grid[0][0]-grid[0][1] != grid[1][0]-grid[1][1] || grid[1][0]-grid[1][1] != grid[2][0]-grid[2][1] {
		fmt.Println("1")
		result = false
	}

	if grid[0][1]-grid[0][2] != grid[1][1]-grid[1][2] || grid[1][1]-grid[1][2] != grid[2][1]-grid[2][2] {
		fmt.Println("2")
		result = false
	}

	if grid[0][0]-grid[1][0] != grid[0][1]-grid[1][1] || grid[0][1]-grid[1][1] != grid[0][2]-grid[1][2] {
		fmt.Println("3")
		result = false
	}

	if grid[1][0]-grid[2][0] != grid[1][1]-grid[2][1] || grid[1][1]-grid[2][1] != grid[1][2]-grid[2][2] {
		fmt.Println("4")
		result = false
	}

	if result {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
