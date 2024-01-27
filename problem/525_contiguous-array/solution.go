package main

// 前缀和 + 哈希表
// - 时间复杂度：$\mathcal{O}(n)$。
// - 空间复杂度：$\mathcal{O}(n)$。
func findMaxLength(nums []int) int {

	var n = len(nums)
	var preSum = make([]int, n+1)
	for i, num := range nums {
		//preSum[i+1] = preSum[i] + (num*2 - 1) // 简洁写法

		if num == 0 {
			num = -1
		}
		preSum[i+1] = preSum[i] + num
	}

	var count = map[int]int{}
	var ans = 0 // math.MaxInt32 // math.MinInt32
	for i, sum := range preSum {
		j, ok := count[sum]
		if ok {
			ans = max(ans, i-j)
		} else {
			count[sum] = i
		}
	}

	return ans
}

// findMaxLength 的空间优化
//
// 前缀和 + 哈希表
// - 时间复杂度：$\mathcal{O}(n)$。
// - 空间复杂度：$\mathcal{O}(n)$。
func findMaxLength2(nums []int) int {
	var sum = 0
	var count = map[int]int{
		0: -1, // 初始化 preSum[0] = 0
	}

	var ans = 0
	for i, num := range nums {
		sum += num*2 - 1

		j, ok := count[sum]
		if ok {
			ans = max(ans, i-j)
		} else {
			count[sum] = i
		}
	}
	return ans
}
