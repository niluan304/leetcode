package main

// 滑动窗口, 使用哈希表记录当前子串中各字符的重复次数
func characterReplacement(s string, k int) int {
	var (
		set  = map[byte]int{}
		left = 0
		ans  = 0
	)

	NeedCut := func() bool {
		var (
			longest = 0
			count   = 0
		)
		for _, n := range set {
			count += n
			longest = max(n, longest)
		}

		count -= longest
		// 被替换的数量应当小于k
		return count > k
	}

	for right := 0; right < len(s); right++ {
		set[s[right]]++

		for NeedCut() && left <= right {
			set[s[left]]--
			if set[s[left]] == 0 {
				delete(set, s[left])
			}
			left++
		}

		ans = max(ans, right-left+1)
	}

	return ans
}

// 滑动窗口, 记录每次子串中最大重复字符次数
func characterReplacement2(s string, k int) int {
	var (
		ans   = 0         //
		count = [26]int{} // 记录当前子串中, 各字符的重复次数
		left  = 0         // 滑动窗口的左边界
		most  = 0         // 记录当前子串中, 最大重复字符次数
	)

	for right := range s {
		x := s[right] - 'A'
		count[x]++

		// 更新最大重复字符次数,
		// most代表这 s[left:right]中重复次数最多的字符 s[m],
		// 当前子串 s[left:right+1] 的重复次数最多的字符, 要么是count[x]++后的s[right]字符, 要么是 s[m]字符
		most = max(most, count[x])

		// 要替换的字符串数量: 当前子串长度 - 字符最大出现次数, 不能超过k
		if right-left+1-most > k {
			count[s[left]-'A']--
			left++
		}

		ans = max(ans, right-left+1)
	}

	return ans
}

// 穷举法, 时间复杂度: O(N^3)
func characterReplacement3(s string, k int) int {
	var ans = 0

	for i := range s {
		var count = [26]int{}
		// 枚举所有子字符串 O(N^2)
		for j := i; j < len(s); j++ {
			count[s[j]-'A']++
			var n = 0

			// 计算子串中的字符次数 O(N)
			for _, c := range count {
				n = max(n, c)
			}

			l := j - i + 1
			if l > k+n {
				break
			}

			ans = max(ans, l)
		}
	}
	return ans
}
