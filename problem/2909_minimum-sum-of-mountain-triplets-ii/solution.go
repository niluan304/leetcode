package main

import (
	"math"
)

// 请看 [视频讲解](https://www.bilibili.com/video/BV12w411B7ia/)。
//
// 遇到这种三元组的题目，一个通常的做法是枚举中间的数。
//
// 知道了 $\textit{nums}[j]$，只需要知道 $j$ 左边的最小值和右边的最小值，就知道了三元组的和的最小值。
//
// 左右的最小值可以递推算出来。定义 $\textit{suf}[i]$ 表示从 $\textit{nums}[i]$ 到 $\textit{nums}[n-1]$ 的最小值（后缀最小值），则有
//
// $$
// \textit{suf}[i] = \min(\textit{suf}[i+1], \textit{nums}[i])
// $$
//
// 前缀最小值 $\textit{pre}$ 的计算方式同理，可以和答案一起算，所以只需要一个变量。
//
// 那么答案就是
//
// $$
// \textit{pre} + \textit{nums}[j] + \textit{suf}[j+1]
// $$
//
// 的最小值。
//
// #### 复杂度分析
//
// - 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
// - 空间复杂度：$\mathcal{O}(n)$。值 pre\textit{pre}pre 的计算方式同理，可以和答案一起算，所以只需要一个变量。
//
// 作者：灵茶山艾府
// 链接：https://leetcode.cn/problems/minimum-sum-of-mountain-triplets-ii/solutions/2493548/mei-ju-numsj-qian-hou-zhui-fen-jie-pytho-tskf/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func minimumSum(nums []int) int {
	var n = len(nums)
	var prefix, suffix = make([]int, n), make([]int, n)
	prefix[0], suffix[n-1] = math.MaxInt32, math.MaxInt32

	for i := 0; i < n-1; i++ {
		prefix[i+1] = min(prefix[i], nums[i])
	}
	for i := n - 1; i >= 1; i-- {
		suffix[i-1] = min(suffix[i], nums[i])
	}
	var ans = math.MaxInt32
	for i := 0; i < n; i++ {
		if nums[i] > prefix[i] && nums[i] > suffix[i] {
			ans = min(ans, prefix[i]+nums[i]+suffix[i])
		}
	}

	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}
