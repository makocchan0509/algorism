package main

import (
	"fmt"
	"sort"
)

type Work struct {
	first, second int
}

type Tasks []Work

func (w Tasks) Len() int {
	return len(w)
}
func (w Tasks) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}
func (w Tasks) Less(i, j int) bool {
	return w[i].first < w[j].first
}

func main() {
	var n int
	fmt.Scan(&n)
	var ss, tt int
	var s, t []int
	s = make([]int, n)
	t = make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&ss)
		s[i] = ss
	}
	for i := 0; i < n; i++ {
		fmt.Scan(&tt)
		t[i] = tt
	}

	tasks := make(Tasks, n)

	for i := 0; i < n; i++ {
		tasks[i].first = t[i]
		tasks[i].second = s[i]
	}
	sort.Sort(tasks)

	ans := 0
	l := 0

	for i := 0; i < n; i++ {
		if l < tasks[i].second {
			ans++
			l = tasks[i].first
		}
	}
	fmt.Println(ans)

}
