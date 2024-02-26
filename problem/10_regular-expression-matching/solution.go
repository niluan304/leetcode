package main

import (
	"regexp"
	"strings"

	. "github.com/niluan304/leetcode/copypasta/dp"
)

// 1 <= s.length <= 20
// 1 <= p.length <= 20
// s 只包含从 a-z 的小写字母。
// p 只包含从 a-z 的小写字母，以及字符 . 和 *。
// 保证每次出现字符 * 时，前面都匹配到有效的字符

// dfs + 记忆化搜索
// - 时间复杂度：$\mathcal{O}(n \cdot m)$。
// - 空间复杂度：$\mathcal{O}(n \cdot m)$。
func isMatch(s string, p string) bool {

	var n, m = len(s), len(p)
	var ans bool

	var dfs func(i, j int) bool
	dfs = func(i, j int) bool {
		if s[i:] == p[j:] {
			return true
		}

		if j <= m-2 && p[j+1] == '*' {
			if dfs(i, j+2) { // 0次
				return true
			}

			for k := i; k < n; k++ { // 多次
				if p[j] != '.' && s[k] != p[j] {
					break
				}
				if dfs(k+1, j+2) {
					return true
				}
			}
			return false
		}
		if i < n && j < m && (p[j] == '.' || s[i] == p[j]) {
			if dfs(i+1, j+1) {
				return true
			}
		}
		return false
	}

	MemorySearch2(&dfs)
	ans = dfs(0, 0)
	return ans
}

// 正则表达式
// go的正则保证时间复杂度是线性的，极端情况下也不会是指数级别，因此并不会 timeout
//
// - 时间复杂度：$\mathcal{O}(n)$。
// - 空间复杂度：$\mathcal{O}(1)$。
func isMatch2(s, p string) bool {
	// 增加前后缀，避免 * 结尾的字符串
	fields := strings.FieldsFunc("^"+p+"$", func(r rune) bool {
		// strings.FieldsFunc 通过指定的 rune 分割字符串
		// 将连续的 * 替换为单个 *
		return r == '*'
	})

	exp := strings.Join(fields, "*") // 修复之后的正则表达式
	return regexp.MustCompile(exp).MatchString(s)
}
