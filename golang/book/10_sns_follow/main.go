package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {

	var n, q int
	fmt.Scan(&n, &q)

	u := make([][]string, n)

	for i := 0; i < n; i++ {
		userL := make([]string, n)
		for y := 0; y < n; y++ {
			userL[y] = "N"

		}
		u[i] = userL
	}

	sc.Split(bufio.ScanWords)

	l := make([][]int, q)
	for i := 0; i < q; i++ {
		l[i] = makeArrIntExt(3)
	}

	for i := 0; i < len(l); i++ {

		s1 := l[i][0]
		s2 := l[i][1]

		if s1 == 1 {
			s3 := l[i][2]
			u[s2][s3] = "Y"
		} else if s1 == 2 {
			for i := 0; i < len(u); i++ {
				if i != s2 && u[i][s2] == "Y" {
					u[s2][i] = "Y"
				}
			}
		} else {
			//outArr(u)
			for i, v := range u[s2] {
				if v == "Y" {
					for ii, vv := range u[i] {
						if vv == "Y" {
							u[s2][ii] = "V"
						}
					}
				}
			}
			for i, v := range u[s2] {
				if v == "V" {
					u[s2][i] = "Y"
				}
			}
		}
	}

	outArr(u)

}

func makeArrIntExt(c int) []int {

	arr := make([]int, c)

	sc.Scan()
	s1, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	s2, _ := strconv.Atoi(sc.Text())
	s2 = s2 - 1

	arr[0] = s1
	arr[1] = s2

	if s1 == 1 {
		sc.Scan()
		s3, _ := strconv.Atoi(sc.Text())
		s3 = s3 - 1
		arr[2] = s3
	}

	return arr
}

func outArr(arr [][]string) {

	for i := 0; i < len(arr); i++ {
		result := strings.Join(arr[i], "")
		fmt.Println(result)
	}

}
