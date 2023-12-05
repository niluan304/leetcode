package main

// 滑动窗口 + 单调队列
// 时间复杂度: O(n)
// 空间复杂度: O(n)
func maximumRobots(chargeTimes []int, runningCosts []int, budget int64) int {
	var left, sum, ans = 0, 0, 0
	var queue []int // 单调队列，记录当前窗口的最大值
	for right, cost := range runningCosts {
		m := len(queue)
		for m > 0 && chargeTimes[queue[m-1]] <= chargeTimes[right] {
			m--
		}
		queue = append(queue[:m], right)
		sum += cost

		// 不满足条件时，向前移动滑动窗口的左边界
		for len(queue) > 0 && (sum*(right-left+1)+chargeTimes[queue[0]]) > int(budget) {
			sum -= runningCosts[left]
			left++
			for len(queue) > 0 && queue[0] < left {
				queue = queue[1:]
			}
		}

		ans = max(ans, right-left+1)
	}
	return ans
}
