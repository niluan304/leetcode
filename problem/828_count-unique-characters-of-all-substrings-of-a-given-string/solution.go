package main

import "cmp"

// 动态规划
// 时间复杂度: O(n)
// 空间复杂度: O(n)
func uniqueLetterString(s string) int {
	var n = len(s)
	var dp = make([]int, n+1)
	var idx [26][2]int
	for i := 1; i <= n; i++ {
		b := s[i-1] - 'A'
		v := idx[b]
		dp[i] = dp[i-1] + (i - v[0]) - (v[0] - v[1])
		idx[b] = [2]int{i, v[0]}
	}

	ans := Sum(dp)
	return ans
}

func Sum[S ~[]E, E cmp.Ordered](x S) E {
	m := x[0]
	for i := 1; i < len(x); i++ {
		m += x[i]
	}

	return m
}

// #### 方法一：分别计算每个字符的贡献
//
// **思路**
//
// 对于下标为 $i$ 的字符 $c_i$，当它在某个子字符串中仅出现一次时，它会对这个子字符串统计唯一字符时有贡献。
// 只需对每个字符，计算有多少子字符串仅包含该字符一次即可。
// 对于 $c_i$， 记同字符上一次出现的位置为 $c_j$，下一次出现的位置为 $c_k$，
// 那么这样的子字符串就一共有 $(c_i - c_j) \times (c_k - c_i)$ 种，
// 即子字符串的起始位置有 $c_j$（不含）到 $c_i$（含）之间这 $(c_i - c_j)$ 种可能，
// 到结束位置有 $(c_k - c_i)$ 种可能。可以预处理 $s$，
// 将相同字符的下标放入数组中，方便计算。最后对所有字符进行这种计算即可。
// 时间复杂度: O(n)
// 空间复杂度: O(n)
//
// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/count-unique-characters-of-all-substrings-of-a-given-string/solutions/1802699/tong-ji-zi-chuan-zhong-de-wei-yi-zi-fu-b-h9pj/
func uniqueLetterString2(s string) (ans int) {
	idx := map[rune][]int{}
	for i, c := range s {
		idx[c] = append(idx[c], i)
	}
	for _, arr := range idx {
		arr = append(append([]int{-1}, arr...), len(s))
		for i := 1; i < len(arr)-1; i++ {
			ans += (arr[i] - arr[i-1]) * (arr[i+1] - arr[i])
		}
	}
	return
}
