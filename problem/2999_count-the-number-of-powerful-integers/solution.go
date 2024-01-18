package main

import (
	"strconv"
	"strings"
)

func numberOfPowerfulInt(start int64, finish int64, limit int, s string) int64 {
	startStr := strconv.FormatInt(start-1, 10)
	finishStr := strconv.FormatInt(finish, 10)
	x, y := digitalDP(finishStr, limit, s), digitalDP(startStr, limit, s)
	return x - y
}

func digitalDP(s string, limit int, suffix string) int64 {
	m, n := len(s), len(suffix)
	if m < n || (m == n && s < suffix) {
		return 0
	}

	type Key struct {
		index          int
		isLimit, isNum bool
	}
	cache := make(map[Key]int, m)
	var dfs func(i int, isLimit, isNum bool) (res int)
	dfs = func(i int, isLimit, isNum bool) (res int) {
		key := Key{index: i, isLimit: isLimit, isNum: isNum}
		if v, ok := cache[key]; ok {
			return v
		}
		defer func() { cache[key] = res }()

		if i == m {
			if isNum {
				return 1
			}
			return 0
		}

		// 问：记忆化四个状态有点麻烦，能不能只记忆化 (i,mask) 这两个状态？
		//
		// 答：是可以的。比如 n=234，第一位填 2，第二位填 3，后面无论怎么递归，都不会再次递归到第一位填 2，第二位填 3 的情况，所以不需要记录。
		// 又比如，第一位不填，第二位也不填，后面无论怎么递归也不会再次递归到这种情况，所以也不需要记录。

		// 处理递归

		up := 9
		if isLimit {
			up = int(s[i] - '0')
		}

		// 不用担心前导 0
		for d := 0; d <= limit && d <= up; d++ { // 枚举要填入的数字 d
			if j := i - (m - n); j >= 0 && d != int(suffix[j]-'0') {
				continue
			}

			v := dfs(i+1, isLimit && d == up, true)
			res += v
		}

		return res
	}

	// 递归入口：dfs(0, 0, ture, false)，表示：
	// - 从 s[0] 开始枚举;
	// - 一开始集合中没有数字;
	// - 一开始要受到 n 的约束（否则就可以随意填）;
	// - 一开始没有填数字。
	var ans = dfs(0, true, false)
	return int64(ans)
}

func numberOfPowerfulInt2(start int64, finish int64, limit int, s string) int64 {
	high := strconv.FormatInt(finish, 10)
	low := strconv.FormatInt(start, 10)

	var n = len(high)
	low = strings.Repeat("0", n-len(low)) + low // 前导 0 方便判断
	var cache = make([]*int, n)

	// v2基础版：不支持前导零
	var dfs func(i int, limitLow, limitHigh bool) (res int)
	dfs = func(i int, limitLow, limitHigh bool) (res int) {
		if i == n {
			return 1
		}

		if !limitLow && !limitHigh {
			ptr := &cache[i]
			if *ptr != nil {
				return **ptr
			}
			defer func() { *ptr = &res }()
		}

		// 第 i 个数位可以从 lo 枚举到 hi
		// 如果对数位还有其它约束，应当只在下面的 for 循环做限制，不应修改 lo 或 hi
		lo, hi := 0, 9
		if limitLow {
			lo = int(low[i] - '0')
		}
		if limitHigh {
			hi = int(high[i] - '0')
		}

		for d := lo; d <= hi; d++ { // 枚举这个数位填什么

			if d > limit { // 本题的 limit 约束
				break
			}
			if diff := n - len(s); i >= diff && d != int(s[i-diff]-'0') { // 本题的后缀判断，这个数位只能填 s[i-diff]
				continue
			}

			v := dfs(i+1, limitLow && d == lo, limitHigh && d == hi)
			res += v
		}

		return res
	}

	var ans = dfs(0, true, true)
	return int64(ans)
}
