package dp

import "slices"

// LIS
// 常见变形
// - 最长 递增 / 递减 子序列
// - 最长 不降序 / 不升序 子序列
// - 自定义拼接条件

// LIS
// 返回最长递增子序列的长度
//
// LC300 https://leetcode.cn/problems/longest-increasing-subsequence/
// LC1626 https://leetcode.cn/problems/best-team-with-no-conflicts/
// - 时间复杂度：O(n^2)。
// - 空间复杂度：O(n)。
func LIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	ans := 0
	for i, num := range nums {
		dp[i] = 1
		for j, v := range nums[:i] {
			if v < num { // 改为 <= 为非减子序列，改为 > 为递减子序列
				dp[i] = max(dp[i], dp[j]+1)
			}
		}

		ans = max(ans, dp[i])
	}
	return ans
}

// LISGreedy
//
// 贪心下二分
//
// LC1964 https://leetcode.cn/problems/find-the-longest-valid-obstacle-course-at-each-position/
// 需转换：
// - LC354 https://leetcode.cn/problems/russian-doll-envelopes
//
// - 时间复杂度：O(nlogn)
// - 空间复杂度：O(n)。
func LISGreedy(nums []int) int {
	var greedy []int // 定义 greedy[i] 为长度为 i+1 的上升子序列的末尾元素的最小值

	for _, num := range nums {
		i, _ := slices.BinarySearch(greedy, num) // 改成 num+1 为非严格递增
		if i < len(greedy) {
			greedy[i] = num
		} else {
			greedy = append(greedy, num)
		}
	}

	return len(greedy)
}

// LISPath
// 返回最长递增子序列本身，需要额外的 Last 字段来记录是从哪里转移过来的
//
// LC2901 https://leetcode.cn/problems/longest-unequal-adjacent-groups-subsequence-ii/
//
// - 时间复杂度：O(n^2)。
// - 空间复杂度：O(n)。
func LISPath(nums []int) []int {
	type Item struct{ Length, Last int }

	n := len(nums)
	dp := make([]Item, n+1)
	mxIdx := 0
	for i := 1; i <= n; i++ {
		dp[i] = Item{Length: 1, Last: 0}

		for j := 1; j < i; j++ {
			// 遍历 [1, n] 而不是 [0, n-1]，因此索引需要向左偏移 1 位
			if nums[i-1] > nums[j-1] {
				dp[i] = Item{Length: dp[j].Length + 1, Last: j}
			}
		}

		if dp[i].Length > dp[mxIdx].Length {
			mxIdx = i
		}
	}

	var ans = make([]int, dp[mxIdx].Length)
	for i := mxIdx; i > 0; i = dp[i].Last {
		ans[dp[i].Length-1] = nums[i-1]
	}
	return ans
}
