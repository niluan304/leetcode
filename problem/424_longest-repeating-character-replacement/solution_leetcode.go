package main

// 0 ms 的代码示例
// 2.25 MB 的代码示例
func leetcodeMinTime(s string, k int) int {
	// 记录窗口内的字符数量
	// 当窗口内的字符减去重复的字符之后>k则不能替换，此时需要收缩窗口
	l, r := 0, 0
	n := len(s)
	m := [26]int{}
	maxL := 0
	for r < n {
		cur := s[r] - 'A'
		m[cur] += 1
		maxL = max(maxL, m[cur])
		if r-l+1-maxL > k {
			m[s[l]-'A'] -= 1
			l++
		}
		r++
	}
	return r - l
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
