package main

// ## 贪心 + 暴力穷举
//
// - 时间复杂度：$\mathcal{O}(n \cdot \Sigma)$，其中 $n$ 为 $s$ 的长度，$\Sigma = 26$ 为字符集的大小。
// - 空间复杂度：$\mathcal{O}(n)$。
func repeatLimitedString(s string, repeatLimit int) string {
	var count [26]int
	for _, b := range s {
		count[b-'a']++
	}

	var ans = []rune{-1} // 设置哨兵，避免数组越界
	var repeat = 0       // 末尾字符的使用次数

	// 寻找下一个添加的字符：使用次数小于限制次数，且字典序最大的字符
	// 最大的添加次数为 len(s)
	for i := 0; i < len(s); i++ {
		for char := 'z'; char >= 'a'; char-- {
			if count[char-'a'] == 0 {
				continue
			}

			if char == ans[len(ans)-1] {
				// 当前字符与末尾字符相同，先检查该字符的使用次数
				if repeat == repeatLimit {
					// 如果使用次数已达到限制次数，那说明要使用下一个字典序最大的字符
					continue
				} else {
					// 如果使用次数小于限制次数，末尾字符的使用次数 +1，并添加该字符
					repeat++
				}
			} else {
				// 当前字符与末尾字符不相同，重置使用次数，并添加该字符
				repeat = 1
			}

			count[char-'a']--
			ans = append(ans, char)
			break // 完成了本轮添加字符，开始下一轮查找
		}
	}
	return string(ans[1:])
}
