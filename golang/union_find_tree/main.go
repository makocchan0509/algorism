package main

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

func (u UnionFind) size(x int) int {
	return -u.par[u.root(x)]
}
func main() {

	N := 10
	a := 1
	b := 2
	// インスタンスの生成
	u := newUnionFind(N)

	// 各メソッドの利用
	u.unite(a, b)
	u.same(a, b)
	u.size(a)
	u.root(a)

}
