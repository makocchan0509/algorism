package main

import (
	"container/heap"
	"fmt"
)

type Pair struct {
	B, C int
}

type Route struct {
	to, priority, index int
}

type PQ []*Route

func main() {
	var N, M int
	var A, B, C int
	fmt.Scan(&N, &M)
	Roads := make([][]Pair, N)

	for i := 0; i < M; i++ {
		fmt.Scan(&A, &B, &C)
		A--
		B--
		Roads[A] = append(Roads[A], Pair{B, C})
	}
	//fmt.Println(Roads)

	for i := 0; i < N; i++ {
		result := initIntSlice(N, 999999)
		pq := make(PQ, 0)
		heap.Init(&pq)

		push := func(v, d int) {
			if result[v] <= d {
				return
			}
			result[v] = d
			heap.Push(&pq, &Route{v, d, 0})
		}

		for _, v := range Roads[i] {
			push(v.B, v.C)
		}
		//fmt.Println(result)
		for pq.Len() != 0 {
			r := heap.Pop(&pq).(*Route)
			//fmt.Println(r, r.to)
			if result[r.to] != r.priority {
				continue
			}
			for _, v := range Roads[r.to] {
				push(v.B, r.priority+v.C)
			}
		}
		ans := result[i]
		if ans == 999999 {
			ans = -1
		}
		fmt.Println(ans)
	}
}
func initIntSlice(count int, def int) []int {

	slice := make([]int, count)
	for i := 0; i < count; i++ {
		slice[i] = def
	}
	return slice
}

func (pq PQ) Len() int           { return len(pq) }
func (pq PQ) Less(i, j int) bool { return pq[i].priority < pq[j].priority }
func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *PQ) Push(x interface{}) {
	n := len(*pq)
	route := x.(*Route)
	route.index = n
	*pq = append(*pq, route)
}
func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

// var count int
// var max int

// func main() {

// 	var cy, rt int
// 	fmt.Scan(&cy, &rt)
// 	max = rt

// 	var start, end, time int

// 	route := make(map[int]*Route)
// 	indexes := make([]int, rt)

// 	for i := 0; i < rt; i++ {
// 		fmt.Scan(&start, &end, &time)
// 		route[start] = &Route{
// 			start: start,
// 			goal:  end,
// 			time:  time,
// 		}
// 		indexes[i] = start
// 	}

// 	for _, v := range indexes {
// 		//fmt.Println(v)
// 		count = 0
// 		fmt.Println(recursive_route(route, route[v].goal, route[v].start, route[v].time))
// 	}

// }

// func recursive_route(route map[int]*Route, rd int, goal int, sum int) int {
// 	//fmt.Println(count, rd, goal, sum)
// 	if count == max {
// 		return -1
// 	}
// 	count = count + 1
// 	if route[rd] == nil {
// 		return -1
// 	}
// 	if goal == rd {
// 		return sum
// 	} else {
// 		//return sum + route[rd].time
// 		return recursive_route(route, route[rd].goal, goal, route[rd].time+sum)
// 	}
// }
