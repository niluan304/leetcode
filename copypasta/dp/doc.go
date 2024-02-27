/*
Package dp copy form https://github.com/EndlessCheng/codeforces-go/copypasta/dp.go

*/

package dp

var _ = []any{
	// dfs + 记忆化搜索
	MemorySearch[int, int], MemorySearch2[int, int, int],

	// 线性 DP
	// LCS 最长公共子序列

	// LIS 最长递增子序列
	LIS, LISGreedy, LISPath,

	// 状态机 DP

	// 背包 DP
	ZeroOneKnapsack, ZeroOneKnapsackExactlyFull, ZeroOneKnapsackAtLeastFillUp, ZeroOneWaysToSum, // 0-1 背包
	UnboundedKnapsack, UnboundedKnapsackExactlyFull, UnboundedKnapsackAtLeastFillUp, UnboundedKnapsackWaysToSum, // 完全背包
	GroupKnapsack, GroupKnapsackFill, GroupKnapsackWaysToSum, // 分组背包

	// 区间 DP

	// 树形 DP

	// 状态压缩 DP

	// 数位 DP
	// DigitDp, DigitDpExtraLeadingZero,
}
