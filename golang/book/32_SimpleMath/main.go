package main

import "fmt"

func main() {
	var A, B, C int

	fmt.Scan(&A, &B, &C)

	// 1 + 2 + ... + C = 1/2C(C+1)
	D := A * (A + 1) / 2
	E := B * (B + 1) / 2
	F := C * (C + 1) / 2

	ans := (D * E * F) % 998244353

	fmt.Println(ans)
}
