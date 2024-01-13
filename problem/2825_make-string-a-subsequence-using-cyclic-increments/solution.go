package main

// 双指针 + 贪心匹配
// - 时间复杂度：$\mathcal{O}(n)$。
// - 空间复杂度：$\mathcal{O}(1)$。
func canMakeSubsequence(str1 string, str2 string) bool {
	var m, n = len(str1), len(str2)
	if m < n {
		return false
	}

	var j = 0
	for i := 0; i < m && j < n; i++ {
		// 如果可以 字符 str1[i] 可以匹配 str[j] 那就马上匹配
		switch rune(str2[j]) - rune(str1[i]) {
		case 0, 1, 'a' - 'z':
			j++
		}

		//if (str1[i] == str2[j]) || (str1[i]+1 == str2[j]) || (str1[i] == 'z' && str2[j] == 'a') {
		//	j++
		//}
	}

	return j == n
}
