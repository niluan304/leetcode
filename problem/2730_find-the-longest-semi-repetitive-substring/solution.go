package main

func longestSemiRepetitiveSubstring(s string) int {
	var left, right = 0, 0
	var n = len(s)
	var cnt, ans = 0, 0
	for ; right < n; right++ {
		if right > 0 && s[right] == s[right-1] {
			cnt++
		}
		for cnt > 1 {
			left++
			if s[left] == s[left-1] {
				cnt--
			}
		}
		ans = max(ans, right-left+1)
	}
	return ans
}
