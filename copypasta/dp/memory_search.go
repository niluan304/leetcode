package dp

// MemorySearch
// 实现 Python 的 cache 装饰器，返回 cache 用于 debug
func MemorySearch[X comparable, Y any](dfs *func(X) Y) (_cache map[X]Y) {
	var cache = make(map[X]Y, 1<<10)
	f := *dfs
	*dfs = func(x X) (res Y) {
		if v, ok := cache[x]; ok {
			return v
		}
		defer func() { cache[x] = res }()

		return f(x)
	}
	return cache
}

type key2[X, Y any] struct {
	x X
	y Y
}

// MemorySearch2
// todo 整合 MemorySearch
//
// 实现 Python 的 cache 装饰器，返回 cache 用于 debug
func MemorySearch2[X, Y comparable, Z any](dfs *func(X, Y) Z) (_cache map[key2[X, Y]]Z) {
	var cache = make(map[key2[X, Y]]Z, 1<<10)
	f := *dfs
	*dfs = func(x X, y Y) (res Z) {
		key := key2[X, Y]{x: x, y: y}
		if v, ok := cache[key]; ok {
			return v
		}
		defer func() { cache[key] = res }()

		return f(x, y)
	}
	return cache
}

type key3[X, Y, Z any] struct {
	x X
	y Y
	z Z
}

// MemorySearch3
// todo 整合 MemorySearch
//
// 实现 Python 的 cache 装饰器，返回 cache 用于 debug
func MemorySearch3[X, Y, Z comparable, R any](dfs *func(X, Y, Z) R) (_cache map[key3[X, Y, Z]]R) {
	var cache = make(map[key3[X, Y, Z]]R, 1<<10)
	f := *dfs
	*dfs = func(x X, y Y, z Z) (res R) {
		key := key3[X, Y, Z]{x: x, y: y, z: z}
		if v, ok := cache[key]; ok {
			return v
		}
		defer func() { cache[key] = res }()

		return f(x, y, z)
	}
	return cache
}

// 子集型回溯模板2
// 答案的视角，选哪个数
// 回溯三问：
// 当前操作？
// 子问题？
// 下一个子问题？
//
// LC78 https://leetcode.cn/problems/subsets/
// LC306 https://leetcode.cn/problems/additive-number/
// LC131 https://leetcode.cn/problems/palindrome-partitioning
// LC784 https://leetcode.cn/problems/letter-case-permutation/
// LC2698 https://leetcode.cn/problems/find-the-punishment-number-of-an-integer
// LC2397 https://leetcode.cn/problems/maximum-rows-covered-by-columns/
func backtrack(nums []int) [][]int {
	n := len(nums)
	ans := make([][]int, 0, 1<<n)
	path := make([]int, 0, n)
	var dfs func(int)
	dfs = func(i int) {
		ans = append(ans, append([]int(nil), path...)) // 固定答案
		if i == n {
			return
		}
		for j := i; j < n; j++ { // 枚举选择的数字
			path = append(path, nums[j])
			dfs(j + 1)
			path = path[:len(path)-1] // 恢复现场
		}
	}
	dfs(0)
	return ans
}
