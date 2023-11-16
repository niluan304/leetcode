package main

// 方法一：哈希表
// 思路和算法
//
// 我们考虑枚举数组中的每个数 x，考虑以其为起点，不断尝试匹配 x+1,x+2,⋯ 是否存在，
// 假设最长匹配到了 x+y，那么以 x 为起点的最长连续序列即为 x,x+1,x+2,⋯, x+y，
// 其长度为 y+1，我们不断枚举并更新答案即可。
func leetcode1(nums []int) int {
	numSet := map[int]bool{}
	for _, num := range nums {
		numSet[num] = true
	}
	longestStreak := 0
	for num := range numSet {
		if !numSet[num-1] {
			currentNum := num
			currentStreak := 1
			for numSet[currentNum+1] {
				currentNum++
				currentStreak++
			}
			if longestStreak < currentStreak {
				longestStreak = currentStreak
			}
		}
	}
	return longestStreak
}
