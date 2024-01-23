package main

// 分组循环
// - 时间复杂度：$\mathcal{O}(n)$。
// - 空间复杂度：$\mathcal{O}(1)$。
func longestAlternatingSubarray(nums []int, threshold int) int {

	var n = len(nums)
	var ans = 0 // math.MaxInt32 // math.MinInt32
	for i := 0; i < n; {
		// 外循环开始的判断
		if nums[i]%2 != 0 || nums[i] > threshold {
			i++
			continue
		}

		// 内循环
		j := i
		for i++; i < n; i++ {
			if nums[i]%2 != (i-j)%2 || nums[i] > threshold {
				break
			}
		}
		ans = max(ans, i-j)
	}
	return ans
}
