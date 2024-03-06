package main

import (
	. "github.com/niluan304/leetcode/container"
)

type StockPrice struct {
	current  int
	prices   map[int]int
	minPrice *PriorityQueue[int] // price 为优先级，timestamp 为值，这里用 Heap[Pair] 可读性更好
	maxPrice *PriorityQueue[int] // price 为优先级，timestamp 为值，这里用 Heap[Pair] 可读性更好
}

func Constructor() StockPrice {
	return StockPrice{
		current:  0,
		prices:   make(map[int]int),
		minPrice: new(PriorityQueue[int]),
		maxPrice: new(PriorityQueue[int]),
	}
}

func (s *StockPrice) Update(timestamp int, price int) {
	s.minPrice.Insert(timestamp, -price) // 取负数，将大根堆变为小根堆
	s.maxPrice.Insert(timestamp, price)

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
		head := (*s.maxPrice)[0]
		price, timestamp := head.Priority(), head.Value()
		if price == s.prices[timestamp] {
			return price
		}
		s.maxPrice.PopHead()
	}
}

func (s *StockPrice) Minimum() int {
	for {
		head := (*s.minPrice)[0]
		price, timestamp := -head.Priority(), head.Value()
		if price == s.prices[timestamp] {
			return price
		}
		s.minPrice.PopHead()
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
