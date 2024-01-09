package main

// ## 前置知识：动态规划入门
//
// 详见 [动态规划入门：从记忆化搜索到递推【基础算法精讲 17】](https://www.bilibili.com/video/BV1Xj411K7oF/)。
//
// > APP 用户需要分享到微信打开视频链接。
//
// ## 一、寻找子问题
//
// > 为了方便转成递推，从后往前考虑。
//
// 设 $n$ 为 $s$ 的长度。我们可以：
//
// - 直接跳过 $s$ 的最后一个字符，那么问题变成 $s$ 的前 $n-1$ 个字符的子问题。
// - 考虑「枚举选哪个」，如果从 $s[j]$ 开始的后缀在 $\textit{dictionary}$ 中，那么问题变成 $s$ 的前 $j-1$ 个字符的子问题。
//
// ## 二、记忆化搜索
//
// 根据上面的讨论，定义 $\textit{dfs}(i)$ 表示 $s$ 的前 $i$ 个字符的子问题。
//
// - 跳过 $s$ 的最后一个字符，有 $\textit{dfs}(i)=\textit{dfs}(i-1)+1$。
// - 考虑「枚举选哪个」，如果从 $s[j]$ 到 $s[i]$ 的子串在 $\textit{dictionary}$ 中，有
//
// $$
// \textit{dfs}(i)=\min_{j=0}^{i}\textit{dfs}(j-1)
// $$
//
// 这两种情况取最小值。
//
// 递归边界：$\textit{dfs}(-1)=0$。
//
// 答案：$\textit{dfs}(n-1)$。
// #### 复杂度分析
//
// - 时间复杂度：$\mathcal{O}(L + n^3)$，其中 $n$ 为 $s$ 的长度，$L$ 为 $\textit{dictionary}$ 所有字符串的长度之和。预处理哈希表需要 $\mathcal{O}(L)$ 的时间。动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题中状态个数等于 $\mathcal{O}(n)$，单个状态的计算时间为 $\mathcal{O}(n^2)$，因此时间
// 复杂度为 $\mathcal{O}(n^3)$。所以总的时间复杂度为 $\mathcal{O}(L + n^3)$。
// - 空间复杂度：$\mathcal{O}(n+L)$。
func endlessCheng(s string, dictionary []string) int {
	has := map[string]bool{}
	for _, s := range dictionary {
		has[s] = true
	}
	n := len(s)
	memo := make([]int, n)
	for i := range memo {
		memo[i] = -1
	}
	var dfs func(int) int
	dfs = func(i int) int {
		if i < 0 {
			return 0
		}
		p := &memo[i]
		if *p != -1 {
			return *p
		}

		// 不选
		res := dfs(i-1) + 1

		// 枚举选哪个
		for j := 0; j <= i; j++ {
			if has[s[j:i+1]] {
				res = min(res, dfs(j-1))
			}
		}

		*p = res
		return res
	}
	return dfs(n - 1)
}
