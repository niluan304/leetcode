package main

// 方法一：滑动窗口
// 优化
//
// 注意到每次窗口滑动时，只统计了一进一出两个字符，却比较了整个 cnt1 和 cnt2 数组。
//
// 从这个角度出发，我们可以用一个变量 diff 来记录 cnt1 与 cnt2 的不同值的个数，
// 这样判断 cnt1 和 是否相等就转换成了判断 diff 是否为 0.
func leetcode1(s1, s2 string) bool {
	n, m := len(s1), len(s2)
	if n > m {
		return false
	}
	cnt := [26]int{}
	for i, ch := range s1 {
		cnt[ch-'a']--
		cnt[s2[i]-'a']++
	}
	diff := 0
	for _, c := range cnt[:] {
		if c != 0 {
			diff++
		}
	}
	if diff == 0 {
		return true
	}
	for i := n; i < m; i++ {
		x, y := s2[i]-'a', s2[i-n]-'a'
		if x == y {
			continue
		}
		if cnt[x] == 0 {
			diff++
		}
		cnt[x]++
		if cnt[x] == 0 {
			diff--
		}
		if cnt[y] == 0 {
			diff++
		}
		cnt[y]--
		if cnt[y] == 0 {
			diff--
		}
		if diff == 0 {
			return true
		}
	}
	return false
}

// 方法二：双指针
// 回顾方法一的思路，我们在保证区间长度为 n 的情况下，去考察是否存在一个区间使得 cnt 的值全为 0。
//
// 反过来，还可以在保证 cnt 的值不为正的情况下，去考察是否存在一个区间，其长度恰好为 n。
// 初始时，仅统计 s1 中的字符，则 cnt 的值均不为正，且元素值之和为 −n。
//
// 然后用两个指针 left 和 right 表示考察的区间 [left,right]。right 每向右移动一次，就统计一次进入区间的字符 x。
// 为保证 cnt，若此时 cnt[x]>0，则向右移动左指针，减少离开区间的字符的 cnt 值直到 cnt[x]≤0。
//
// 注意到 [left,right] 的长度每增加 1，cnt 的元素值之和就增加 1。
// 当 [left,right] 的长度恰好为 n 时，就意味着 cnt 的元素值之和为 0。
// 由于 cnt 的值不为正，元素值之和为 0 就意味着所有元素均为 0，这样我们就找到了一个目标子串。
func leetcode2(s1, s2 string) bool {
	n, m := len(s1), len(s2)
	if n > m {
		return false
	}
	cnt := [26]int{}
	for _, ch := range s1 {
		cnt[ch-'a']--
	}
	left := 0
	for right, ch := range s2 {
		x := ch - 'a'
		cnt[x]++
		for cnt[x] > 0 {
			cnt[s2[left]-'a']--
			left++
		}
		if right-left+1 == n {
			return true
		}
	}
	return false
}
