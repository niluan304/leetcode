package main

import (
	. "github.com/niluan304/leetcode/container"
)

type Pair struct{ price, timestamp int }

type StockPrice struct {
	current int
	prices  map[int]int

	maxPrice *Heap[Pair] // 大根堆
	minPrice *Heap[Pair] // 小根堆
}

func Constructor() StockPrice {
	return StockPrice{
		current: 0,
		prices:  make(map[int]int),

		maxPrice: NewEmptyHeap(func(x, y Pair) bool { return x.price > y.price }), // 大根堆
		minPrice: NewEmptyHeap(func(x, y Pair) bool { return x.price < y.price }), // 小根堆
	}
}

func (s *StockPrice) Update(timestamp int, price int) {
	pair := Pair{price: price, timestamp: timestamp}
	s.minPrice.Insert(pair)
	s.maxPrice.Insert(pair)

	s.prices[timestamp] = price
	if timestamp > s.current {
		s.current = timestamp
	}
}

func (s *StockPrice) Current() int {
	return s.prices[s.current]
}

func (s *StockPrice) Maximum() int {
	return s.extremum(s.maxPrice)
}

func (s *StockPrice) Minimum() int {
	return s.extremum(s.minPrice)
}

func (s *StockPrice) extremum(hp *Heap[Pair]) int {
	for {
		root := hp.Head()
		if root.price == s.prices[root.timestamp] {
			return root.price
		}
		hp.PopHead()
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
