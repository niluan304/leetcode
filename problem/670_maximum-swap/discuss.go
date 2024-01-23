package main

// https://leetcode.cn/problems/maximum-swap/solutions/2614470/yi-ci-bian-li-jian-ji-xie-fa-pythonjavac-c9b1
// 思考题
// 1. 如果要交换的两个数，下标相差不能超过 k 呢？
// 2. 如果可以交换两次呢？

func discuss(num string, k int) string {
	var n = len(num)
	var maxSuffix = make([]int, n)
	var queue []int

	for i := n - 1; i >= 0; i-- {
		for len(queue) > 0 && queue[0]-i > k { // 队首的数据已过期，弹出
			queue = queue[1:]
		}

		m := len(queue)
		for m > 0 && num[i] > num[queue[m-1]] { // 维护队列的单调性
			m--
		}

		queue = append(queue[:m], i)
		maxSuffix[i] = queue[0] // 区间 [i, i+k] 的最大值为 nums[queue[0]]
	}

	s := []byte(num)
	for i, j := range maxSuffix {
		if s[i] != s[j] {
			s[i], s[j] = s[j], s[i]
			break
		}
	}

	return string(s)
}
