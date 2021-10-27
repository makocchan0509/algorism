package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Scan(&str)

	ca := strings.Count(str, "a")
	cb := strings.Count(str, "b")
	cc := strings.Count(str, "c")

	if ca < cb {
		if cb < cc {
			fmt.Println("c")
		} else {
			fmt.Println("b")
		}
	} else if ca < cc {
		fmt.Println("c")
	} else {
		fmt.Println("a")
	}

}
