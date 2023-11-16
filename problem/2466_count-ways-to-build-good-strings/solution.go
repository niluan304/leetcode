package main

// dfs + 记忆化搜索
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func countGoodStrings(low int, high int, zero int, one int) int {
	const Mod = 1_000_000_007

	// cache[i] 可能为 0, 使用数组作为cache, 需要额外处理, 否则可能死循环
	var cache = make(map[int]int, high+1)
	cache[0] = 1 // 处理空字符串
	var dfs func(n int) int
	dfs = func(n int) int {
		if n < 0 {
			return 0
		}

		v, ok := cache[n]
		if ok {
			return v
		}

		v = (dfs(n-zero) + dfs(n-one)) % Mod
		cache[n] = v
		return v
	}

	//// 只写 dfs(high) 的话，可能 dp[high-1] 等等这样的状态是算不出来的。
	//// 直接在最后的 for 循环把 dfs(i) 累加起来就行。
	//// 作者：灵茶山艾府
	//// 链接：https://leetcode.cn/problems/count-ways-to-build-good-strings/solutions/1964910/by-endlesscheng-4j22/comments/2033765
	//dfs(high)

	var ans = 0
	for i := low; i <= high; i++ {
		ans = (ans + dfs(i)) % Mod
	}
	return ans
}

// dp 动态规划
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func countGoodStrings2(low int, high int, zero int, one int) int {
	const Mod = 10*10000*10000 + 7

	var dp = make([]int, high+1)
	dp[0] = 1 // 处理空字符串

	for i := 1; i <= high; i++ {
		i0 := i - zero
		i1 := i - one

		var x = 0
		if i0 >= 0 {
			x += dp[i0]
		}
		if i1 >= 0 {
			x += dp[i1]
		}
		dp[i] = x % Mod
	}

	var ans = 0
	for i := low; i <= high; i++ {
		ans += dp[i]
		ans %= Mod
	}
	return ans
}

// dfs + 记忆化搜索
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func countGoodStrings3(low int, high int, zero int, one int) int {
	const Mod = 1_000_000_007
	const Empty = -1

	// cache[i] 可能为 0, 使用数组作为cache, 需要额外处理, 否则可能死循环
	var cache = make([]int, high+1)
	cache[0] = 1 // 处理空字符串
	for i := 1; i <= high; i++ {
		cache[i] = Empty
	}
	var dfs func(n int) int
	dfs = func(n int) int {
		if n < 0 {
			return 0
		}

		v := cache[n]
		if v != Empty {
			return v
		}

		v = (dfs(n-zero) + dfs(n-one)) % Mod
		cache[n] = v
		return v
	}

	//// 只写 dfs(high) 的话，可能 dp[high-1] 等等这样的状态是算不出来的。
	//// 直接在最后的 for 循环把 dfs(i) 累加起来就行。
	//// 作者：灵茶山艾府
	//// 链接：https://leetcode.cn/problems/count-ways-to-build-good-strings/solutions/1964910/by-endlesscheng-4j22/comments/2033765
	//dfs(high)

	var ans = 0
	for i := low; i <= high; i++ {
		ans = (ans + dfs(i)) % Mod
	}
	return ans
}
