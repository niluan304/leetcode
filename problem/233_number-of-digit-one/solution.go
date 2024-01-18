package main

import (
	"math/bits"
	"strconv"
	"strings"
)

func countDigitOne(n int) int {
	high := strconv.FormatInt(int64(n), 10)
	low := strconv.FormatInt(0, 10)

	var m = len(high)
	low = strings.Repeat("0", m-len(low)) + low // 前导 0 方便判断

	type Key struct {
		Index               int
		Value               int
		LimitLow, LimitHigh bool
	}
	var cache = make(map[Key]int, m)

	var ans = 0
	// v2基础版本：不支持前导零
	var dfs func(i int, value int, limitLow, limitHigh bool) (res int)
	dfs = func(i int, value int, limitLow, limitHigh bool) (res int) {
		if i == m {
			return value
		}

		key := Key{
			Index:     i,
			Value:     value,
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
			val := value
			if d == 1 {
				val++
			}

			v := dfs(i+1, val, limitLow && d == lo, limitHigh && d == hi)
			res += v
		}

		return res
	}

	ans = dfs(0, 0, true, true)
	return ans
}

func countDigitOne2(n int) int {
	high := strconv.FormatInt(int64(n), 10)
	low := strconv.FormatInt(0, 10)

	var m = len(high)
	low = strings.Repeat("0", m-len(low)) + low // 前导 0 方便判断

	var cache = make([][1 << 10]*int, m)

	var ans = 0
	// v2基础版本：不支持前导零
	var dfs func(i int, mask uint, limitLow, limitHigh bool) (res int)
	dfs = func(i int, mask uint, limitLow, limitHigh bool) (res int) {
		if i == m {
			return bits.OnesCount(mask)
		}

		if !limitLow && !limitHigh {
			ptr := &cache[i][mask]
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
			mask := mask
			if d == 1 {
				mask |= 1 << (i + 1)
			}

			v := dfs(i+1, mask, limitLow && d == lo, limitHigh && d == hi)
			res += v
		}

		return res
	}

	ans = dfs(0, 0, true, true)
	return ans
}
