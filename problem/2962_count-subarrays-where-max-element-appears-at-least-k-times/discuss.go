package main

import (
	"cmp"
	"math"
)

func violent(nums []int, k int) int64 {
	var n, ans = len(nums), 0
	for i := 0; i < n; i++ {
		var count, mx = map[int]int{}, math.MinInt
		for j := i; j < n; j++ {
			count[nums[j]]++
			mx = max(mx, nums[j])
			if count[mx] >= k {
				ans++
			}
		}
	}
	return int64(ans)
}

// 修改题目
//
// 给你一个整数数组 nums 和一个 正整数 k 。
//
// 如果子数组中的最大元素出现了至少 k 次，那么称这个子数组为 美丽子数组。
// 请你统计数组 nums 中 有多少 美丽子数组。
//
// 子数组是数组中的一个连续元素序列。
func discuss(nums []int, k int) int64 {
	var n = len(nums)
	var ans = 0
	var left, right = make([]int, n), make([]int, n)
	var stack []int

	// right[i] 表示 i 的右边 [i : right[i]] 的数都小于 nums[i]，不包括等于的情况
	// 从前向后遍历，利用单调栈更新 right[:]
	for i := 0; i < n; i++ {
		right[i] = n // 初始化，假设 nums[i] 在后缀中是唯一最大的

		m := len(stack)
		for m > 0 {
			j := stack[m-1]
			if nums[i] < nums[j] {
				break
			}
			m--
			right[j] = i // nums[i] >= nums[j]，更新 right[j]
		}
		stack = append(stack[:m], i)
	}

	stack = []int{}

	// left[i] 表示 i 的左边 [left[i]: i] 的数都小于等于 nums[i]
	// 从后向前遍历，利用单调栈更新 left[:]
	for i := n - 1; i >= 0; i-- {
		left[i] = -1 // 初始化，假设 nums[i] 在前缀中是最大的

		m := len(stack)
		for m > 0 {
			j := stack[m-1]
			if nums[i] <= nums[j] {
				break
			}
			m--
			left[j] = i // nums[i] > nums[j]，更新 left[i]
		}
		stack = append(stack[:m], i)
	}

	var index = map[int][]int{} // 记录相同 num 的索引位置
	// 从前向后遍历，假设 nums[i] 是 最短子数组的 右边界
	for i, num := range nums {
		index[num] = append(index[num], i)
		m := len(index[num])
		if m < k {
			// num 出现次数不足 k 次，跳过
			continue
		}
		j := index[num][m-k] // j 是 最短子数组的 左边界
		if j < left[i] {
			// 但是 [j : i] 之间有一个数比 nums[i] 更大，在 区间[j: i] 内 num 不是最大值，跳过
			continue
		}

		// nums[i] 为最大值时，满足条件子数组的个数：j 左边小于等于 num 的数量 × i 右边小于 num 的数量
		ans += (j - left[j]) * (right[i] - i)
	}
	return int64(ans)
}

func discuss2(nums []int, k int) int64 {
	var ans, sum, n = 0, 0, len(nums)
	type Item struct{ Index, Count int }

	stack := make([]Item, 0)
	index := make(map[int][]int)

	for i := 0; i < n; i++ {
		m := len(stack)
		for m > 0 && nums[stack[m-1].Index] <= nums[i] {
			sum -= stack[m-1].Count
			m--
		}
		stack = stack[:m]

		top := -1 // 单调递减栈, 栈顶元素 nums[top] 是前一个比 nums[i] 大的元素
		if m > 0 {
			top = stack[m-1].Index
		}

		index[nums[i]] = append(index[nums[i]], i)
		idx := index[nums[i]]

		cnt := 0
		if len(idx) >= k {
			cnt = max(0, idx[len(idx)-k]-top)
		}

		sum += cnt
		ans += sum
		stack = append(stack, Item{Index: i, Count: cnt})
	}

	return int64(ans)
}

// 单调栈 + 动态规划
// - 时间复杂度：$\mathcal{O}(n)$。
// - 空间复杂度：$\mathcal{O}(n)$。
func discuss3(nums []int, k int) int64 {
	var n = len(nums)

	stack := make([]int, 0)
	index := make(map[int][]int)
	dp := make([]int, n+1) //  dp[i] 表示 nums[i] 为右边界时，子数组数量

	for i := 0; i < n; i++ {
		m := len(stack)
		for m > 0 && nums[stack[m-1]] <= nums[i] {
			m--
		}
		stack = stack[:m] // 小于 nums[i] 的元素出栈

		top := -1 // 单调递减栈, 栈顶元素 nums[top] 是前一个比 nums[i] 大的元素
		if m > 0 {
			top = stack[m-1]
		}

		index[nums[i]] = append(index[nums[i]], i)
		idx := index[nums[i]]

		cnt := 0 // cnt 表示 nums[i] 为子数组最大值 且 为右边界时 的子数组数量
		if len(idx) >= k {
			cnt = max(0, idx[len(idx)-k]-top) // cnt 等于左边界到 top 的距离
		}

		// 动态规划递推公式： dp[i] = dp[top] + cnt
		// dp[top] 表示 nums[top]为右边界的子数组数量
		// 因为 nums[top+1:i] <= nums[top]，所以将 nums[top+1:i] 拼到 nums[top] 右边时，会增加 dp[top] 的子数组数量
		// 如果 top == -1, 说明在区间 [0:i] 中， nums[i] 是最大值
		dp[i+1] = dp[top+1] + cnt
		stack = append(stack, i) // 入栈，nums[i] 是新的最小值
	}

	var ans = Sum(dp)
	return int64(ans)
}

func Sum[S ~[]E, E cmp.Ordered](x S) E {
	var m E
	for i := 0; i < len(x); i++ {
		m += x[i]
	}
	return m
}
