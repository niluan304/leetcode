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

type Value3 struct {
	top, n int
}

type StockSpanner3 struct {
	stack []Value3
}

// Constructor3
// 线段栈？
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func Constructor3() *StockSpanner3 {
	return &StockSpanner3{
		stack: []Value3{
			{top: math.MinInt, n: 0},
		},
	}
}

func (s *StockSpanner3) Next(price int) int {
	n := len(s.stack)
	if price < s.stack[n-1].top {
		s.stack = append(s.stack, Value3{
			top: price,
			n:   1,
		})
		return 1
	}

	s.stack[n-1].top = price
	s.stack[n-1].n++

	var ans = s.stack[n-1].n
	for i := n - 2; i >= 0; i-- {
		stack := s.stack[i]
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

type Value struct {
	price int // 股票价格
	day   int // 小于或等于今天价格的最大连续日数（从今天开始往回数，包括今天）
}

type StockSpanner4 struct {
	stack []Value
}

// Constructor4
// 单调栈
// 时间复杂度：O(1)
// 空间复杂度：O(n)
func Constructor4() StockSpanner4 {
	return StockSpanner4{
		stack: make([]Value, 0),
	}
}

func (s *StockSpanner4) Next(price int) int {
	ans := 1
	for n := len(s.stack); n > 0; n-- {
		top := s.stack[n-1]
		if price < top.price {
			break
		}
		s.stack = s.stack[:n-1]
		ans += top.day
	}
	s.stack = append(s.stack, Value{
		price: price,
		day:   ans,
	})
	return ans
}
