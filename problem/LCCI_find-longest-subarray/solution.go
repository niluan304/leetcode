package main

// 前缀和 + 哈希表
// - 时间复杂度：$\mathcal{O}(n)$。
// - 空间复杂度：$\mathcal{O}(n)$。
func findLongestSubarray(array []string) []string {
	var n = len(array)
	var prefix = make([]int, n+1)
	for i, s := range array {
		if '0' <= s[0] && s[0] <= '9' {
			prefix[i+1] = prefix[i] + 1
		} else {
			prefix[i+1] = prefix[i] - 1
		}
	}

	l, r := 0, 0
	var count = make(map[int]int, n)
	for i, sum := range prefix {
		if j, ok := count[sum]; ok {
			if i-j > r-l { // 若存在多个最长子数组，返回左端点下标值最小的子数组
				l, r = j, i
			}
		} else {
			count[sum] = i
		}
	}
	return array[l:r]
}

// findLongestSubarray 的空间优化版本
//
// 前缀和 + 哈希表
// - 时间复杂度：$\mathcal{O}(n)$。
// - 空间复杂度：$\mathcal{O}(n)$。
func findLongestSubarray2(array []string) []string {
	var sum = 0
	var count = map[int]int{
		0: 0, // 初始化 prefix[0] = 0
	}

	l, r := 0, 0
	for i := 1; i <= len(array); i++ {
		char := array[i-1][0]
		if '0' <= char && char <= '9' {
			sum++
		} else {
			sum--
		}

		if j, ok := count[sum]; ok {
			if i-j > r-l { // 若存在多个最长子数组，返回左端点下标值最小的子数组
				l, r = j, i
			}
		} else {
			count[sum] = i
		}
	}
	return array[l:r]
}
