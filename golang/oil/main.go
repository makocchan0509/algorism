package main

import (
	"container/heap"
	"fmt"
)

func main() {
	var N, L, P, a, b int
	fmt.Scan(&N, &L, &P)
	A := make([]int, N)
	B := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scan(&a)
		A[i] = a
	}
	for i := 0; i < N; i++ {
		fmt.Scan(&b)
		B[i] = b
	}

	pq := make(PQ, 0)
	heap.Init(&pq)

	pos := 0
	ans := 0
	tank := P
	nA := append(A, 25)

	for i := 0; i < len(nA); i++ {
		d := nA[i] - pos
		for tank-d < 0 {
			o := heap.Pop(&pq).(int)
			tank += o
			ans++
		}

		pos = nA[i]
		if i != len(nA)-1 {
			heap.Push(&pq, B[i])
		}
		tank -= d
	}
	fmt.Println(ans)

}

func initIntSlice(count int, def int) []int {

	slice := make([]int, count)
	for i := 0; i < count; i++ {
		slice[i] = def
	}
	return slice
}

type PQ []int

func (pq PQ) Len() int           { return len(pq) }
func (pq PQ) Less(i, j int) bool { return pq[i] > pq[j] }
func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PQ) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}
func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}
