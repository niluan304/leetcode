package main

// 相向双指针
func twoSum(numbers []int, target int) []int {
	var (
		left  = 0
		right = len(numbers) - 1
	)

	// 两个指针相向移动
	// 时间复杂度: O(n)
	// 空间复杂度: O(1)
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			// 答案唯一, 可以直接return
			// 或者 break
			return []int{left + 1, right + 1}
		}

		// numbers 是有序数组,
		// 如果和大于目标值, 说明右指针的值太大, 需要左移
		// 如果和小于目标值, 说明左指针的值太小, 需要右移
		if sum > target {
			right--
		} else {
			left++
		}
	}

	return []int{left + 1, right + 1}
}

// 暴力解法
func twoSum2(numbers []int, target int) []int {
	var n = len(numbers)
	// 时间复杂度: O(n^2)
	// 空间复杂度: O(1)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if numbers[i]+numbers[j] == target {
				return []int{i + 1, j + 1}
			}
		}
	}

	return nil
}
