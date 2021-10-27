package main

import "fmt"

type arr []int

func main() {

	var N,M int
	fmt.Scan(&N,&M)

	G := make([][]int,N)

	for i:=0;i<M;i++{
		var a,b int
		fmt.Scan(&a,&b)
		a -= 1
		b -= 1
		G[a] = append(G[a],b)
		G[b] = append(G[b],a)
	}

	fmt.Println(G)

	dist := make([]int,N)
	for i:=0;i<N;i++{
		dist[i] = -1
	}

	dist[0] = 0
	var q []int

	q = append(q,0)

	for len(q) > 0 {

		i := q[0]
		q = q[1:]

		for _,v := range G[i]{
			if dist[v] == -1{
				dist[v] = dist[i] + 1
				q = append(q,v)
			}
		}
		//fmt.Println(dist)
	}

	//fmt.Println(dist)

	if dist[N-1] == 2{
		fmt.Println("POSSIBLE")
	}else{
		fmt.Println("IMPOSSIBLE")
	}

}


//func (x arr) pop() (int,error){
//
//	if len(x) == 0{
//		return 0,errors.New("len=0")
//	}
//	v := x[0]
//	x = x[1:]
//
//	return v,nil
//}
//
//func (x arr) push(v int) []int{
//	x = append(x,v)
//	return x
//}
