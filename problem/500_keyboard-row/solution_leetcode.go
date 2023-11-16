package main

import "unicode"

// 方法一：遍历
// 我们为每一个英文字母标记其对应键盘上的行号，然后检测字符串中所有字符对应的行号是否相同。
// 我们可以预处理计算出每个字符对应的行号。
// 遍历字符串时，统一将大写字母转化为小写字母方便计算。
func leetcode1(words []string) (ans []string) {
	const rowIdx = "12210111011122000010020202"
next:
	for _, word := range words {
		idx := rowIdx[unicode.ToLower(rune(word[0]))-'a']
		for _, ch := range word[1:] {
			if rowIdx[unicode.ToLower(ch)-'a'] != idx {
				continue next
			}
		}
		ans = append(ans, word)
	}
	return
}
