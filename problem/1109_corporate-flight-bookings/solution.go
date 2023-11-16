package main

// 差分数组
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func corpFlightBookings(bookings [][]int, n int) []int {
	// 差分数组
	var diff = make([]int, n)

	for _, booking := range bookings {
		left, right, x := booking[0], booking[1], booking[2]
		diff[left-1] += x
		if right < n {
			diff[right] -= x
		}
	}

	// 复原数组
	for i := 1; i < n; i++ {
		diff[i] += diff[i-1]
	}
	return diff
}

// 差分数组
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func corpFlightBookings2(bookings [][]int, n int) []int {
	// 初始化差分数组, 长度 +1 避免数组越界判断
	var diff = make([]int, n+1)

	for _, booking := range bookings {
		left, right, x := booking[0], booking[1], booking[2]
		diff[left-1] += x
		diff[right] -= x
	}

	// 复原数组
	for i := 1; i < n; i++ {
		diff[i] += diff[i-1]
	}
	return diff[:n]
}
