package main

// 差分数组
// - 时间复杂度：$\mathcal{O}(n)$。
// - 空间复杂度：$\mathcal{O}(n)$。
func checkArray(nums []int, k int) bool {
	nums = append(nums, 0) // 补一个尾 0，方便计算，[left, right+1] 变动
	var n = len(nums)
	var diff = make([]int, n+1)
	diff[1] = nums[0]
	for i := 1; i < n; i++ {
		diff[i+1] = nums[i] - nums[i-1]
	}

	end := n + 1 - k
	for left := 1; left < end; left++ { // 可进行区间减一的范围：left ∈ [1, end-1]
		diff[left] += diff[left-1]
		if diff[left] < 0 {
			return false
		}

		// right = left +k -1
		// 差分数组，区间 [left, right] 减去 diff[left]
		diff[left+k] += diff[left] // diff[right+1] -= x
		diff[left] = 0             // diff[left] +=x
	}

	for i := end; i < n; i++ {
		if diff[i] != 0 {
			return false
		}
	}
	return true
}
