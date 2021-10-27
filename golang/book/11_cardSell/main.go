package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Split(bufio.ScanWords)

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	c := make([]int, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		cc, _ := strconv.Atoi(sc.Text())
		c[i] = cc
	}

	sc.Scan()
	q, _ := strconv.Atoi(sc.Text())

	sell := 0
	z := 0
	s := 0

	min_s := 1000000000
	min_z := 1000000000

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			min_s = min(min_s, c[i])
		} else {
			min_z = min(min_z, c[i])
		}
	}

	for i := 0; i < q; i++ {
		sc.Scan()
		s1, _ := strconv.Atoi(sc.Text())

		sc.Scan()
		s2, _ := strconv.Atoi(sc.Text())

		if s1 == 1 {
			s2 = s2 - 1
			sc.Scan()
			s3, _ := strconv.Atoi(sc.Text())
			cardx := 0
			if s2%2 == 0 {
				cardx = c[s2] - s - z
			} else {
				cardx = c[s2] - z
			}

			if cardx >= s3 {
				c[s2] -= s3
				sell += s3

				if s2%2 == 0 {
					min_s = min(min_s, c[s2])
				} else {
					min_z = min(min_z, c[s2])
				}
			}
		} else if s1 == 2 {
			if min_s-s-z >= s2 {
				s += s2
			}
		} else {
			if min(min_s-s-z, min_z-z) >= s2 {
				z += s2
			}
		}
	}

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			sell += s
		}
	}

	sell += z * n
	fmt.Println(sell)

}

func min(l int, r int) int {
	if l > r {
		return r
	} else {
		return l
	}
}
