package main

// 暴力解法
func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {
	var (
		left = 0
	)

	for right, num := range nums {
		if right-left > indexDiff {
			left++
		}

		for i := left; i < right; i++ {
			if _abs(nums[i]-num) <= valueDiff {
				return true
			}
		}
	}

	return false
}

func _abs(x int) int {
	if x > 0 {
		return x
	}
	return -x

}

func containsNearbyAlmostDuplicate2(nums []int, indexDiff int, valueDiff int) bool {
	// 更高效的解法

	return false
}
