package main

func longestSubarray(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	var (
		cost = 0
		left = 0
		ans  = 0
	)

	for right, num := range nums {
		if num != 1 {
			cost++
		}

		// for循环:
		// 使得子串 nums[left:right+1] 满足要求: 0的个数 <= 1
		for cost > 1 {
			if nums[left] == 0 {
				cost--
			}
			left++
		}

		// ans = 子串长度 - 1,
		// 所以也不需要额外处理全1的情况
		ans = max(ans, right-left)
	}

	return ans
}

func longestSubarray2(nums []int) int {
	var (
		cost = 0
		left = 0
	)

	for _, num := range nums {
		if num != 1 {
			cost++
		}

		if cost > 1 {
			if nums[left] == 0 {
				cost--
			}
			left++
		}
	}

	// ans = 滑动窗口长度 - 1
	return len(nums) - left - 1
}
