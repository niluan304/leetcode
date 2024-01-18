package main

import (
	"math/big"
	"strings"
)

const MOD = 1_000_000_007

func count(num1 string, num2 string, minSum int, maxSum int) int {
	x := new(big.Int)
	x.SetString(num1, 10)
	num1 = x.Sub(x, big.NewInt(1)).String()

	return (digitDp(num2, minSum, maxSum) - digitDp(num1, minSum, maxSum) + MOD) % MOD
}

func digitDp(s string, minSum, maxSum int) int {

	m := len(s)
	cache := make([][]*int, m)
	for i, _ := range cache {
		cache[i] = make([]*int, maxSum+1)
	}

	var dfs func(i, sum int, isLimit, isNum bool) (res int)
	dfs = func(i, sum int, isLimit, isNum bool) (res int) {
		if i == m {
			if isNum && minSum <= sum && sum <= maxSum {
				return 1
			}
			return 0
		}

		// 问：记忆化四个状态有点麻烦，能不能只记忆化 (i,mask) 这两个状态？
		//
		// 答：是可以的。比如 n=234，第一位填 2，第二位填 3，后面无论怎么递归，都不会再次递归到第一位填 2，第二位填 3 的情况，所以不需要记录。
		// 又比如，第一位不填，第二位也不填，后面无论怎么递归也不会再次递归到这种情况，所以也不需要记录。
		if !isLimit && isNum {
			ptr := &cache[i][sum]
			if *ptr != nil {
				return **ptr
			}
			defer func() { *ptr = &res }()
		}

		if !isNum { // 前导 0 的情况
			v := dfs(i+1, sum, false, false)
			res = (v%MOD + res%MOD) % MOD
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
			x := sum + d
			if x > maxSum {
				break
			}
			v := dfs(i+1, x, isLimit && d == up, true)
			res = (v%MOD + res%MOD) % MOD
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

func count2(num1 string, num2 string, minSum int, maxSum int) int {
	high := num2
	low := num1

	var m = len(high)
	low = strings.Repeat("0", m-len(low)) + low // 前导 0 方便判断
	type Key struct {
		Index               int
		Sum                 int
		LimitLow, LimitHigh bool
	}
	var cache = make(map[Key]int, m)

	var ans = 0
	// v2基础版本：不支持前导零
	var dfs func(i int, sum int, limitLow, limitHigh bool) (res int)
	dfs = func(i int, sum int, limitLow, limitHigh bool) (res int) {
		if i == m {
			if minSum <= sum && sum <= maxSum {
				return 1
			}
			return 0
		}

		key := Key{
			Index:     i,
			Sum:       sum,
			LimitLow:  limitLow,
			LimitHigh: limitHigh,
		}

		if v, ok := cache[key]; ok {
			return v
		}
		defer func() { cache[key] = res }()

		lo, hi := 0, 9
		if limitLow {
			lo = int(low[i] - '0')
		}
		if limitHigh {
			hi = int(high[i] - '0')
		}

		for d := lo; d <= hi; d++ {
			v := dfs(i+1, sum+d, limitLow && d == lo, limitHigh && d == hi)
			res = (res%MOD + v%MOD) % MOD
		}

		return res
	}

	ans = dfs(0, 0, true, true)
	return ans
}

func count3(num1 string, num2 string, minSum int, maxSum int) int {
	high := num2
	low := num1

	var m = len(high)
	low = strings.Repeat("0", m-len(low)) + low // 补前导零，和 high 对齐

	var cache = make([][21*9 + 1]*int, m) // 最大和：9*21

	// v2基础版本：不支持前导零
	var dfs func(i int, sum int, limitLow, limitHigh bool) (res int)
	dfs = func(i int, sum int, limitLow, limitHigh bool) (res int) {
		if sum > maxSum { // 减枝
			return 0
		}

		if i == m {
			if minSum <= sum && sum <= maxSum {
				return 1
			}
			return 0
		}

		if !limitLow && !limitHigh {
			ptr := &cache[i][sum]
			if *ptr != nil {
				return **ptr
			}
			defer func() { *ptr = &res }()
		}

		lo, hi := 0, 9
		if limitLow {
			lo = int(low[i] - '0')
		}
		if limitHigh {
			hi = int(high[i] - '0')
		}

		for d := lo; d <= hi; d++ {
			v := dfs(i+1, sum+d, limitLow && d == lo, limitHigh && d == hi)
			res = (res + v) % MOD
		}

		return res
	}

	var ans = dfs(0, 0, true, true)
	return ans
}
