package main

// 前后缀分解
// - 时间复杂度：$\mathcal{O}(n)$。
// - 空间复杂度：$\mathcal{O}(n)$。
func maximumTripletValue(nums []int) int64 {
	var n = len(nums)
	var prefix, suffix = make([]int, n), make([]int, n)
	for i := 1; i < n; i++ {
		prefix[i] = max(prefix[i-1], nums[i-1])
	}
	for i := n - 2; i >= 0; i-- {
		suffix[i] = max(suffix[i+1], nums[i+1])
	}

	var ans = 0
	for i, num := range nums {
		ans = max(ans, (prefix[i]-num)*suffix[i])
	}
	return int64(ans)
}

// maximumTripletValue 的空间优化版本，并同时更新答案和前缀最大值
func maximumTripletValue2(nums []int) int64 {
	var n = len(nums)
	var suffix = make([]int, n+1)

	for i := n - 1; i >= 0; i-- {
		suffix[i] = max(suffix[i+1], nums[i])
	}

	var ans = 0
	var prefix = 0
	for i, num := range nums {
		ans = max(ans, (prefix-num)*suffix[i+1])
		prefix = max(prefix, nums[i])
	}
	return int64(ans)
}
