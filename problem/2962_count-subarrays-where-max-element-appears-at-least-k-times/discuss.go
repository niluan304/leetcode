package main

import (
	"math"
)

// 修改题目
//
// 给你一个整数数组 nums 和一个 正整数 k 。
//
// 如果子数组中的最大元素出现了至少 k 次，那么称这个子数组为 美丽子数组。
// 请你统计数组 nums 中 有多少 美丽子数组。
//
// 子数组是数组中的一个连续元素序列。

// 修改后的题目很像：[907. 子数组的最小值之和](https://leetcode.cn/problems/sum-of-subarray-minimums/description/)
// 但是多一个要求：出现次数至少为 k

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

	index := make(map[int][]int, n) // 记录相同 num 的出现索引
	stack := []int{-1}              // 哨兵
	dp := make([]int, n+1)          // dp[i] 表示 nums[i] 为右边界时，子数组数量
	ans := 0

	for i := 0; i < n; i++ {
		m := len(stack)
		for m > 1 && nums[stack[m-1]] <= nums[i] {
			m--
		}
		stack = stack[:m]        // 小于 nums[i] 的元素出栈
		top := stack[m-1]        // 单调递减栈, 栈顶元素 nums[top] 是前一个比 nums[i] 大的元素
		stack = append(stack, i) // nums[i] 作为新的最小值入栈

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
		ans += dp[i+1]
	}
	return int64(ans)
}

// discuss 的优化版本，只需要一次遍历
//
// 维护一个单调递减栈，每次加入新元素 nums[i] 前 需要将小于当前元素的出栈，
// 假设有栈 stack，其大小为 m，此时 stack[m-1] 出栈，说明 nums[stack[m-1]] 的右边界（第一个大于）的下标就是 i，
//
// 如果 stack 的大小 m 是大于等于 k的，且 nums[stack[m-k]]==nums[stack[m-1]]
// 说明单调栈里倒数 k个都的下标的值都是一样的
// 因为仍在单调栈里，说明中间一定没有比他们大的 否则早就出栈了！
// 那么在区间 [stack[m-k], stack[m-1]] 的最大值，都会是 nums[stack[m-1]]，
//
// 此时就可以计算区间 [stack[m-k], stack[m-1]] 合格子数组的数量
// 其左边界就是第一个大于等于 nums[stack[m-1]]的 回顾我们栈单调不增的特性，
// 容易得到下标就是 stack[m-k-1] 就是在 再向左一个，如果m-k == 0， 那么左边界是-1
// 其右边界是一个大于 nums[stack[m-1]]的下标 其实就是i，
// 因为 nums[i] 要入栈才导致他们要出栈的，那么子数组的数量为：(stack[m-k] - stack[m-k-1]) * (i - stack[m-1])
func discuss4(nums []int, k int) int64 {
	n := len(nums)
	stack := []int{-1} // 左边界哨兵，表示 nums[-1] 位置有一个非常大的数字，用于计算最短子数组的左边界的覆盖范围
	ans := 0

	nums = append(nums, math.MaxInt) // 右边界哨兵，表示 nums[n] 位置有一个非常大的数字，用于弹出单调栈内剩余元素
	for i := 0; i <= n; i++ {
		// 需要将小于当前值的出栈意味着下标 i，就是被出栈元素的右边界
		// 如果只考虑 k=1的情况下 对于下标为 i的元素其左边界为第一个 大于等于 nums[i]的数 不妨令其下标为 l
		// 右边界是第一个 大于 nums[i]的数 令其下标为 r 此时 nums[i] 作为最大数的子数组个数为 (r-i) * (i-l)
		// 相同的思想 可以把连续 k个作为一个整体 然后套上面的思路同时也要求了 栈里尾部至少有 k个值相同的下标才行

		m := len(stack)
		for m > 1 && nums[stack[m-1]] < nums[i] {
			// 栈里至少有 k个数 才能保证出栈的第一个元素有 k个 是一个必要条件
			// 栈里 倒数第k个数 和倒数第一个相等 因为有单调递减栈，相当于快速判断了 末尾k个都是相等的
			if j := m - k; j >= 1 && nums[stack[m-1]] == nums[stack[j]] {
				// 检测倒数第 k个左边时候还有值 也就是这个整体左边界 第一个大于等于的下标
				// 因此我们需要判断倒数第k个是不是最大数 如果是 说明其左边界是-1 不然左边界是栈里再往左一个
				// 再次回顾 维护的栈是单调不减的

				// 栈倒数第 k个的左边还有值 再往左一个的下标就是其边界,
				// 令倒数第 k个 在 stack 的下标为 j 的话 对应回前面的思想，其左边数量是 stack[j] - stack[j-1]
				// 对应回上面的 (i-l)*(r-i) 把区间 [j, m-1] 看成一个整体

				x := stack[j] - stack[j-1] // 左边界的最大覆盖范围
				y := i - stack[m-1]        // 右边界的最大覆盖范围
				ans += x * y               // 贡献法：左边界的最大覆盖范围 * 右边界的最大覆盖范围
			}
			m--
		}
		stack = append(stack[:m], i)
	}
	return int64(ans)
}
