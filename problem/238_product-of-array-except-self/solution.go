package main

func productExceptSelf(nums []int) []int {
	var n = len(nums)
	var prefix, suffix = make([]int, n), make([]int, n)
	prefix[0], suffix[n-1] = 1, 1
	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] * nums[i-1]
	}
	for i := n - 2; i >= 0; i-- {
		suffix[i] = suffix[i+1] * nums[i+1]
	}

	var ans = make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = prefix[i] * suffix[i]
	}
	return ans
}
