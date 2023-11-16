package main

// dp 动态规划
// 时间复杂度：O(n^2)
// 空间复杂度：O(n)
func lengthOfLIS(nums []int) int {
	var dp = make([]int, len(nums)) // 以 nums[i] 为末尾的最长递增子序列长度
	var ans = 0

	for i, num := range nums {
		for j, v := range nums[:i] {
			if num <= v {
				continue
			}
			dp[i] = _max(dp[i], dp[j]) // 寻找最长递增子序列的
		}
		dp[i]++
		ans = _max(ans, dp[i])
	}

	//// debug: 打印 dp 数组
	//fmt.Println(dp)

	return ans
}

func _max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
