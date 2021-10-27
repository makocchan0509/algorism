package main

import "fmt"

func main() {
	var A, B, C int
	fmt.Scan(&A, &B, &C)
	var ahead, rear int

	if C == 0 {
		ahead = A
		rear = B
		A = A - 1
	} else {
		ahead = B
		rear = A
		B = B - 1
	}

	if A < B {
		fmt.Println("Aoki")
		return
	} else if A > B {
		fmt.Println("Takahashi")
		return
	}

	for true {

		ahead--
		if ahead == 0 && C == 0 {
			fmt.Println("Aoki")
			return
		} else if ahead == 0 && C == 1 {
			fmt.Println("Takahashi")
			return
		}
		rear--
		if rear == 0 && C == 0 {
			fmt.Println("Takahashi")
			return
		} else if rear == 0 && C == 1 {
			fmt.Println("Aoki")
			return
		}
	}

}
