package main

func removeElement(nums []int, val int) int {
	i := 0
	for i < len(nums) {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i++
		}
	}
	return i
}

func removeElement2(nums []int, val int) int {
	i, j := 0, len(nums)
	for i < j {
		if nums[i] == val {
			nums[i] = nums[j-1]
			j--
		} else {
			i++
		}
	}

	nums = nums[:i]
	return i
}

func leetcode1(nums []int, val int) int {
	left := 0
	for _, v := range nums { // v 即 nums[right]
		if v != val {
			nums[left] = v
			left++
		}
	}
	return left
}
