package main

// 模拟
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func minCount(coins []int) int {
	var ans = 0
	for _, coin := range coins {
		ans += coin/2 + coin%2
	}
	return ans
}

// 模拟
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func minCount2(coins []int) int {
	var ans = 0
	for _, coin := range coins {
		//ans += (coin + 1) / 2
		ans += (coin + 1) >> 1
	}
	return ans
}
