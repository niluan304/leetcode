package main

import (
	"strconv"
)

var sn [1001]int

func init() {
	for i := 1; i <= 1000; i++ {
		sn[i] = sn[i-1] + isPunishment(i)
	}
}

// 回溯
// 时间复杂度：O(n^2)
// 空间复杂度：O(n)
func isPunishment(x int) (ans int) {
	y := x * x
	s := strconv.Itoa(y)
	n := len(s)

	var dfs func(i int, sum int)
	dfs = func(i int, sum int) {
		if i == n && sum == x {
			ans = y
			return
		}

		for j := i + 1; j <= n; j++ {
			v, _ := strconv.Atoi(s[i:j])
			if sum+v > x || ans != 0 {
				break
			}
			dfs(j, sum+v)
		}
	}
	dfs(0, 0)
	return
}

func punishmentNumber(n int) int {
	return sn[n]
}

var cache = [][2]int{
	{1, 1},
	{9, 81},
	{10, 100},
	{36, 1296},
	{45, 2025},
	{55, 3025},
	{82, 6724},
	{91, 8281},
	{99, 9801},
	{100, 10000},
	{235, 55225},
	{297, 88209},
	{369, 136161},
	{370, 136900},
	{379, 143641},
	{414, 171396},
	{657, 431649},
	{675, 455625},
	{703, 494209},
	{756, 571536},
	{792, 627264},
	{909, 826281},
	{918, 842724},
	{945, 893025},
	{964, 929296},
	{990, 980100},
	{991, 982081},
	{999, 998001},
	{1000, 1000000},
}

// 面向 oj 编程
func punishmentNumber2(n int) int {
	var ans = 0
	for _, v := range cache {
		if v[0] <= n {
			ans += v[1]
		}
	}
	return ans
}
