package main

// 差分数组
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func checkArray(nums []int, k int) bool {
	var n = len(nums)
	var diff = make([]int, n+1)
	diff[0] = nums[0]
	diff[n] = -nums[n-1] // [left, right+1] 变动
	for i := 1; i < n; i++ {
		diff[i] = nums[i] - nums[i-1]
	}

	for i := 0; i+k <= n; i++ {
		if diff[i] < 0 {
			return false // 只能减 1，不能加 1
		}

		diff[i+k] += diff[i] // -= -diff[left]
		diff[i] = 0          // += -diff[left]
	}
	for i := 0; i < k; i++ {
		if diff[n-i] != 0 {
			return false
		}
	}
	return true
}

// 差分数组
// 逆向操作： 如果 全 0 的数组可以通过 +1 的操作，变成目标数组，那么说明原数组可以通过 -1 的操作变为全 0 数组
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func checkArray2(nums []int, k int) bool {
	var n = len(nums)
	var diff = make([]int, n+1)
	diff[0] = nums[0]
	diff[n] = -nums[n-1]
	for i := 1; i < n; i++ {
		diff[i] = nums[i] - nums[i-1]
	}

	var zeros = make([]int, n+1)
	for i := 0; i+k <= n; i++ {
		d := diff[i] - zeros[i]
		if d < 0 {
			return false // 题目要求只能减 1，不能加 1
		}
		zeros[i] += d
		zeros[i+k] -= d
	}

	for i := 0; i < k; i++ {
		if diff[n-i] != zeros[n-i] {
			return false
		}
	}
	return true
}
