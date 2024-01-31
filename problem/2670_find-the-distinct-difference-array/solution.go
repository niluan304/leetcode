package main

func distinctDifferenceArray(nums []int) []int {

	var n = len(nums)
	var ans = make([]int, n)
	prefix, suffix := make([]int, n), make([]int, n+1)
	count := map[int]struct{}{}
	for i := 0; i < n; i++ {
		count[nums[i]] = struct{}{}
		prefix[i] = len(count)
	}

	count = make(map[int]struct{})
	for i := n - 1; i >= 0; i-- {
		count[nums[i]] = struct{}{}
		suffix[i] = len(count)
	}

	for i, _ := range ans {
		ans[i] = prefix[i] - suffix[i+1]
	}

	return ans
}
