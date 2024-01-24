package main

import "math"

// 暴力穷举
// - 时间复杂度：$\mathcal{O}(n^2)$。
// - 空间复杂度：$\mathcal{O}(1)$。
// Deprecated: 超时
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
// 这个做法主要是求 以 maxHeights[i] 为山顶时，会浪费多少 height，然后使用 Sum(maxHeight) - min(prefix + suffix)
// 但计算有些重复了，因为使用单调栈可以求 以 maxHeight[i] 为山顶时的最大和
// - 时间复杂度：$\mathcal{O}(n)$。
// - 空间复杂度：$\mathcal{O}(n)$。
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
// 以 maxHeight[i] 为山顶时，高度和的最大值 = 左边的最大高度和 + 右边的最大高度和
// 左右两边的最大高度和，都可以使用单调栈预处理完成。
//
// - 时间复杂度：$\mathcal{O}(n)$。
// - 空间复杂度：$\mathcal{O}(n)$。
func maximumSumOfHeights4(nums []int) int64 {

	var n = len(nums)
	var ans = 0 // math.MaxInt32 // math.MinInt32
	var left, right = make([]int, n+1), make([]int, n+1)

	var stack = []int{-1} // 方便计算的哨兵，假定 nums[-1] = 0
	for i := 0; i < n; i++ {
		m := len(stack)
		for m > 1 && nums[i] <= nums[stack[m-1]] { // 单调递增栈
			m--
		}
		top := stack[m-1]            // 左边第一个小于 nums[i] 的元素
		stack = append(stack[:m], i) // 大于等于 nums[i] 的元素出栈， nums[i] 作为新的最大值入栈

		left[i+1] = (i-top)*nums[i] + left[top+1] // 前缀和计算
	}

	stack = []int{n} // 方便计算的哨兵，假定 nums[n] = 0
	for i := n - 1; i >= 0; i-- {
		m := len(stack)
		for m > 1 && nums[i] <= nums[stack[m-1]] { // 单调递增栈
			m--
		}
		top := stack[m-1]            // 右边第一个小于 nums[i] 的元素
		stack = append(stack[:m], i) // 大于等于 nums[i] 的元素出栈， nums[i] 作为新的最大值入栈

		right[i] = (top-i)*nums[i] + right[top] // 后缀和计算
	}

	for i, num := range nums {
		ans = max(ans, left[i+1]+right[i]-num)
	}

	return int64(ans)
}
