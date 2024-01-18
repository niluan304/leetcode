package main

import (
	"strconv"
	"strings"
)

func atMostNGivenDigitSet(digits []string, n int) int {
	s := strconv.Itoa(n)
	m := len(s)
	cache := make([]*int, m)

	var dfs func(i int, isLimit, isNum bool) (res int)
	dfs = func(i int, isLimit, isNum bool) (res int) {
		if i == len(s) {
			if isNum {
				return 1 // 如果填了数字，则为 1 种合法方案
			}
			return 0
		}

		// 问：记忆化四个状态有点麻烦，能不能只记忆化 (i,mask) 这两个状态？
		//
		// 答：是可以的。比如 n=234，第一位填 2，第二位填 3，后面无论怎么递归，都不会再次递归到第一位填 2，第二位填 3 的情况，所以不需要记录。
		// 又比如，第一位不填，第二位也不填，后面无论怎么递归也不会再次递归到这种情况，所以也不需要记录。
		if !isLimit && isNum { // 在不受到任何约束的情况下，返回记录的结果，避免重复运算
			ptr := &cache[i]
			if *ptr != nil {
				return **ptr
			}
			defer func() { *ptr = &res }()
		}

		if !isNum { // 前面不填数字，那么可以跳过当前数位，也不填数字
			// isLimit 改为 false，因为没有填数字，位数都比 n 要短，自然不会受到 n 的约束
			// isNum 仍然为 false，因为没有填任何数字
			res += dfs(i+1, false, false)
		}

		// 处理递归
		// 根据是否受到约束，决定可以填的数字的上限
		up := byte('9')
		if isLimit {
			up = s[i]
		}
		// 只能使用 digits 构建第 i 位数字
		// 注意：对于一般的题目而言，如果此时 isNum 为 false，则必须从 1 开始枚举，由于本题 digits 没有 0，所以无需处理这种情况
		for _, digit := range digits {
			d := digit[0]
			if d > up { // d 超过上限，由于 digits 是有序的，后面的 d 都会超过上限，故退出循环
				break
			}

			// isLimit：如果当前受到 n 的约束，且填的数字等于上限，那么后面仍然会受到 n 的约束
			// isNum 为 true，因为填了数字
			res += dfs(i+1, isLimit && d == s[i], true)
		}

		return res
	}

	// 递归入口：dfs(0, ture, false)，表示：
	// - 从 s[0] 开始枚举;
	// - 一开始要受到 n 的约束（否则就可以随意填）;
	// - 一开始没有填数字。
	var ans = dfs(0, true, false)
	return ans
}

func atMostNGivenDigitSet2(digits []string, n int) int {
	high := strconv.FormatInt(int64(n), 10)
	low := strconv.FormatInt(1, 10)

	var m = len(high)
	low = strings.Repeat("0", m-len(low)) + low // 前导 0 方便判断

	var cache = make([]*int, m)

	var set = make([]bool, 10)
	for _, digit := range digits {
		set[digit[0]-'0'] = true
	}

	// v2拓展版本：支持前导零
	var dfs func(i int, limitLow, limitHigh bool, isNum bool) (res int)
	dfs = func(i int, limitLow, limitHigh bool, isNum bool) (res int) {
		if i == m {
			if isNum {
				return 1
			}
			return 0
		}

		if !limitLow && !limitHigh && isNum {
			ptr := &cache[i]
			if *ptr != nil {
				return **ptr
			}
			defer func() { *ptr = &res }()
		}

		if !isNum && low[i] == '0' {
			res += dfs(i+1, true, false, false)
		}

		lo, hi := 0, 9
		if limitLow {
			lo = int(low[i] - '0')
		}
		if limitHigh {
			hi = int(high[i] - '0')
		}

		d0 := 1
		if isNum {
			d0 = 0
		}

		for d := max(lo, d0); d <= hi; d++ {
			if !set[d] { // 本题约束：只能使用 digits 数组构建数字
				continue
			}

			v := dfs(i+1, limitLow && d == lo, limitHigh && d == hi, true)
			res += v
		}

		return res
	}

	var ans = dfs(0, true, true, false)
	return ans
}

func atMostNGivenDigitSet3(digits []string, n int) int {
	high := strconv.FormatInt(int64(n), 10)
	low := strconv.FormatInt(1, 10)

	var m = len(high)
	low = strings.Repeat("0", m-len(low)) + low // 前导 0 方便判断

	type Key struct {
		Index               int
		LimitLow, LimitHigh bool
		IsNum               bool
	}
	var cache = make(map[Key]int, m)

	var set = make([]bool, 10)
	for _, digit := range digits {
		set[digit[0]-'0'] = true
	}

	// v2拓展版本：支持前导零
	var dfs func(i int, limitLow, limitHigh bool, isNum bool) (res int)
	dfs = func(i int, limitLow, limitHigh bool, isNum bool) (res int) {
		if i == m {
			if isNum {
				return 1
			}
			return 0
		}

		key := Key{
			Index:     i,
			LimitLow:  limitLow,
			LimitHigh: limitHigh,
			IsNum:     isNum,
		}
		if v, ok := cache[key]; ok {
			return v
		}
		defer func() { cache[key] = res }()

		if !isNum && low[i] == '0' {
			res += dfs(i+1, true, false, false)
		}

		lo, hi := 0, 9
		if limitLow {
			lo = int(low[i] - '0')
		}
		if limitHigh {
			hi = int(high[i] - '0')
		}

		d0 := 1
		if isNum {
			d0 = 0
		}

		for d := max(lo, d0); d <= hi; d++ {
			if !set[d] { // 本题约束：只能使用 digits 数组构建数字
				continue
			}

			v := dfs(i+1, limitLow && d == lo, limitHigh && d == hi, true)
			res += v
		}

		return res
	}

	var ans = dfs(0, true, true, false)
	return ans
}
