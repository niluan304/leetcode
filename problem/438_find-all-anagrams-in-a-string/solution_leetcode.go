package main

// 方法一：滑动窗口
// 思路
//
// 根据题目要求，我们需要在字符串 s 寻找字符串 p 的异位词。因为字符串 p 的异位词的长度一定与字符串 p 的长度相同，
// 所以我们可以在字符串 s 中构造一个长度为与字符串 p 的长度相同的滑动窗口，并在滑动中维护窗口中每种字母的数量；
// 当窗口中每种字母的数量与字符串 p 中每种字母的数量相同时，则说明当前窗口为字符串 p 的异位词。
func leetcode1(s, p string) (ans []int) {
	sLen, pLen := len(s), len(p)
	if sLen < pLen {
		return
	}

	var sCount, pCount [26]int
	for i, ch := range p {
		sCount[s[i]-'a']++
		pCount[ch-'a']++
	}
	if sCount == pCount {
		ans = append(ans, 0)
	}

	for i, ch := range s[:sLen-pLen] {
		sCount[ch-'a']--
		sCount[s[i+pLen]-'a']++
		if sCount == pCount {
			ans = append(ans, i+1)
		}
	}
	return
}

// 方法二：优化的滑动窗口
// 思路和算法
//
// 在方法一的基础上，我们不再分别统计滑动窗口和字符串 p 中每种字母的数量，而是统计滑动窗口和字符串 p 中每种字母数量的差；
// 并引入变量 differ 来记录当前窗口与字符串 p 中数量不同的字母的个数，并在滑动窗口的过程中维护它。
//
// 在判断滑动窗口中每种字母的数量与字符串 p 中每种字母的数量是否相同时，只需要判断 differ 是否为零即可。
func leetcode2(s, p string) (ans []int) {
	sLen, pLen := len(s), len(p)
	if sLen < pLen {
		return
	}

	count := [26]int{}
	for i, ch := range p {
		count[s[i]-'a']++
		count[ch-'a']--
	}

	differ := 0
	for _, c := range count {
		if c != 0 {
			differ++
		}
	}
	if differ == 0 {
		ans = append(ans, 0)
	}

	for i, ch := range s[:sLen-pLen] {
		if count[ch-'a'] == 1 { // 窗口中字母 s[i] 的数量与字符串 p 中的数量从不同变得相同
			differ--
		} else if count[ch-'a'] == 0 { // 窗口中字母 s[i] 的数量与字符串 p 中的数量从相同变得不同
			differ++
		}
		count[ch-'a']--

		if count[s[i+pLen]-'a'] == -1 { // 窗口中字母 s[i+pLen] 的数量与字符串 p 中的数量从不同变得相同
			differ--
		} else if count[s[i+pLen]-'a'] == 0 { // 窗口中字母 s[i+pLen] 的数量与字符串 p 中的数量从相同变得不同
			differ++
		}
		count[s[i+pLen]-'a']++

		if differ == 0 {
			ans = append(ans, i+1)
		}
	}
	return
}
