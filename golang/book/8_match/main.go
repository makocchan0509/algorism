package main

import (
	"fmt"
	"strings"
)

func main() {

	t := "abcdefghijklmnopqrstuvwxyz."

	var s string
	fmt.Scan(&s)
	var result []string

	for _, c := range t {
		if match(s, string(c)) {
			result = append(result, string(c))
		}
	}
	fmt.Println(result)

	if len(s) >= 2 {
		for _, c1 := range t {
			for _, c2 := range t {
				str := string(c1) + string(c2)
				if match(s, str) {
					result = append(result, str)
				}
			}
		}
		fmt.Println(result)
	}

	if len(s) >= 3 {
		for _, c1 := range t {
			for _, c2 := range t {
				for _, c3 := range t {
					str := string(c1) + string(c2) + string(c3)
					if match(s, str) {
						result = append(result, str)
					}
				}
			}
		}
		fmt.Println(result)
	}

	fmt.Println(len(result))

}

func match(s string, t string) bool {

	scnt := len(s)
	tcnt := len(t)

	arrS := strings.Split(s, "")
	arrT := strings.Split(t, "")

	for i := 0; i <= scnt-tcnt; i++ {
		result := true

		for j := 0; j < tcnt; j++ {
			if arrS[i+j] != arrT[j] && arrT[j] != "." {
				result = false
			}
		}
		if result {
			return true
		}
	}

	return false
}
