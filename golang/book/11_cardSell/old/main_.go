package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main_() {
	sc.Split(bufio.ScanWords)

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	cm := map[int]int{}
	for i := 0; i < n; i++ {
		sc.Scan()
		c, _ := strconv.Atoi(sc.Text())
		cm[i] = c
	}

	//fmt.Println(cm)

	sc.Scan()
	q, _ := strconv.Atoi(sc.Text())

	count := 0
	for i := 0; i < q; i++ {

		sc.Scan()
		s1, _ := strconv.Atoi(sc.Text())
		sc.Scan()
		s2, _ := strconv.Atoi(sc.Text())
		fmt.Println(s1, s2)
		if s1 == 1 {
			s2 = s2 - 1
			sc.Scan()
			s3, _ := strconv.Atoi(sc.Text())
			if cm[s2] >= s3 {
				count = count + s3
				cm[s2] = cm[s2] - s3
			}
		} else if s1 == 2 {
			b := 0
			sell := true
			t := map[int]int{}
			for i := 0; i < len(cm); i++ {

				if i%2 == 0 {
					if cm[i] >= s2 {
						b = b + s2
						t[i] = s2
					} else {
						sell = false
						break
					}
				}
			}

			if sell {
				count = count + b
				for i := 0; i < len(cm); i++ {
					cm[i] = cm[i] - t[i]
				}
			}

		} else {
			c := 0
			sell := true
			t := map[int]int{}
			for i := 0; i < len(cm); i++ {

				if cm[i] >= s2 {
					c = c + s2
					t[i] = s2
				} else {
					sell = false
					break
				}
			}

			if sell {
				count = count + c
				for i := 0; i < len(cm); i++ {
					cm[i] = cm[i] - t[i]
				}
			}
		}

		//fmt.Println(cm)
	}

	fmt.Println(count)
}
