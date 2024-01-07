package main

import "sort"

// ## 方法二：二分+贪心
//
// 也可以用贪心做。
//
// 考虑到只需要计算个数，在从左到右遍历的情况下只要当前房子可以偷，就立刻偷。
//
// 例如 $\textit{nums}=[1,2,3,4],\ mx=3$，如果不偷 $\textit{nums}[0]=1$ 去偷 $\textit{nums}[1]=2$，那么只能偷一间房子。而如果偷 $\textit{nums}[0]=1$ 和 $\textit{nums}[2]=3$，
// 就可以偷两间房子。
//
// 严格证明如下：
//
// 根据递推式 $f[i] = \max(f[i-1], f[i-2] + 1)$ 可知，
//
// $$
// f[i] \ge f[i-1]
// $$
//
// 所以 $f$ 数组是递增的。
//
// 此外，$f[i]-f[i-1]$ 不会超过 $1$，毕竟我们统计的是「个数」，从 $f[i-1]$ 到 $f[i]$ 最多增加 $1$。
//
// 因此
//
// $$
// f[i-2]+1\ge f[i-1]
// $$
//
// 必然成立。也就是说，如果 $\textit{nums}[i]\le mx$，则
//
// $$
// f[i] = f[i-2] + 1
// $$
//
// 上式表明，在从左到右遍历 $\textit{nums}$ 时，能偷就偷。如果 $\textit{nums}[i]\le mx$，我们可以偷 $\textit{nums}[i]$，并跳过 $\textit{nums}[i+1]$。
//
// #### 复杂度分析
//
// - 时间复杂度：$\mathcal{O}(n\log U)$，其中 $n$ 为 $\textit{nums}$ 的长度，$U=max(\textit{nums})$。
// - 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。
func endlessCheng(nums []int, k int) int {
	return sort.Search(1e9, func(mx int) bool {
		cnt, n := 0, len(nums)
		for i := 0; i < n; i++ {
			if nums[i] <= mx {
				cnt++ // 偷 nums[i]
				i++   // 跳过下一间房子
			}
		}
		return cnt >= k
	})
}
