package main

import (
	"math/bits"
	"strings"
)

// 状态压缩 dp
// 由于状态比较离散，数组效率不如哈希表，使用 数组记录缓存反而会超时
//
// - 时间复杂度：$\mathcal{O}(m \cdot n \cdot 2^n \cdot 2^n)$。
// - 空间复杂度：$\mathcal{O}(n)$。
func maxStudents(seats [][]byte) int {

	var m, n = len(seats), len(seats[0])
	var ans = 0 // math.MaxInt32 // math.MinInt32

	type Key struct {
		i, j     int
		pre, cur int
	}
	var cache = make(map[Key]int, 1<<16)

	var dfs func(i, j, pre, cur int) int
	dfs = func(i, j, pre, cur int) (res int) {
		if j >= n {
			i, j = i+1, 0
			pre, cur = cur, 0
		}
		if i == m {
			return 0
		}

		key := Key{i: i, j: j, pre: pre, cur: cur}
		if v, ok := cache[key]; ok {
			return v
		}
		defer func() { cache[key] = res }()

		// 使用该座位
		k := n - j
		if seats[i][j] == '.' && // 该座位完好
			(cur>>(k+1)&1) == 0 && // 左边无人
			pre>>(k+1)&1 == 0 && // 左上无人
			pre>>(k-1)&1 == 0 { // 右上无人

			res = dfs(i, j+2, pre, cur|1<<k) + 1
		}

		// 不使用该座位
		res = max(res, dfs(i, j+1, pre, cur))
		return res
	}

	ans = dfs(0, 0, 0, 0)
	return ans
}

func maxStudents2(seats [][]byte) int {

	var m, n = len(seats), len(seats[0])
	var ans = 0 // math.MaxInt32 // math.MinInt32

	type Key struct {
		i               int
		prefix, current string
	}
	var cache = make(map[Key]int, 1<<n)

	var dfs func(i int, current, prefix string) int
	dfs = func(i int, current, prefix string) (res int) {
		j := len(current) - 1

		if j >= n {
			i, j = i+1, 0
			prefix, current = current+"1", "1"
		}
		if i == m {
			return 0
		}

		key := Key{i: i, prefix: prefix, current: current}
		if v, ok := cache[key]; ok {
			return v
		}
		defer func() { cache[key] = res }()

		// 不使用该座位
		res = dfs(i, current+"1", prefix)
		// 用 1 表示座位无人， 0 表示座位有人
		// 使用该座位
		//k := n - j
		if seats[i][j] == '.' && // 该座位完好
			current[j] == '1' && // 左边无人
			prefix[j] == '1' && // 左上无人
			prefix[j+2] == '1' { // 右上无人

			res = max(res, dfs(i, current+"0", prefix)+1)
		}
		return res
	}
	// 用 1 表示座位无人， 0 表示座位有人
	ans = dfs(0, "1", strings.Repeat("1", n+2))
	return ans
}

func maxStudents3(seats [][]byte) int {
	var m, n = len(seats), len(seats[0])
	var ans = 0

	type Key struct{ i, cur, pre int }
	var cache = make(map[Key]int, 1<<n)

	// 用 1 表示座位无人， 0 表示座位有人
	var dfs func(i int, cur, pre int) int
	dfs = func(i int, cur, pre int) (res int) {
		// cur 表示第 i 排 已操作过的座位，接下来要安排 seats[i][j] 能否坐人，
		// 为方便计算 j, 可以令 cur 中 1 为无人的，0 为有人的，且最左边还有一个无人的位置，
		// 那么最高位的 1 的位置，就是 j。
		j := bits.Len(uint(cur)) - 1

		if j >= n {
			i, j = i+1, 0
			pre, cur = (cur<<1)|1, 1
		}
		if i == m {
			return 0
		}

		key := Key{i: i, cur: cur, pre: pre}
		if v, ok := cache[key]; ok {
			return v
		}
		defer func() { cache[key] = res }()

		// 不使用该座位
		res = dfs(i, cur<<1|1, pre)

		// 使用该座位
		k := n - j
		if seats[i][j] == '.' && // 该座位完好，可以坐人
			cur&1 == 1 && // 左边无人
			(pre>>(k+1))&1 == 1 && // 左上无人
			(pre>>(k-1))&1 == 1 { // 右上无人

			res = max(res, dfs(i, cur<<1|0, pre)+1)
		}

		return res
	}
	ans = dfs(0, 1, 1<<(n+2)-1)
	return ans
}
