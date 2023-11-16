package main

// 双指针
// 使用双指针，左指针指向当前已经处理好的序列的尾部，右指针指向待处理序列的头部。
// 右指针不断向右移动，每次右指针指向非零数，则将左右指针对应的数交换，同时左指针右移。
func leetcode(nums []int) {
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}

// 8ms 代码示例
func leetcode2(nums []int) {
	length := len(nums)
	var count int
	for _, num := range nums {
		if num != 0 {
			nums[count] = num
			count++
		}
	}

	for count < length {
		nums[count] = 0
		count++
	}
}
