package main

func longestOnes(nums []int, k int) int {
	var (
		count = 0
		left  = 0
		ans   = 0
	)

	for right, num := range nums {
		if num == 0 {
			count++
		}
		for count > k && left <= right {
			if nums[left] == 0 {
				count--
			}
			left++
		}
		ans = _max(ans, right-left+1)
	}

	return ans
}

func _max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
