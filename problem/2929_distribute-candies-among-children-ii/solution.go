package main

// 枚举糖果A 的数量，再计算 (B, C) 的合法方案数
//
// 仿照枚举糖果A 的思路，可以继续枚举糖果B 的合法数量，这时候已经确定糖果A,B 的数量了，那糖果C的数量也确定了。
//
// 假设 B+C = x，那么有
// 0 <= B <= min(x, limit)
// 0 <= C <= min(x, limit)
// 代入 B+C = x，有 B 的取值范围：
// [L, R]  L = max(0, x-limit), R = min(limit, n)
// 那么合法的方案数为 R-L+1
//
// - 时间复杂度：$\mathcal{O}(n)$。
// - 空间复杂度：$\mathcal{O}(1)$。
func distributeCandies(n int, limit int) int64 {
	if n > limit*3 {
		return 0
	}

	var ans = 0
	for i := 0; i <= min(limit, n); i++ {
		x := n - i
		r, l := min(limit, x), max(0, x-limit)
		ans += max(0, r-l+1)
	}

	return int64(ans)
}

// 要计算合法方案数（每个小朋友分到的糖果都不超过 $\textit{limit}$），可以先计算所有方案数（没有 $\textit{limit}$ 限制），再减去不合法的方案数（至少一个小朋友分到的糖果超过
// $\textit{limit}$）。
//
// #### 所有方案数
//
// 相当于把 $n$ 个无区别的小球放入 $3$ 个有区别的盒子，允许空盒的方案数。
//
// 隔板法：假设 $n$ 个球和 $2$ 个隔板放到 $n+2$ 个位置，第一个隔板前的球放入第一个盒子，第一个隔板和第二个隔板之间的球放入第二个盒子，第二个隔板后的球放入第三个盒子。那么从
// $n+2$ 个位置中选 $2$ 个位置放隔板，有 $C(n+2, 2)$ 种放法。注意隔板可以放在最左边或最右边，也可以连续放，对应着空盒的情况。
//
// #### 至少一个小朋友分到的糖果超过 limit
//
// 设三个小朋友分别叫 $A,B,C$。
//
// 只关注 $A$。如果 $A$ 分到的糖果超过 $\textit{limit}$，那么他至少分到 $\textit{limit}+1$ 颗糖果，问题变成剩下 $n-(\textit{limit}+1)$ 颗糖果分给三个小朋友的方案数（注意 $B$
// 和 $C$ 分到的糖果是否超过 $\textit{limit}$ 我们是不关注的），即 $C(n-(\textit{limit}+1)+2, 2)$。
//
// 只关注 $B$ 的情况和只关注 $C$ 的情况同上，均为 $C(n-(\textit{limit}+1)+2, 2)$。
//
// 直接加起来，就是 $3\cdot C(n-(\textit{limit}+1)+2, 2)$，但这样就重复统计了「至少两个小朋友分到的糖果超过 $\textit{limit}$」的情况，要减去。
//
// #### 至少两个小朋友分到的糖果超过 limit
//
// 只关注 $A$ 和 $B$。如果他们俩分到的糖果超过 $\textit{limit}$，那么至少分出去了 $2\cdot (\textit{limit}+1)$ 颗糖果，问题变成剩下 $n-2\cdot (\textit{limit}+1)$ 颗糖果分给三
// 个小朋友的方案数（注意 $C$ 分到的糖果是否超过 $\textit{limit}$ 我们是不关注的），即 $C(n-2\cdot(\textit{limit}+1)+2, 2)$。
//
// 只关注 $A,C$ 的情况和只关注 $B,C$ 的情况同上，均为 $C(n-2\cdot(\textit{limit}+1)+2, 2)$。
//
// 直接加起来，就是 $3\cdot C(n-2\cdot(\textit{limit}+1)+2, 2)$，但这样就重复统计了「三个小朋友分到的糖果均超过 $\textit{limit}$」的情况，要减去。
//
// #### 三个小朋友分到的糖果超过 limit
//
// 至少分出去了 $3\cdot (\textit{limit}+1)$ 颗糖果，问题变成剩下 $n-3\cdot (\textit{limit}+1)$ 颗糖果分给三个小朋友的方案数，即 $C(n-3\cdot(\textit{limit}+1)+2, 2)$。
//
// #### 总结
//
// 不合法的方案数为「至少一个」减去「至少两个」加上「三个」，这就是**容斥原理**。
//
// 最后用所有方案数减去不合法的方案数，整理得到答案：
//
// $$
// C(n+2, 2) - 3\cdot C(n-\textit{limit}+1, 2) + 3\cdot C(n-2\cdot\textit{limit}, 2) - C(n-3\cdot \textit{limit}-1, 2)
// $$
//
// #### 复杂度分析
//
// - 时间复杂度：$\mathcal{O}(1)$。
// - 空间复杂度：$\mathcal{O}(1)$。
//
// 作者：灵茶山艾府
// 链接：https://leetcode.cn/problems/distribute-candies-among-children-ii/solutions/2522969/o1-rong-chi-yuan-li-pythonjavacgo-by-end-2woj
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func distributeCandies2(n int, limit int) int64 {
	return c2(n+2) - 3*c2(n-limit+1) + 3*c2(n-2*limit) - c2(n-3*limit-1)
}

func c2(n int) int64 {
	if n < 2 {
		return 0
	}
	return int64(n) * int64(n-1) / 2
}

func distributeCandies3(n int, limit int) int64 {
	ans := combination(n+2, 2) -
		3*combination(n-limit+1, 2) +
		3*combination(n-2*limit, 2) -
		combination(n-3*limit-1, 2)
	return int64(ans)
}

// 计算组合数
// 需要记忆化，否则会超时
func combination(n, k int) int {
	if n < k {
		return 0
	}

	// 基本情况
	if k == 0 || k == n {
		return 1
	}

	// 递归计算组合数
	return combination(n-1, k-1) + combination(n-1, k)
}

// 计算组合数
func combination2(n, k int) int {
	if n < k {
		return 0
	}

	x, y := 1, 1
	for i := n - k + 1; i <= n; i++ {
		x *= i
	}
	for i := 1; i <= k; i++ {
		y *= i
	}
	return x / y
}
