package main

import (
	"container/heap"
	"fmt"
)

type heapInt []int

func (h heapInt) Len() int           { return len(h) }
func (h heapInt) Less(i, j int) bool { return h[i] < h[j] }
func (h heapInt) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *heapInt) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *heapInt) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func main() {
	var N, l int

	fmt.Scan(&N)

	h := make(heapInt, 0)
	heap.Init(&h)

	for i := 0; i < N; i++ {
		fmt.Scan(&l)
		heap.Push(&h, l)
	}

	ans := 0
	for h.Len() > 1 {

		l1 := heap.Pop(&h).(int)
		l2 := heap.Pop(&h).(int)

		L := l1 + l2
		ans += L
		heap.Push(&h, L)
	}
	fmt.Println(ans)

}
