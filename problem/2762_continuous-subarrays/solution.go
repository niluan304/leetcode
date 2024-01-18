package main

import "github.com/emirpasic/gods/trees/redblacktree"

// 暴力穷举
// - 时间复杂度：$\mathcal{O}(n^2)$。
// - 空间复杂度：$\mathcal{O}(n)$。
// Deprecated: 超时
func continuousSubarrays(nums []int) int64 {
	const limit = 2

	var ans = 0
	var n = len(nums)
	for i, num := range nums {
		r := i + 1
		for ; r < n; r++ {
			if nums[r] > num+limit || nums[r] <= num {
				break
			}
		}
		l := i - 1
		for ; l >= 0; l-- {
			if nums[l] > num+limit || nums[l] < num {
				break
			}
		}
		x, y := i-l, r-i
		ans += x * y
	}
	return int64(ans)
}

// 单调队列
// 假设 nums[i] 为最小值
// - 时间复杂度：$\mathcal{O}(n)$。
// - 空间复杂度：$\mathcal{O}(n)$。
func continuousSubarrays2(nums []int) int64 {
	const limit = 2

	var ans = 0
	var n = len(nums)
	var queue []int
	var right, left = make([]int, n), make([]int, n)

	// right[i] 表示 i 的右边 [i : right[i]] 的数范围：[nums[i]+1, nums[i]+limit]
	// 从前向后遍历，利用单调队列更新 right[:]
	for i := 0; i < n; i++ {
		right[i] = n
		// 1. 入
		for m := len(queue); m > 0; m-- {
			if nums[queue[m-1]] < nums[i] { // 改动：把 `<=` 改为 `<`，就解答了。
				break
			}
			right[queue[m-1]] = i
			queue = queue[:m-1] // 维护 q 的单调性
		}
		queue = append(queue, i)

		// 2. 出
		for len(queue) > 0 && nums[queue[0]]+limit < nums[i] {
			right[queue[0]] = i
			// 队首不满足 limit 要求，需要移除队首元素
			queue = queue[1:]
		}
	}

	queue = []int{}
	// left[i] 表示 i 的左边 [left[i]: i] 的数范围：[nums[i], nums[i]+limit]
	// 从后向前遍历，利用单调队列更新 left[:]
	for i := n - 1; i >= 0; i-- {
		left[i] = -1

		// 1. 入
		for m := len(queue); m > 0; m-- {
			if nums[queue[m-1]] <= nums[i] {
				break
			}
			left[queue[m-1]] = i
			queue = queue[:m-1] // 维护 q 的单调性
		}
		queue = append(queue, i)

		// 2. 出
		for len(queue) > 0 && nums[queue[0]]+limit < nums[i] {
			left[queue[0]] = i
			// 队首不满足 limit 要求，需要移除队首元素
			queue = queue[1:]
		}
	}

	for i := 0; i < n; i++ {
		x, y := i-left[i], right[i]-i
		ans += y * x
	}
	return int64(ans)
}

// 暴力穷举
//
// - 时间复杂度：$\mathcal{O}(n^2)$。
// - 空间复杂度：$\mathcal{O}(n)$。
// Deprecated: 超时
func continuousSubarrays3(nums []int) int64 {
	const limit = 2

	var ans = 0
	var n = len(nums)
	for i, num := range nums {
		mn, mx := num, num
		for j := i; j < n; j++ {
			mn = min(mn, nums[j])
			mx = max(mx, nums[j])

			if mx-mn > limit {
				break
			}

			ans++
		}
	}
	return int64(ans)
}

// 滑动窗口 + 有序集合
// 使用合适的的数据结构，可以有效降低题目难度
//
// - 时间复杂度：$\mathcal{O}(n\log n)$。
// - 空间复杂度：$\mathcal{O}(n)$。
func continuousSubarrays4(nums []int) int64 {
	const limit = 2
	var tree = redblacktree.NewWithIntComparator()

	var ans = 0
	var left = 0
	for right, num := range nums {
		node := tree.GetNode(num)
		if node == nil {
			node = &redblacktree.Node{Value: 0}
		}
		tree.Put(num, node.Value.(int)+1)

		for {
			mx, mn := tree.Right().Key.(int), tree.Left().Key.(int)
			if mx-mn <= limit {
				break
			}

			num := nums[left]
			node := tree.GetNode(num)
			v := node.Value.(int) - 1

			if v == 0 {
				tree.Remove(num)
			} else {
				tree.Put(num, v)
			}

			left++
		}

		ans += right - left + 1
	}
	return int64(ans)
}
