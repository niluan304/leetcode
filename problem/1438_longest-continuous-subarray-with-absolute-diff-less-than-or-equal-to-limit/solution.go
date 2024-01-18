package main

import "github.com/emirpasic/gods/trees/redblacktree"

// 暴力穷举
// - 时间复杂度：$\mathcal{O}(n^2)$。
// - 空间复杂度：$\mathcal{O}(n)$。
// Deprecated: 超时
func longestSubarray(nums []int, limit int) int {
	var ans = 0
	var n = len(nums)
	for i, num := range nums {
		r := i + 1
		for ; r < n; r++ {
			if nums[r] > num+limit || nums[r] < num {
				break
			}
		}
		l := i - 1
		for ; l >= 0; l-- {
			if nums[l] > num+limit || nums[l] < num {
				break
			}
		}
		ans = max(ans, r-l-1)
	}
	return ans
}

// 双重单调队列
// 时间复杂度: O(n)
// 空间复杂度: O(n)
// 使用单调队列求出 [nums[i], nums[i]+limit] 的左右边界
func longestSubarray2(nums []int, limit int) int {
	var ans = 0
	var n = len(nums)
	var queue []int
	var right, left = make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		right[i] = n
		// 1. 入
		for m := len(queue); m > 0; m-- {
			if nums[queue[m-1]] <= nums[i] {
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
		ans = max(ans, right[i]-left[i]-1)
	}
	return ans
}

// 暴力穷举
// - 时间复杂度：$\mathcal{O}(n^2)$。
// - 空间复杂度：$\mathcal{O}(n)$。
// Deprecated: 超时
func longestSubarray3(nums []int, limit int) int {
	var ans = 0
	var n = len(nums)
	for i, num := range nums {
		mn, mx := num, num
		j := i + 1
		for ; j < n; j++ {
			mn = min(mn, nums[j])
			mx = max(mx, nums[j])

			if mx-mn > limit {
				break
			}
		}
		ans = max(ans, j-i)
	}
	return ans
}

func longestSubarray4(nums []int, limit int) (ans int) {
	var tree = redblacktree.NewWithIntComparator()

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

		ans = max(ans, right-left+1)
	}
	return ans
}
