package main

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	return max(robI(nums[1:]), robI(nums[:len(nums)-1]))
}

// 动态规划, 空间优化
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func robI(nums []int) int {
	var n = len(nums)
	var dp0, dp1 = 0, 0

	for i := 0; i < n; i++ {
		// 2. 第i个房间的最大值, 选不选 i 的问题
		v := max(dp0, dp1+nums[i])
		dp1 = dp0
		dp0 = v

		//// 5. debug 打印dp数组
		//fmt.Print(dp0, " ")
	}
	return dp0
}
