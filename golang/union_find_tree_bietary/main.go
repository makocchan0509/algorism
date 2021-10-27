package main

import "fmt"

type UnionFind struct {
	par []int
}

func newUnionFind(N int) *UnionFind {
	u := new(UnionFind)
	u.par = make([]int, N)
	for i := range u.par {
		u.par[i] = -1
	}
	return u
}

func (u UnionFind) root(x int) int {
	if u.par[x] < 0 {
		return x
	}
	u.par[x] = u.root(u.par[x])
	return u.par[x]
}

func (u UnionFind) unite(x, y int) {
	xr := u.root(x)
	yr := u.root(y)
	if xr == yr {
		return
	}

	u.par[yr] += u.par[xr]
	u.par[xr] = yr
}

func (u UnionFind) same(x, y int) bool {
	if u.root(x) == u.root(y) {
		return true
	} else {
		return false
	}
}

func main() {

	var N, K int
	fmt.Scan(&N, &K)
	var t, x, y int

	T := make([]int, K)
	X := make([]int, K)
	Y := make([]int, K)

	for i := 0; i < K; i++ {
		fmt.Scan(&t, &x, &y)
		T[i] = t
		X[i] = x
		Y[i] = y
	}

	u := newUnionFind(N * 3)
	ans := 0

	for i := 0; i < K; i++ {
		x := X[i] - 1
		y := Y[i] - 1

		if x < 0 || N <= x || y < 0 || N <= y {
			ans++
			continue
		}

		if T[i] == 1 {
			if u.same(x, y+N) || u.same(x, y+N*2) {
				ans++
			} else {
				u.unite(x, y)
				u.unite(x, y+N)
				u.unite(x, y+N*2)
			}
		} else {
			if u.same(x, y) || u.same(x, y+N*2) {
				ans++
			} else {
				u.unite(x, y+N)
				u.unite(x+N, y+2*N)
				u.unite(x+N*2, y)
			}
		}
	}
	fmt.Println(ans)
}
