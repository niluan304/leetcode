package main

// 方法二：滑动窗口 + 单调队列
// 思路和解法
//
// 在方法一中，我们仅需要统计当前窗口内的最大值与最小值，因此我们也可以分别使用两个单调队列解决本题。
//
// 在实际代码中，我们使用一个单调递增的队列 queMin 维护最小值，一个单调递减的队列 queMax 维护最大值。
// 这样我们只需要计算两个队列的队首的差值，即可知道当前窗口是否满足条件。
func leetcode2(nums []int, limit int) (ans int) {
	var minQ, maxQ []int
	left := 0
	for right, v := range nums {
		for len(minQ) > 0 && minQ[len(minQ)-1] > v {
			minQ = minQ[:len(minQ)-1]
		}
		minQ = append(minQ, v)
		for len(maxQ) > 0 && maxQ[len(maxQ)-1] < v {
			maxQ = maxQ[:len(maxQ)-1]
		}
		maxQ = append(maxQ, v)
		for len(minQ) > 0 && len(maxQ) > 0 && maxQ[0]-minQ[0] > limit {
			if nums[left] == minQ[0] {
				minQ = minQ[1:]
			}
			if nums[left] == maxQ[0] {
				maxQ = maxQ[1:]
			}
			left++
		}
		ans = max(ans, right-left+1)
	}
	return
}
