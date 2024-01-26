package main

// 前缀和 + 取模转换 + 两数之和
// - 时间复杂度：$\mathcal{O}(n)$。
// - 空间复杂度：$\mathcal{O}(n)$。## 相似题目（前缀和+哈希表）
//
// 请看 [视频讲解](https://www.bilibili.com/video/BV1Nj411178Z/) 第三题。
//
// 对于本题，由于需要统计 $\textit{cnt}$，我们可以把满足 $\textit{nums}[i]\bmod \textit{modulo} = k$ 的 $\textit{nums}[i]$ 视作 $1$，不满足则视作 $0$。
//
// 如此转换后，算出 $\textit{nums}$ 的前缀和数组 $s$，那么题目中的 $\textit{cnt}\bmod \textit{modulo} = k$ 等价于
//
// $$
// (s[r+1]-s[l])\bmod \textit{modulo} = k
// $$
//
// 上式等价于（推导过程请看视频）
//
// $$
// s[l]\bmod \textit{modulo} = (s[r+1]-k)\bmod \textit{modulo}
// $$
//
// 根据上式，我们可以一边枚举 $r$，一边用一个哈希表统计有多少个 $s[r+1]\bmod \textit{modulo}$。这样可以快速知道有多少个 $(s[r+1]-k)\bmod \textit{modulo}$，也就是 $s[l]\bmod \textit{modulo}$ 的个数，把个数加到答案中。
//
// ## 相似题目（前缀和+哈希表）
//
// 推荐按照顺序完成。
//
// - [560. 和为 K 的子数组](https://leetcode.cn/problems/subarray-sum-equals-k/)
// - [974. 和可被 K 整除的子数组](https://leetcode.cn/problems/subarray-sums-divisible-by-k/)
// - [523. 连续的子数组和](https://leetcode.cn/problems/continuous-subarray-sum/)
// - [525. 连续数组](https://leetcode.cn/problems/contiguous-array/)
func countInterestingSubarrays(nums []int, mod int, k int) int64 {
	var ans = 0 // math.MaxInt32 // math.MinInt32
	var n = len(nums)

	var preSum = make([]int, n+1)
	for i, num := range nums {
		preSum[i+1] = preSum[i]
		if num%mod == k {
			preSum[i+1]++
		}
	}
	var count = map[int]int{}
	for _, sum := range preSum {
		s := sum % mod
		ans += count[(s-k+mod)%mod]
		count[s]++
	}

	return int64(ans)
}

// countInterestingSubarrays 空间优化版本
// 前缀和 + 取模转换 + 两数之和
// - 时间复杂度：$\mathcal{O}(n)$。
// - 空间复杂度：$\mathcal{O}(mod)$。
func countInterestingSubarrays2(nums []int, mod int, k int) int64 {
	var ans = 0
	var count = map[int]int{0: 1} // 把 preSum[0] = 0 计算进去
	var preSum = 0
	for _, num := range nums {
		if num%mod == k {
			preSum = (preSum + 1) % mod // 这里取模，下面 cnt[s]++ 就不需要取模了
		}

		ans += count[(preSum-k+mod)%mod] // +mod 避免减法出现负数
		count[preSum]++
	}
	return int64(ans)
}

// 暴力穷举
// - 时间复杂度：$\mathcal{O}(n^2)$。
// - 空间复杂度：$\mathcal{O}(1)$。
// Deprecated: 超时
func bruteForce(nums []int, modulo int, k int) int64 {
	var ans = 0
	var n = len(nums)
	for i := 0; i < n; i++ {
		cnt := 0
		for j := i; j < n; j++ {
			if nums[j]%modulo == k {
				cnt++
			}
			if cnt%modulo == k {
				ans++
			}
		}
	}

	return int64(ans)
}
