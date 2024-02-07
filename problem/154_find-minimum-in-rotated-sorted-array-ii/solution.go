package main

func findMin(nums []int) int {
	var (
		left  = 0
		right = len(nums) - 1
		last  = nums[right]
	)

	for left < right {
		mid := left + (right-left)/2

		if nums[mid] > last {
			left = mid + 1
		} else if nums[mid] < last {
			right = mid
		} else {
			break
		}
	}

	v := nums[right]
	for _, num := range nums[left:right] {
		if v > num {
			v = num
		}
	}

	return v
}

func findMin2(nums []int) int {
	var (
		left  = 0
		right = len(nums) - 1
	)

	for left < right {
		mid := left + (right-left)/2
		vMid := nums[mid]
		vRight := nums[right]

		// 设 min := nums[X]
		// vMid > vRight 说明 mid 在 [0, X-1]范围内, 即最小是的左侧
		// vMid < vRight 说明 mid 在 [X, len(nums)-1]范围内, 即最小是的右侧
		if vMid > vRight {
			left = mid + 1
		} else if vMid < vRight {
			right = mid
		} else {
			right--
		}
	}

	return nums[right]
}
