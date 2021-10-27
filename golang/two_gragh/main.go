package main

import "fmt"

var G [][]int
var color []int
var V int

func dfs(v, c int) bool {
	fmt.Println(G)
	fmt.Println(color)
	color[v] = c
	for i := 0; i < len(G[i]); i++ {
		if color[G[v][i]] == c {
			return false
		}
		if color[G[v][i]] == 0 && !dfs(G[v][i], -c) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Scan(&V)
	makeColor(V)
	makeGraph(V)

	for i := 0; i < V; i++ {
		if color[i] == 0 {
			if !dfs(i, 1) {
				fmt.Println("No")
				return
			}
		}
	}
	fmt.Println("Yes")
}

func makeGraph(v int) {
	G = make([][]int, v)
	for i := 0; i < v; i++ {
		G[i] = make([]int, v)
	}
}

func makeColor(v int) {
	color = make([]int, V)
	for i := 0; i < v; i++ {
		color[i] = 0
	}
}
