package main

// 暴力穷举
// 时间复杂度: O(n^2)
// 空间复杂度: O(1)
func finalPrices(prices []int) []int {
	var n = len(prices)
	var ans = make([]int, n)

	for i, price := range prices {
		mn := 0
		for j := i + 1; j < n; j++ {
			if prices[j] <= price {
				mn = prices[j]
				break
			}
		}
		ans[i] = prices[i] - mn
	}
	return ans
}

// 单调栈
// 时间复杂度: O(n)
// 空间复杂度: O(n)
func finalPrices2(prices []int) []int {
	var n = len(prices)
	var ans = make([]int, n)
	var stack []int

	for i, price := range prices {
		ans[i] = price

		for m := len(stack); m > 0; m-- {
			j := stack[m-1]
			if price > prices[j] {
				break
			}
			stack = stack[:m-1]
			ans[j] -= price
		}
		stack = append(stack, i)
	}
	return ans
}
