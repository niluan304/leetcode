package main

// dp 动态规划，LIS
// 时间复杂度：O(n^2)
// 空间复杂度：O(n)
func minimumMountainRemovals(nums []int) int {
	var n, mm = len(nums), 0
	var dp, dp2 = make([]int, n), make([]int, n)

	for i := 0; i < n; i++ {
		length := 0
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				length = _max(length, dp[j])
			}
		}
		dp[i] = length + 1
	}
	for i := 0; i < n; i++ {
		x := n - 1 - i
		length := 0
		for j := x + 1; j < n; j++ {
			if nums[x] > nums[j] {
				length = _max(length, dp2[j])
			}
		}
		dp2[x] = length + 1
	}

	for i := 0; i < n-1; i++ {
		if dp[i] == 1 || dp2[i] == 1 {
			continue
		}
		mm = _max(mm, dp[i]+dp2[i]-1)
	}
	return n - mm
}

// TODO：fail pass case4, case5
// dp 动态规划，LIS
// 时间复杂度：O(n^3)
// 空间复杂度：O(n)
func minimumMountainRemovals2(nums []int) int {
	var n = len(nums)
	var mm = 0
	for i := 1; i < n-1; i++ {
		mm = _max(mm,
			LongestOfLIS(nums[:i], func(num, vj int) bool { return num > vj })+
				LongestOfLIS(nums[i:], func(num, vj int) bool { return num < vj }),
		)
	}

	return n - mm
}

func _max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func _min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func _abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func LongestOfLIS(nums []int, less func(num, vj int) bool) int {
	var n = len(nums)
	var dp = make([]int, n)
	var ans = 0

	for i, num := range nums {
		length := 0

		for j := 0; j < i; j++ {
			if !less(num, nums[j]) {
				continue
			}
			length = _max(length, dp[j])
		}

		dp[i] = length + 1
		ans = _max(ans, dp[i])
	}

	return ans
}
