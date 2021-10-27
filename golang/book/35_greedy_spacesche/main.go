package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type task struct {
	start int
	end   int
}

func main() {

	sc.Split(bufio.ScanWords)
	sc.Scan()

	N, _ := strconv.Atoi(sc.Text())

	tasks := make([]task,N)

	for i := 0; i < N; i++ {
		sc.Scan()
		a,_ := strconv.Atoi(sc.Text())
		sc.Scan()
		b,_ := strconv.Atoi(sc.Text())

		var t task
		t.start = a
		t.end = b

		tasks[i] = t
	}

	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].end < tasks[j].end
	})

	ans := 0
	last := 0

	for _,t := range tasks {
		if t.start > last{
			ans += 1
			last = t.end
		}
	}
	fmt.Println(ans)
}
