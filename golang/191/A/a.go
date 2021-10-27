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

	var slice [4]int

	if sc.Scan() {
		strs := strings.Split(sc.Text(), " ")
		for i, v := range strs {
			slice[i], _ = strconv.Atoi(v)
		}
	}
	if slice[0]*slice[1] > slice[3] || slice[0]*slice[2] < slice[3] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
