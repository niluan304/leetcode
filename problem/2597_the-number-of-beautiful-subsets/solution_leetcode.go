package main

import "sort"

// # 方法二：动态规划
//
// ### 前置知识：动态规划基础（打家劫舍）
//
// 见[【基础算法精讲 17】](https://www.bilibili.com/video/BV1Xj411K7oF/)。
//
// ### 前置知识：乘法原理
//
// 见 [乘法原理](https://baike.baidu.com/item/%E4%B9%98%E6%B3%95%E5%8E%9F%E7%90%86/7538447)。
//
// ### 前置知识：同余
//
// 两个数 $x$ 和 $y$，如果 $(x-y)\bmod k = 0$，则称 $x$ 与 $y$ 对模 $k$ 同余，记作
//
// $$
// x\equiv y \pmod k
// $$
//
// 例如 $42\equiv 12 \pmod {10}$，$-17\equiv 3 \pmod {10}$。
//
// ### 思路
//
// 如果两个数模 $k$ **不同余**，那么必然无法相差 $k$。
//
// 所以我们可以按照模 $k$ 的结果分组，每一组用哈希表/有序集合统计元素及其出现次数。
//
// 每一组怎么思考呢？
//
// 按照 key 从小到大排序后（设这些 key 组成了数组 $g$），相邻的 key 如果相差 $k$，那么不能同时选（类似 [198. 打家劫舍](https://leetcode.cn/problems/house-robber/)）。
//
// 为什么不考虑非相邻的 key？因为这个组里面的 key 都是模 $k$ 同余的，非相邻的 key 相差大于 $k$。
//
// 设 $g$ 的大小为 $m$。考虑最大的数 $g[m-1]$「选或不选」：
//
// - 如果不选 $g[m-1]$，那么问题变成一个 $m-1$ 个数的子问题。
// - 如果选 $g[m-1]$：
// - 这有 $2^c-1$ 种方案，这里 $c$ 为 $g[m-1]$ 的出现次数；
// - 如果 $g[m-1]-g[m-2] = k$，那么 $g[m-2]$ 绝对不能选，问题变成一个 $m-2$ 个数的子问题。
// - 如果 $g[m-1]-g[m-2] \ne k$，那么 $g[m-2]$ 可选可不选，问题变成一个 $m-1$ 个数的子问题。
//
// 定义 $f[i]$ 表示考虑前 $i$ 个 key 的方案数，可以得到一个类似打家劫舍的转移方程：
//
// - 如果 $g[i]-g[i-1]=k$，那么 $f[i]=f[i-1]+f[i-2] \cdot( 2^{c_i}-1)$。
// - 如果 $g[i]-g[i-1]\ne k$，那么 $f[i]=f[i-1]\cdot 2^{c_i}$。
//
// 其中 $c_i$ 为 $g[i]$ 的出现次数。
//
// 代码实现时，为了避免负数下标，需要偏移一位。
//
// 每组的初始值为 $f[0]=1$，$f[1] = 2^{c_0}$。
//
// 每组的答案为 $f[m]$（因为偏移了一位）。
//
// 根据乘法原理，最终答案为每组的答案的乘积。注意把空集去掉。
//
// ### 复杂度分析
//
// - 时间复杂度：$O(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
// - 空间复杂度：$O(n)$。
func beautifulSubsets3(nums []int, k int) int {
	groups := map[int]map[int]int{}
	for _, x := range nums {
		if groups[x%k] == nil {
			groups[x%k] = map[int]int{}
		}
		groups[x%k][x]++
	}
	ans := 1
	for _, cnt := range groups {
		m := len(cnt)
		type pair struct{ x, c int }
		g := make([]pair, 0, m)
		for x, c := range cnt {
			g = append(g, pair{x, c})
		}
		sort.Slice(g, func(i, j int) bool { return g[i].x < g[j].x })
		f := make([]int, m+1)
		f[0] = 1
		f[1] = 1 << g[0].c
		for i := 1; i < m; i++ {
			if g[i].x-g[i-1].x == k {
				f[i+1] = f[i] + f[i-1]*(1<<g[i].c-1)
			} else {
				f[i+1] = f[i] << g[i].c
			}
		}
		ans *= f[m]
	}
	return ans - 1 // -1 去掉空集
}
