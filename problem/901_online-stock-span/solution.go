package main

import "math"

type StockSpanner struct {
	prices []int
}

func Constructor() *StockSpanner {
	return &StockSpanner{
		prices: make([]int, 0),
	}
}

func (s *StockSpanner) Next(price int) int {
	var ans = 1
	for i := len(s.prices) - 1; i >= 0; i-- {
		if s.prices[i] > price {
			break
		}
		ans++
	}
	s.prices = append(s.prices, price)
	return ans
}

func Constructor2() *StockSpanner2 {
	return &StockSpanner2{
		stacks: make([][]int, 0),
		last:   math.MaxInt,
	}
}

type StockSpanner2 struct {
	stacks [][]int
	last   int
}

func (s *StockSpanner2) Next(price int) int {
	defer func() {
		s.last = price
	}()

	if price < s.last {
		s.stacks = append(s.stacks, []int{price})
		return 1
	}

	var n = len(s.stacks)
	var ans = 1
	for i := n - 1; i >= 0; i-- {
		stack := s.stacks[i]
		m := len(stack)
		if price < stack[m-1] {
			break
		}
		ans += m
	}
	s.stacks[n-1] = append(s.stacks[n-1], price)
	return ans
}

func Constructor3() *StockSpanner3 {
	return &StockSpanner3{
		stacks: []Stack3{
			{top: math.MinInt, n: 0},
		},
	}
}

type Stack3 struct {
	top, n int
}

type StockSpanner3 struct {
	stacks []Stack3
}

func (s *StockSpanner3) Next(price int) int {
	n := len(s.stacks)
	if price < s.stacks[n-1].top {
		//n++
		//s.stacks = append(s.stacks, Stack3{})

		s.stacks = append(s.stacks, Stack3{
			top: price,
			n:   1,
		})
		return 1
	}

	s.stacks[n-1].top = price
	s.stacks[n-1].n++

	var ans = s.stacks[n-1].n
	for i := n - 2; i >= 0; i-- {
		stack := s.stacks[i]
		if price < stack.top {
			break
		}
		ans += stack.n
	}

	return ans
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
