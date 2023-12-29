package main

import (
	"sort"
)

func buyChoco(prices []int, money int) int {
	sort.Ints(prices)
	sum := prices[0] + prices[1]
	if sum > money {
		return money
	}
	return money - sum
}

func buyChoco2(prices []int, money int) int {
	mn1, mn2 := min(prices[0], prices[1]), max(prices[0], prices[1])
	for i := 2; i < len(prices); i++ {
		if prices[i] < mn1 {
			mn1, mn2 = prices[i], mn1
		} else if prices[i] < mn2 {
			mn2 = prices[i]
		}
	}

	if sum := mn1 + mn2; sum > money {
		return money
	} else {
		return money - sum
	}
}
