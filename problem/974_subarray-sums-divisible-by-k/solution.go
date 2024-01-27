package main

// 前缀和 + 哈希表
// - 时间复杂度：$\mathcal{O}(n)$。
// - 空间复杂度：$\mathcal{O}(n)$。
func subarraysDivByK(nums []int, k int) int {

	var n = len(nums)
	var ans = 0 // math.MaxInt32 // math.MinInt32
	var preSum = make([]int, n+1)
	for i, num := range nums {
		preSum[i+1] = preSum[i] + num
	}

	var count = map[int]int{}
	for _, sum := range preSum {
		sum = (sum%k + k) % k // 负数取模

		ans += count[sum]
		count[sum]++
	}

	return ans
}

// subarraysDivByK 的空间优化 + 一次遍历
//
// 前缀和 + 哈希表
// - 时间复杂度：$\mathcal{O}(n)$。
// - 空间复杂度：$\mathcal{O}(n)$。
func subarraysDivByK2(nums []int, k int) int {
	var ans = 0 // math.MaxInt32 // math.MinInt32
	var sum = 0

	var count = map[int]int{
		0: 1, // 初始化 preSum[0] = 0
	}
	for _, num := range nums {
		sum = ((sum+num)%k + k) % k // 负数取模
		ans += count[sum]
		count[sum]++
	}

	return ans
}
