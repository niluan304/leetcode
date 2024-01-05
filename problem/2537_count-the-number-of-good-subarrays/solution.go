package main

// 滑动窗口
// 时间复杂度: O(n)
// 空间复杂度: O(n)
func countGood(nums []int, k int) int64 {
	var count, sum, ans = map[int]int{}, 0, 0

	var left, n = 0, len(nums)
	for right, num := range nums {
		sum += count[num] // 4个1有 6对，5个1有 10对，增加了4对
		count[num]++
		for sum >= k {
			ans += n - right
			count[nums[left]]--
			sum -= count[nums[left]] // 6个1有15对，5个1有 10对，减少了5对
			left++
		}
	}
	return int64(ans)
}
