package main

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (right-left)>>1 + left // (left + right) / 2 防溢出的写法

		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}

func Search2(nums []int, left, right, target int) int {

	if left > right {
		return -1
	}

	mid := (left + right) / 2 // 会溢出 (right-left)/2 + left, 又(x/2 == x>>1)

	if nums[mid] < target {
		return Search2(nums, mid+1, right, target)
	} else if nums[mid] > target {
		return Search2(nums, left, mid-1, target)
	} else {
		return mid
	}

}
