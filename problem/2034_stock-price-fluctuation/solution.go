package main

import (
	"container/heap"
)

type StockPrice struct {
	current  int
	prices   map[int]int
	minPrice *PriorityQueue
	maxPrice *PriorityQueue
}

func Constructor() StockPrice {
	return StockPrice{
		current:  0,
		prices:   make(map[int]int),
		minPrice: new(PriorityQueue),
		maxPrice: new(PriorityQueue),
	}
}

func (s *StockPrice) Update(timestamp int, price int) {
	heap.Push(s.minPrice, &Item{
		timestamp: timestamp,
		price:     -price,
	})
	heap.Push(s.maxPrice, &Item{
		timestamp: timestamp,
		price:     price,
	})

	s.prices[timestamp] = price
	if timestamp > s.current {
		s.current = timestamp
	}
}

func (s *StockPrice) Current() int {
	return s.prices[s.current]
}

func (s *StockPrice) Maximum() int {
	for {
		if head := (*s.maxPrice)[0]; head.price == s.prices[head.timestamp] {
			return head.price
		}
		heap.Pop(s.maxPrice)
	}
}

func (s *StockPrice) Minimum() int {
	for {
		if head := (*s.minPrice)[0]; -head.price == s.prices[head.timestamp] {
			return -head.price
		}
		heap.Pop(s.minPrice)
	}
}

/**
 * Your StockPrice object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Update(timestamp,price);
 * param_2 := obj.Current();
 * param_3 := obj.Maximum();
 * param_4 := obj.Minimum();
 */

// An Item is something we manage in a priority queue.
type Item struct {
	timestamp int
	price     int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].price > pq[j].price }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x any) { *pq = append(*pq, x.(*Item)) }

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}
