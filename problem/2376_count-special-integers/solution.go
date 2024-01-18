package main

import "strconv"

func countSpecialNumbers(n int) int {
	s := strconv.Itoa(n)
	m := len(s)
	cache := make([][1 << 10]*int, m)
	//for i, _ := range cache {
	//	cache[i] = [1024]*int{}
	//}

	// v1
	var dfs func(i, mask int, isLimit, isNum bool) (res int)
	dfs = func(i, mask int, isLimit, isNum bool) (res int) {
		if i == len(s) {
			if isNum {
				return 1
			} else {
				return 0
			}
		}

		// 问：记忆化四个状态有点麻烦，能不能只记忆化 (i,mask) 这两个状态？
		//
		// 答：是可以的。比如 n=234，第一位填 2，第二位填 3，后面无论怎么递归，都不会再次递归到第一位填 2，第二位填 3 的情况，所以不需要记录。
		// 又比如，第一位不填，第二位也不填，后面无论怎么递归也不会再次递归到这种情况，所以也不需要记录。
		if !isLimit && isNum {
			ptr := &cache[i][mask]
			if *ptr != nil {
				return **ptr
			}
			defer func() { *ptr = &res }()
		}

		if !isNum {
			res += dfs(i+1, mask, false, false)
		}

		// 处理递归
		d := 0
		if !isNum {
			d = 1 // 如果前面没有填数字，必须从 1 开始（因为不能有前导零）
		}
		up := 9
		if isLimit {
			up = int(s[i]) - '0' // 如果前面填的数字都和 n 的一样，那么这一位至多填数字 s[i]（否则就超过 n 啦）
		}
		for ; d <= up; d++ { // 枚举要填入的数字 d
			if mask>>d&1 == 0 { // d 不在 mask 中
				res += dfs(i+1, mask|1<<d, isLimit && d == up, true)
			}
		}

		return res
	}

	// 递归入口：dfs(0, 0, ture, false)，表示：
	// - 从 s[0] 开始枚举;
	// - 一开始集合中没有数字;
	// - 一开始要受到 n 的约束（否则就可以随意填）;
	// - 一开始没有填数字。
	var ans = dfs(0, 0, true, false)
	return ans
}
