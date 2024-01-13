package main

// 滑动窗口模板题
// - 时间复杂度：$\mathcal{O}(n)$。
// - 空间复杂度：$\mathcal{O}(n)$。
func maxSubarrayLength(nums []int, k int) int {
	var count = make(map[int]int, len(nums))

	var ans = 0
	var left = 0
	for right, num := range nums {
		count[num]++
		for count[num] > k {
			count[nums[left]]--
			left++
		}

		ans = max(ans, right-left+1)
	}

	return ans
}
