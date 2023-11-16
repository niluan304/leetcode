package main

// 暴力解法
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func search(nums []int, target int) int {
	for i, v := range nums {
		if v == target {
			return i
		}
	}
	return -1
}

// 二分查找
// 时间复杂度：O(logn)
// 空间复杂度：O(1)
func search2(nums []int, target int) int {
	var (
		left  = 0
		right = len(nums) - 1
	)

	for left < right-1 {
		mid := left + (right-left)/2
		vMid := nums[mid]
		vRight := nums[right]

		if vMid < vRight {
			if target < vMid {
				right = mid
			} else {
				if target <= vRight {
					left = mid
				} else {
					right = mid
				}
			}
		} else {
			if target > vMid {
				left = mid
			} else {
				if target <= vRight {
					left = mid
				} else {
					right = mid
				}
			}
		}
	}

	for i := left; i <= right; i++ {
		if nums[i] == target {
			return i
		}
	}

	return -1
}

// 二分查找
// 时间复杂度：O(logn)
// 空间复杂度：O(1)
func search3(nums []int, target int) int {
	var (
		left  = 0
		right = len(nums) - 1
	)

	for left < right-1 {
		mid := left + (right-left)/2
		vMid := nums[mid]
		vRight := nums[right]

		if vMid < vRight {
			if target >= vMid && target <= vRight {
				left = mid
			} else {
				right = mid
			}
		} else {
			if target <= vMid && target > vRight {
				right = mid
			} else {
				left = mid
			}
		}
	}

	for i := left; i <= right; i++ {
		if nums[i] == target {
			return i
		}
	}

	return -1
}

// 二分查找
// 时间复杂度：O(logn)
// 空间复杂度：O(1)
func search4(nums []int, target int) int {
	var (
		left  = 0
		right = len(nums) - 1
		last  = nums[right]
	)

	// if ture, 表示 nums[i] == target, 或者 i 在 target 右侧
	// if false, 表示 i 在 target 左侧
	OnRight := func(i int) bool {
		num := nums[i]
		if target > last {
			return num <= last || num >= target
		} else {
			return num <= last && num >= target
		}
	}

	for left < right {
		mid := left + (right-left)/2
		if OnRight(mid) {
			// target 在 [left, mid]范围内
			right = mid
		} else {
			// target 在 (mid, right]范围内
			left = mid + 1
		}
	}

	if nums[right] != target {
		return -1
	}
	return right
}
