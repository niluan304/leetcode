package main

// 单调队列
// 时间复杂度: O(n)
// 空间复杂度: O(min(k,U)) U 为数组内数字的种类
func maxSlidingWindow(nums []int, k int) []int {
	var queue []int
	var ans = make([]int, 0, len(nums))

	for i, num := range nums {
		// 1. 入
		for m := len(queue); m > 0 && nums[queue[m-1]] <= num; m-- {
			queue = queue[:m-1] // 维护 q 的单调性
		}
		queue = append(queue, i)

		// 2. 出
		if queue[0] <= i-k {
			// 队首不在滑动窗口内了，需要移除队首元素
			queue = queue[1:]
		}

		// 3. 记录答案
		if i >= k-1 {
			// 由于队首到队尾单调递减，所以窗口最大值就是队首
			ans = append(ans, nums[queue[0]])
		}
	}
	return ans
}
