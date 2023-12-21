package main

import "math"

// 暴力穷举
// 时间复杂度: O(n^2)
// 空间复杂度: O(1)
func maximumSumOfHeights(maxHeights []int) int64 {
	var ans = 0

	var n = len(maxHeights)
	for i, height := range maxHeights {
		x, r, l := height, height, height
		for j := i + 1; j < n; j++ {
			r = min(r, maxHeights[j])
			x += r
		}
		for j := i - 1; j >= 0; j-- {
			l = min(l, maxHeights[j])
			x += l
		}
		ans = max(ans, x)
	}
	return int64(ans)
}

// 前后缀分解 + 单调栈
// 时间复杂度: O(n)
// 空间复杂度: O(n)
func maximumSumOfHeights2(maxHeights []int) int64 {
	var n = len(maxHeights)
	var stack []int
	var prefix = make([]int, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i]
		m := len(stack)
		for m > 0 {
			d := maxHeights[stack[m-1]] - maxHeights[i]
			if d < 0 {
				break
			}
			if m > 1 {
				prefix[i+1] += (stack[m-1] - stack[m-2]) * d
			} else {
				prefix[i+1] += (stack[m-1] + 1) * d
			}
			m--
		}
		stack = append(stack[:m], i)
	}

	stack = stack[:0]
	var suffix = make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		suffix[i] = suffix[i+1]
		m := len(stack)
		for m > 0 {
			d := maxHeights[stack[m-1]] - maxHeights[i]
			if d < 0 {
				break
			}
			if m == 1 {
				suffix[i] += (n - stack[m-1]) * d
			} else {
				suffix[i] += (stack[m-2] - stack[m-1]) * d
			}
			m--
		}
		stack = append(stack[:m], i)
	}

	var ans, sum = math.MaxInt, 0
	for i, _ := range maxHeights {
		sum += maxHeights[i]
		ans = min(ans, prefix[i+1]+suffix[i])
	}
	return int64(sum - ans)
}

// 前后缀分解 + 单调栈
// 时间复杂度: O(n)
// 空间复杂度: O(n)
func maximumSumOfHeights3(maxHeights []int) int64 {
	var n = len(maxHeights)
	var stack = []int{-1} // 方便计算前缀和的哨兵
	var prefix = make([]int, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i]
		m := len(stack)
		// stack[0] = -1, 为哨兵，单调栈长度实际从 2 开始
		for m > 1 && maxHeights[stack[m-1]] >= maxHeights[i] {
			prefix[i+1] += (stack[m-1] - stack[m-2]) * (maxHeights[stack[m-1]] - maxHeights[i])
			m--
		}
		stack = append(stack[:m], i)
	}

	stack = []int{n} // 方便计算后缀和的哨兵
	var suffix = make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		suffix[i] = suffix[i+1]
		m := len(stack)
		// stack[0] = n, 为哨兵，单调栈长度实际从 2 开始
		for m > 1 && maxHeights[stack[m-1]] >= maxHeights[i] {
			suffix[i] += (stack[m-2] - stack[m-1]) * (maxHeights[stack[m-1]] - maxHeights[i])
			m--
		}
		stack = append(stack[:m], i)
	}

	var ans, sum = math.MaxInt, 0
	for i, _ := range maxHeights {
		sum += maxHeights[i]
		ans = min(ans, prefix[i+1]+suffix[i])
	}
	return int64(sum - ans)
}

// 前后缀分解 + 单调栈
// 时间复杂度: O(n)
// 空间复杂度: O(n)
func maximumSumOfHeights4(maxHeights []int) int64 {
	var n = len(maxHeights)
	var stack = []int{-1} // 方便计算前缀和的哨兵
	var prefix = make([]int, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] + maxHeights[i]
		m := len(stack)
		// stack[0] = -1, 为哨兵，单调栈长度实际从 2 开始
		for m > 1 && maxHeights[stack[m-1]] >= maxHeights[i] {
			prefix[i+1] -= (stack[m-1] - stack[m-2]) * (maxHeights[stack[m-1]] - maxHeights[i])
			m--
		}
		stack = append(stack[:m], i)
	}

	stack = []int{n} // 方便计算后缀和的哨兵
	var suffix = make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		suffix[i] = suffix[i+1] + maxHeights[i]
		m := len(stack)
		// stack[0] = n, 为哨兵，单调栈长度实际从 2 开始
		for m > 1 && maxHeights[stack[m-1]] >= maxHeights[i] {
			suffix[i] -= (stack[m-2] - stack[m-1]) * (maxHeights[stack[m-1]] - maxHeights[i])
			m--
		}
		stack = append(stack[:m], i)
	}

	var ans = 0
	for i, height := range maxHeights {
		ans = max(ans, prefix[i+1]+suffix[i]-height) // 前后缀和都计算了 maxHeights[i]，需删去一个
	}
	return int64(ans)
}
