package main

import (
	"sort"
)

// ### 复杂度分析
//
// - 时间复杂度：$O(\mathcal{O}(n + m\log m))$，其中 $n$ 为 $s$ 的长度，其中 $m$ 为 $t$ 的长度。
// - 空间复杂度：$O(n)$。
func minimumScore(s string, t string) int {
	var n, m = len(s), len(t)
	var prefix, suffix = make([]int, m), make([]int, m)
	for i := 0; i < m; i++ { // 初始化前后缀
		prefix[i] = n
		suffix[i] = -1
	}

	var left, right = 0, m - 1
	var ans = m
	for i := 0; i < n; i++ {
		if left < m && t[left] == s[i] {
			prefix[left] = i
			left++
			ans = min(ans, m-left) // 只使用前缀匹配 s 的子序列
		}
		if right >= 0 && t[right] == s[n-1-i] {
			suffix[right] = n - 1 - i
			right--
			ans = min(ans, right+1) // 只使用后缀匹配 s 的子序列
		}
	}

	for i := 0; i < m; i++ {
		j := sort.SearchInts(suffix[i+1:], prefix[i]+1)
		if i+j+1 == m {
			continue
		}
		ans = min(ans, j)
	}

	return ans
}

// ### 提示 1
//
// 在 $[\textit{left}, \textit{right}]$ 之间的字符，删除是不影响得分的，且删除后更有机会让剩余部分是 $s$ 的子序列。
//
// 因此只需考虑删除的是 $t$ 的子串，而不是子序列。
//
// ### 提示 2
//
// 删除子串后，剩余部分是 $t$ 的一个前缀和一个后缀。
//
// 假设前缀匹配的是 $s$ 的一个前缀 $s[:i]$，后缀匹配的是 $s$ 的一个后缀 $s[i:]$。这里匹配指子序列匹配。
//
// 那么枚举 $i$，分别计算能够与 $s[:i]$ 和 $s[i:]$ 匹配的 $t$ 的最长前缀和最长后缀，就知道要删除的子串的最小值了。这个技巧叫做「前后缀分解」。
//
// ### 提示 3
//
// 具体来说：
//
// 定义 $\textit{pre}[i]$ 为 $s[:i]$ 对应的 $t$ 的最长前缀的结束下标。
//
// 定义 $\textit{suf}[i]$ 为 $s[i:]$ 对应的 $t$ 的最长后缀的开始下标。
//
// 那么删除的子串就是从 $\textit{pre}[i]+1$ 到 $\textit{suf}[i]-1$ 这段，答案就是 $\textit{suf}[i]-\textit{pre}[i]-1$ 的最小值。
//
// 代码实现时，可以先计算 $\textit{suf}$，然后一边计算 $\textit{pre}$，一边更新最小值，所以 $\textit{pre}$ 可以省略。
//
// 附：[视频讲解](https://www.bilibili.com/video/BV1GY411i7RP/)
//
// ### 复杂度分析
//
// - 时间复杂度：$O(n)$，其中 $n$ 为 $s$ 的长度。注意时间复杂度和 $t$ 的长度无关。
// - 空间复杂度：$O(n)$。
func minimumScore2(s string, t string) int {
	var n, m = len(s), len(t)
	// t 的前缀 t[:prefix[i]] 是 s[:i] 的子序列
	// t 的后缀 t[suffix[i]:] 是 s[i:] 的子序列
	var prefix, suffix = make([]int, n), make([]int, n+1)
	suffix[n] = m

	for i, j := n-1, m-1; i >= 0; i-- {
		if j >= 0 && s[i] == t[j] {
			j--
		}
		suffix[i] = j + 1
	}

	var ans = suffix[0]
	if ans == 0 {
		return 0
	}

	for i, j := 0, 0; i < n; i++ {
		if s[i] == t[j] { // 注意 j 不会等于 m，因为上面 suf[0]>0 表示 t 不是 s 的子序列
			j++
		}
		prefix[i] = j
	}

	for i := 0; i < n; i++ {
		ans = min(ans, suffix[i+1]-prefix[i]) // 删除 t[prefix[i]:suf[i+1]]
	}

	return ans
}
