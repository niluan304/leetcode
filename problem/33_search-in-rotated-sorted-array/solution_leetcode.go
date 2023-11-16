package main

// 二分查找, 红蓝染色
func endlesscheng2(nums []int, target int) int {
	isBlue := func(i int) bool {
		end := nums[len(nums)-1]
		if nums[i] > end {
			return target > end && nums[i] >= target
		} else {
			return target > end || nums[i] >= target
		}
	}
	left, right := -1, len(nums) // 开区间 (-1, n-1)
	for left+1 < right {         // 开区间不为空
		mid := left + (right-left)/2
		if isBlue(mid) { // 蓝色
			right = mid
		} else { // 红色
			left = mid
		}
	}
	if right == len(nums) || nums[right] != target {
		return -1
	}
	return right
}
