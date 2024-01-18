package main

import (
	"math"
	"strconv"
	"strings"
)

func numberOfBeautifulIntegers(l int, h int, k int) int {
	high := strconv.FormatInt(int64(h), 10)
	low := strconv.FormatInt(int64(l), 10)

	var m = len(high)
	low = strings.Repeat("0", m-len(low)) + low // 前导 0 方便判断

	type Key struct {
		Index               int
		Diff, Mod           int
		LimitLow, LimitHigh bool
		IsNum               bool
	}
	var cache = make(map[Key]int, m)

	// v2拓展版本：支持前导零
	var dfs func(i int, diff, mod int, limitLow, limitHigh bool, isNum bool) (res int)
	dfs = func(i int, diff, mod int, limitLow, limitHigh bool, isNum bool) (res int) {
		if i == m {
			if isNum && diff == 0 && mod == 0 {
				return 1
			}
			return 0
		}

		// 摆烂型 cache 查询。
		// 优化做法：(!limitLow && !limitHigh && isNum) == true 时才会重复递归，
		// 那么可以在条件成立时，才查询缓存， 因此 Key 可以不存储这三个参数
		key := Key{
			Index:     i,
			Diff:      diff, // diff 取值范围 [-m, +m]，如果使用数组，可以将 diff 向右偏移 m，并初始化为 m，结束时 判断 diff == m 。
			Mod:       mod,
			LimitLow:  limitLow,
			LimitHigh: limitHigh,
			IsNum:     isNum,
		}
		if v, ok := cache[key]; ok {
			return v
		}
		defer func() { cache[key] = res }()

		if !isNum && low[i] == '0' {
			res += dfs(i+1, diff, mod, true, false, false) // 继续填 0
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
			d0 = 0 // 如果前面没有填数字，必须从 1 开始（因为不能有前导零）
		}

		for d := max(lo, d0); d <= hi; d++ {
			diff := diff + (d%2*2 - 1)                      // 偶数 -1，奇数 +1
			mod := (mod%k + d*int(math.Pow10(m-i-1))%k) % k // 模运算的加法：(a+b) % k = (a%k + b%k)%k，这里的 b 是 d*10^(m-i-1)

			res += dfs(i+1, diff, mod, limitLow && d == lo, limitHigh && d == hi, true)
		}

		return res
	}

	var ans = dfs(0, 0, 0, true, true, false)
	return ans
}
