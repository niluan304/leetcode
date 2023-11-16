package main

import (
	"math"
	"strconv"
)

func splitIntoFibonacci(num string) []int {
	var ans []int
	var path []int
	var n = len(num)
	var dfs func(i int)
	dfs = func(i int) {
		if i == n && len(path) >= 3 {
			ans = append([]int{}, path...)
			return
		}

		for j := i + 1; j <= n; j++ {
			x, _ := strconv.Atoi(num[i:j])
			if (x > math.MaxInt32) || // 拆分出的数必须符合 32 位有符号整数类型，即每个数必须在 [0,2^31−1] 的范围内
				(j > i+1 && num[i] == '0') || // 拆分出的数如果不是 0，则不能以 0 开头
				(len(ans) > 0) { // 找到答案
				break
			}

			l := len(path)
			if len(path) >= 2 {
				y := path[l-1] + path[l-2]
				if x < y {
					continue
				}
				if x > y { // 剪枝：如果列表中至少有 2 个数，并且拆分出的数已经大于最后 2 个数的和，就不需要继续尝试拆分了。
					break
				}
			}
			path = append(path, x)
			dfs(j)
			path = path[:l]
		}
	}

	dfs(0)
	return ans
}
