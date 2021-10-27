package main

import "fmt"

func main() {
	initHeap()
	push(2)
	push(3)
	push(4)
	push(5)
	fmt.Println(heap)
	fmt.Println(pop())
}

var sz int = 0
var n int = 10
var heap []int = make([]int, n)

func initHeap() {
	for i := 0; i < n; i++ {
		heap[i] = 600000
	}
}

func push(x int) {
	sz++
	i := sz

	for i > 0 {
		p := (i - 1) / 2
		if heap[p] <= x {
			break
		}
		heap[i] = heap[p]
		i = p
	}
	heap[i] = x
}

func pop() int {

	ret := heap[0]
	x := heap[sz]
	sz--

	i := 0
	for i*2+1 < sz {
		a := i*2 + 1
		b := i*2 + 2
		if b < sz && heap[b] < heap[a] {
			a = b
		}

		if heap[a] >= x {
			break
		}

		heap[i] = heap[a]
		i = a
	}
	heap[i] = x
	return ret
}
