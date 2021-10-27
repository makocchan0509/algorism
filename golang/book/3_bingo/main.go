package main

import "fmt"

func main() {

	var j1, j2, j3 int
	var n int
	var b int
	bowls := map[int]bool{}

	card := make([][]int, 3)
	cardR := make([][]bool, 3)

	for i := 0; i < len(cardR); i++ {
		cardR[i] = []bool{false, false, false}
	}

	for i := 0; i < len(card); i++ {
		fmt.Scan(&j1, &j2, &j3)
		card[i] = []int{j1, j2, j3}
	}
	fmt.Scan(&n)

	for z := 0; z < n; z++ {
		fmt.Scan(&b)
		bowls[b] = true
	}

	for i, arr := range card {
		for j, v := range arr {
			if bowls[v] {
				cardR[i][j] = true
			}
		}
	}
	fmt.Println(cardR)
	result := false
	for i := 0; i < len(cardR); i++ {
		if cardR[i][0] && cardR[i][1] && cardR[i][2] {
			result = true
			break
		}
		if cardR[0][i] && cardR[1][i] && cardR[2][i] {
			result = true
			break
		}
	}

	if !result {
		if cardR[0][0] && cardR[1][1] && cardR[2][2] {
			result = true
		} else if cardR[0][2] && cardR[1][1] && cardR[2][0] {
			result = true
		}
	}
	fmt.Println(result)

}
