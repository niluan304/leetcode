package main

// 时间复杂度：O(n)，其中 nnn 为 nums 的长度。
// 最坏情况下，数组元素均相同，循环内会一直执行 right 减一，此时循环会执行 O(n) 次。
// 空间复杂度：O(1)，仅用到若干额外变量。
func endlesscheng(nums []int) int {
	left, right := -1, len(nums)-1 // 开区间 (-1, n-1)
	for left+1 < right {           // 开区间不为空
		mid := left + (right-left)/2
		if nums[mid] < nums[right] { // 蓝色
			right = mid
		} else if nums[mid] > nums[right] { // 红色
			left = mid
		} else {
			right--
		}
	}
	return nums[right]
}

// 0 ms 的代码示例
func leetcodeMinTime(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	i := 0
	for i < len(nums) && nums[i] == nums[len(nums)-1] {
		i++
	}
	if i > len(nums)-1 { // 全部相等  2,2,2,2
		return nums[0]
	}

	j := len(nums)
	for i < j {
		mid := (i + j) / 2
		if nums[mid] <= nums[len(nums)-1] {
			j = mid
		} else {
			i = mid + 1
		}
	}
	return nums[i]
}
