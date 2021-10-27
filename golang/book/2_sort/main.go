package main

import "fmt"

func main() {
	var X int

	org := make([]int, 6)

	for i := 0; i < len(org); i++ {
		fmt.Scan(&X)
		org[i] = X
	}

	result := sort(org)
	fmt.Println(result)
	fmt.Println(result[2])
}

func sort(org []int) []int {

	//result := make([]int, len(org))

	for i, _ := range org {
		if i == 0 {
			continue
		}
		if org[i] < org[i-1] {
			tmp := org[i-1]
			org[i-1] = org[i]
			org[i] = tmp
			z := i - 1
			for z > 0 {
				if org[z] < org[z-1] {
					tmp := org[z-1]
					org[z-1] = org[z]
					org[z] = tmp
				}
				z--
			}
		}
	}
	return org
}
