package main

// 0 ms 的代码示例
func leetcodeMinTime(s string) int {
	var (
		l, r int
		mp   [256]int
		res  int
	)

	for ; r < len(s); r++ {
		mp[s[r]]++
		for mp[s[r]] > 1 {
			mp[s[l]]--
			l++
		}

		res = max(res, r-l+1)
	}

	return res
}
