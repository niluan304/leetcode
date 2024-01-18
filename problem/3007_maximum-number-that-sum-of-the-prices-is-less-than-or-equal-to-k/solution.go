package main

import (
	"sort"
	"strconv"
	"strings"
)

// 二分答案 + 数位dp
func findMaximumNumber(k int64, x int) int64 {
	n := int(1e15) + 1 // 不知道上界时的做法：设一个足够大的数
	n = int(k+1) << x  // 数学推导上界

	ans := sort.Search(n, func(num int) bool {
		high := strconv.FormatInt(int64(num), 2)
		low := strconv.FormatInt(0, 10)
		res := digitDp(low, high, x)
		return int64(res) >= k+1
	})

	return int64(ans - 1)
}

func digitDp(low, high string, x int) int {
	var m = len(high)
	low = strings.Repeat("0", m-len(low)) + low // 前导 0 方便判断
	var cache = make([][]*int, m)
	for i, _ := range cache {
		cache[i] = make([]*int, m) // 最大价值为: 1 * m
	}

	var dfs func(i int, value int, limitLow, limitHigh bool) (res int)
	dfs = func(i int, value int, limitLow, limitHigh bool) (res int) {
		if i == m {
			return value
		}

		if !limitLow && !limitHigh {
			ptr := &cache[i][value]
			if *ptr != nil {
				return **ptr
			}
			defer func() { *ptr = &res }()
		}

		lo, hi := 0, 9
		lo, hi = 0, 1 // 本题的约束：只构造二进制数，因此初始化 lo, hi := 0, 1
		if limitLow {
			lo = int(low[i] - '0')
		}
		if limitHigh {
			hi = int(high[i] - '0')
		}

		for d := lo; d <= hi; d++ {
			value := value
			if (m-i)%x == 0 && d == 1 { // 价值 +1
				value++
			}
			res += dfs(i+1, value, limitLow && d == lo, limitHigh && d == hi)
		}
		return res
	}

	ans := dfs(0, 0, true, true)

	return ans
}
