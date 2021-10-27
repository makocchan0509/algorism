package main

import (
	"container/heap"
	"fmt"
	"os"
	"strings"
)

var N, M int

var maze [][]string
var route [][]int
var INF int = 2147483648
var bx = [4]int{1, 0, -1, 0}
var by = [4]int{0, 1, 0, -1}

func main() {
	var m string
	fmt.Scan(&N, &M)
	maze = make([][]string, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&m)
		maze[i] = strings.Split(m, "")
	}
	makeInt(N, M, INF)

	fmt.Println(maze)
	fmt.Println(route)
	ans := bfs()
	fmt.Println(ans)
}

func bfs() int {
	q := make(Queue, 0)
	heap.Init(&q)

	y, x := setStart()
	heap.Push(&q, &Point{y, x})
	route[y][x] = 0

	gy, gx := setGoal()

	for q.Len() > 0 {
		point := heap.Pop(&q).(*Point)

		if point.first == gy && point.second == gx {
			break
		}

		for i := 0; i < 4; i++ {
			ny := point.first + by[i]
			nx := point.second + bx[i]

			if 0 <= ny && ny < N && 0 <= nx && nx < M && maze[ny][nx] != "#" && route[ny][nx] == INF {
				heap.Push(&q, &Point{ny, nx})
				route[ny][nx] = route[point.first][point.second] + 1
			}
		}

		desc()
	}
	return route[gy][gx]
}

func desc() {
	fmt.Println("------- maze ------")
	for i := 0; i < N; i++ {
		fmt.Println(maze[i])
	}
	fmt.Println("------- route ------")
	for i := 0; i < N; i++ {
		fmt.Println(route[i])
	}
}

func makeInt(x, y, def int) {
	route = make([][]int, x)
	for dy := 0; dy < y; dy++ {
		for dx := 0; dx < x; dx++ {
			route[dy] = append(route[dy], INF)
		}
	}
}

func setStart() (int, int) {
	for y := 0; y < N; y++ {
		for x := 0; x < M; x++ {
			if maze[y][x] == "S" {
				return y, x
			}
		}
	}
	fmt.Println("Not found start")
	os.Exit(1)
	return 0, 0
}

func setGoal() (int, int) {
	for y := 0; y < N; y++ {
		for x := 0; x < M; x++ {
			if maze[y][x] == "G" {
				return y, x
			}
		}
	}
	fmt.Println("Not found goal")
	os.Exit(1)
	return 0, 0
}

type Queue []*Point

type Point struct {
	first  int
	second int
}

func (q Queue) Len() int           { return len(q) }
func (q Queue) Less(i, j int) bool { return q[i].first < q[j].first }
func (q Queue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}
func (q *Queue) Push(x interface{}) {
	point := x.(*Point)
	*q = append(*q, point)
}
func (q *Queue) Pop() interface{} {
	old := *q
	n := len(old)
	x := old[n-1]
	*q = old[0 : n-1]
	return x
}
