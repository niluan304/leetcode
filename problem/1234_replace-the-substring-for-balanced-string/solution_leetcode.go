package main

// 方法一：滑动窗口
// 思路与算法
//
// 设 partial=n/4 ，我们选择 s 的一个子串作为「待替换子串」，
// 只有当 s 剩余的部分中 ‘Q’，‘W’，‘E’，‘R’ 的出现次数都小于等于 partial 时，
// 我们才有可能使 s 变为「平衡字符串」。
//
// 如果原始的 s 就是 「平衡字符串」，我们直接返回 0，否则我们按照以下思路求解。
//
// 从小到大枚举「待替换子串」的左端点 l，为了使得替换的长度最小，
// 我们要找到最近的右端点 r，使得去除 [l,r) 之后的剩余部分满足上述条件。
// 不难发现，随着 l 的递增，r 也是递增的。
//
// 具体的，我们使用滑动窗口来维护区间 [l,r) 之外的剩余部分中 ‘Q’，‘W’，‘E’，‘R’ 的出现次数，
// 当其中一种字符的出现次数大于 partial 时，令 s[r] 的出现次数减 1，并使得 r 向右移动一个单位。
// 该操作一直被执行，直到条件被满足或者 r 到达 s 的末尾。
//
// 如果找到了使得条件被满足的 r，我们用 r−l 来更新答案，然后令 s[l] 的出现次数加 1，
// 并使得 l 向右移动一个单位进行下一次枚举。否则，后序的 l 也将不会有合法的 r，此时我们可以直接跳出循环。
// 对于所有合法的 [l,r)，取 r−l 的最小值做为答案。
func leetcode(s string) int {
	cnt := map[byte]int{}
	for _, c := range s {
		cnt[byte(c)]++
	}
	partial := len(s) / 4
	check := func() bool {
		if cnt['Q'] > partial ||
			cnt['W'] > partial ||
			cnt['E'] > partial ||
			cnt['R'] > partial {
			return false
		}
		return true
	}

	if check() {
		return 0
	}

	res := len(s)
	r := 0
	for l, c := range s {
		for r < len(s) && !check() {
			cnt[s[r]]--
			r += 1
		}
		if !check() {
			break
		}
		res = min(res, r-l)
		cnt[byte(c)]++
	}
	return res
}

// 0 ms 的代码示例
func leetcodeMinTime(s string) int {
	cnt := ['X']int{}
	m := len(s) / 4
	for _, c := range s {
		cnt[c]++
	}

	if cnt['Q'] == m && cnt['W'] == m && cnt['E'] == m && cnt['R'] == m {
		return 0
	}

	left, ans := 0, len(s)
	for right, x := range s {
		cnt[x]--
		for cnt['Q'] <= m && cnt['W'] <= m && cnt['E'] <= m && cnt['R'] <= m {
			ans = min(ans, right-left+1)
			cnt[s[left]]++
			left++
		}
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// 2.73 MB 的代码示例
func leetcodeMinMemory(s string) int {
	n := len(s)
	m := n / 4
	mp := map[rune]int{} // 记录子串以外的字符数
	for _, r := range s {
		mp[r]++
	}
	if mp[rune('Q')] <= m && mp[rune('W')] <= m && mp[rune('E')] <= m && mp[rune('R')] <= m {
		return 0
	}
	left, ans := 0, n
	for right, x := range s {
		mp[x]--
		// 如果子串以外的字符数都小于n/4，就继续缩
		for mp[rune('Q')] <= m && mp[rune('W')] <= m && mp[rune('E')] <= m && mp[rune('R')] <= m {
			ans = min(ans, right-left+1)
			mp[rune(s[left])]++
			left++
		}
	}
	return ans
}
