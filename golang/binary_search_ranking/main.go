package main

import (
	"fmt"
	"sort"
)

func (item *Item) String() string {
	return fmt.Sprintf("{%s:%d}", item.id, item.score)
}

type Ranking struct {
	items []*Item
	m     map[string]*Item
}

func (r *Ranking) String() string {
	return fmt.Sprintf("%v", r.items)
}

type Item struct {
	id    string
	score int
}

func (r *Ranking) rankAdd(item *Item) {
	fmt.Println(r, "<=", item)

	//remove
	if old, ok := r.m[item.id]; ok {
		i := sort.Search(len(r.items), func(i int) bool {
			return r.items[i].score < old.score ||
				(r.items[i].score == old.score && r.items[i].id >= old.id)
		})
		fmt.Println(i)
		copy(r.items[i:], r.items[i+1:])
		r.items[len(r.items)-1] = nil
		r.items = r.items[:len(r.items)-1]
		delete(r.m, old.id)
	}

	i := sort.Search(len(r.items), func(i int) bool {
		return r.items[i].score < item.score ||
			(r.items[i].score == item.score && r.items[i].id > item.id)
	})

	r.items = append(r.items, nil)
	copy(r.items[i+1:], r.items[i:])
	r.items[i] = item
	r.m[item.id] = item
}

func add(items []Item, item Item) []Item {
	fmt.Println(items, "<=", item)

	i := sort.Search(len(items), func(i int) bool {
		return items[i].score < item.score ||
			(items[i].score == item.score && items[i].id > item.id)
	})

	fmt.Println(i)

	items = append(items, Item{})

	fmt.Println(items)
	copy(items[i+1:], items[i:])

	fmt.Println(items)
	items[i] = item

	fmt.Println(items)
	return items
}

func main() {

	// items := []Item{}
	// items = add(items, Item{"A", 100})
	// items = add(items, Item{"C", 300})
	// items = add(items, Item{"B", 200})

	// fmt.Println(items)

	// s := [5]int{1, 2, 3, 4, 5}
	// copy(s[2:], s[1:])

	// fmt.Println(s)
	ranking := &Ranking{make([]*Item, 0, 0), make(map[string]*Item)}
	ranking.rankAdd(&Item{"A", 100})
	ranking.rankAdd(&Item{"C", 200})
	ranking.rankAdd(&Item{"B", 200})
	ranking.rankAdd(&Item{"D", 50})
	ranking.rankAdd(&Item{"E", 100})
	ranking.rankAdd(&Item{"A", 10})
	fmt.Println(ranking)
}
