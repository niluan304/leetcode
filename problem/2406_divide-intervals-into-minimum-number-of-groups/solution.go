package main

// 另外一种思路是转换成上下车模型，每个区间看成一个人，他在 left 刻上车，right+1 时刻下车，最后答案为同时在车上的人数的最大值。
// 这可以用差分数组实现，下面代码用的平衡树，方便从小到大计算。
// [1094. 拼车](https://leetcode.cn/problems/car-pooling/)
// 差分数组
// 时间复杂度：(n)
// 空间复杂度：(n)
func minGroups(intervals [][]int) int {
	var n = 0
	for _, interval := range intervals {
		n = _max(n, interval[1])
	}

	var diff = make([]int, n+2)
	for _, interval := range intervals {
		left, right := interval[0], interval[1]
		diff[left]++
		diff[right+1]--
	}

	var ans = 0
	for i := 1; i <= n; i++ {
		diff[i] += diff[i-1]
		ans = _max(ans, diff[i])
	}

	return ans
}

func _max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
