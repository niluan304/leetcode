package copypasta

import "math"

// KMP
// implement Knuth–Morris–Pratt algorithm
// search pattern from text, return all start positions
//
// kmp 在求 前缀表 / next 数组时，借助了动态规划的思想；
// 在求模式串的索引位置时，则是贪心的思想。
//
// [前缀函数与 kmp 算法 - OI Wiki](https://oi.wiki/string/kmp)
// [如何更好地理解和掌握 kmp 算法? - 阮行止的回答](https://www.zhihu.com/question/21923021/answer/1032665486)
func KMP(text, patten []rune) (pos []int) {
	m, n := len(text), len(patten)
	if m < n {
		return nil
	}

	match := MaxPrefix(patten)

	// 第一种写法
	pos = make([]int, 0, m)
	for i, j := 0, 0; i < m; {
		if patten[j] == text[i] {
			// 字符匹配的情况，text, pattern 都向前移动 1 位
			i++
			j++
		} else {
			// 字符不匹配的情况
			if j > 0 { // 回退 pattern
				j = match[j-1]
			} else { // pattern 已经到头了，只能 text 向前移动 1 位
				i++
			}
		}

		if j == n {
			pos = append(pos, i-n)
			j = match[j-1]
		}
	}

	// 第二种写法
	pos = make([]int, 0, m)
	for i, j := 0, 0; i < m; i++ {
		v := text[i]

		// 字符串不相等的情况，j 回退到 x 位置，直到 x 满足 text[i-x:i] == patten[:x]，最多回退到 x == 0，此时两者都为空字符串
		for j > 0 && patten[j] != v {
			j = match[j-1]
		}
		if patten[j] == v { // 两个字符相等，i, j 都都向前移动
			j++
		}
		if j == n {
			pos = append(pos, i-n+1)
			j = match[j-1]
		}
	}
	return pos
}

// MaxPrefix 前缀表
//
// match[i] 表示在子串 str[:i] 中「真前缀」和「真后缀」一致时的最大长度，「真前缀」/「后前缀」指不包含自身的 前缀/后缀。
// 或者说在子串 str[:i-1] 中求「前缀」和「后缀」一致时的最大长度
func MaxPrefix(str []rune) (match []int) {
	n := len(str)

	match = make([]int, n)
	for i := 1; i < n; i++ {
		v, j := str[i], match[i-1]

		// 注意到 match[i+1] <= match[i] + 1，
		// 通俗的讲，如果有 match[i+1] == match[i] + 1，我们叫上升，反之叫下降，
		// 又因为每次最多上升一次，而下降的总数必然是不会超过上升的总数的（不然 match[i] 的值就要变负了），
		// 所以总体上升加下降的次数不会超过 2n，所以算法是 O(n) 的
		for j > 0 && str[j] != v {
			j = match[j-1] // 转移方程
		}
		if str[j] == v {
			j++
		}
		match[i] = j
	}
	return match
}

// KMP2
// 前缀函数的运用
//
// 给定一个文本 t 和一个字符串 s，我们尝试找到并展示 s 在 t 中的所有出现（occurrence）。
// 为了简便起见，我们用 n 表示字符串 s 的长度，用 m 表示文本 t 的长度。
// 我们构造一个字符串 s+#+t，其中 # 为一个既不出现在 s 中也不出现在 t 中的分隔符。
//
// 接下来计算该字符串的前缀函数。
// 现在考虑该前缀函数除去最开始 n+1个值（即属于字符串s和分隔符的函数值）后其余函数值的意义。
// 根据定义，match[i] 为右端点在且同时为一个前缀的最长真子串的长度，具体到我们的这种情况下，
// 其值为与的前缀相同且右端点位于的最长子串的长度。由于分隔符的存在，该长度不可能超过 n。
// 而如果等式 match[i]= n 成立，则意味着 s 完整出现在该位置（即其右端点位于位置i）。
//
// 注意该位置的下标是对字符串 s+#+t而言的。
// 因此如果在某一位置有 match[i]=n 成立，则字符串 s在字符串t的 i - (n-1)-(n+1) = i - 2n处出现。
func KMP2(text, patten []rune) (pos []int) {
	m, n := len(text), len(patten)

	cur := append(patten, math.MaxInt32)
	cur = append(cur, text...)
	match := MaxPrefix(cur)
	for i := n + 1; i < n+m+1; i++ { // i 的起点可以改为 2*n，从 n+1 开始是为了方便翻译
		if match[i] == n {
			pos = append(pos, i-2*n)
		}
	}
	return pos
}
