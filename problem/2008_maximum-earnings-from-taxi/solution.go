package main

// ### 思路：动态规划
// 1. 从左往右思考：有乘客信息 [start,end,tip]，如果知道从 end 开到 n 最多能赚多少钱，那么就知道完成 **当前接单** 最多能赚多少钱，
// start 处可能有多位乘客，所以还需要在 同一 start 的订单中取一个极值，那么就能得到 **从 start 开到 n ** 最多能赚多少钱。
// 2. 还有一种情况：如果 [start+1,end-1] 处有利润更高的订单，那么就应该不接 start 处的订单。
//
// ### 复杂度分析
//
// - 时间复杂度：$\mathcal{O}(n+m)$，其中 $m$ 为 $\textit{rides}$ 的长度。内外层循环次数之和是 $\mathcal{O}(n+m)$，所以时间复杂度是 $\mathcal{O}(n+m)$。
// - 空间复杂度：$\mathcal{O}(n+m)$。
func maxTaxiEarnings(n int, rides [][]int) int64 {
	type Item struct{ End, Earn int }
	var items = make([][]Item, n)
	for _, ride := range rides {
		start, end, tip := ride[0], ride[1], ride[2]
		items[start] = append(items[start], Item{
			End: end,
			// 提前计算利润，也可以只记录 tip，dp中的内循环写为：
			// earn = max(earn, dp[item.End]+item.End-i+item.Tip)
			Earn: end - start + tip,
		})
	}
	var dp = make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		var earn = 0
		for _, item := range items[i] {
			earn = max(earn, dp[item.End]+item.Earn)
		}
		dp[i] = max(dp[i+1], earn)
	}

	return int64(dp[0])
}

// ## 一、启发思考：寻找子问题
//
// 假设 $n=9$。
//
// 我们要解决的问题是从 $1$ 开到 $9$ 最多可以赚多少钱。
//
// 如果没有乘客在 $9$ 下车，或者我们不载在 $9$ 下车的乘客，那么问题变成：
//
// - 从 $1$ 开到 $8$ 最多可以赚多少钱。
//
// 如果有至少一位乘客在 $9$ 下车，我们可以枚举载哪位乘客。假设所载乘客在 $5$ 上车，那么从 $5$ 到 $9$ 不能载其它乘客（题目要求同时最多只能接一个订单），问题变成：
//
// - 从 $1$ 开到 $5$ 最多可以赚多少钱。
//
// 注意可以在一个地点放下一位乘客，并在同一个地点接上另一位乘客。
//
// 不载乘客和枚举载哪个乘客，都会把原问题变成一个**和原问题相似的、规模更小的子问题**，这意味着我们可以用**递归**解决。
//
// > 注：从右往左思考，主要是为了方便把递归翻译成递推。从左往右思考也是可以的。
//
// ### 复杂度分析
//
// - 时间复杂度：$\mathcal{O}(n+m)$，其中 $m$ 为 $\textit{rides}$ 的长度。内外层循环次数之和是 $\mathcal{O}(n+m)$，所以时间复杂度是 $\mathcal{O}(n+m)$。
// - 空间复杂度：$\mathcal{O}(n+m)$。
//
// 作者：灵茶山艾府
// 链接：https://leetcode.cn/problems/maximum-earnings-from-taxi/solutions/2558504/jiao-ni-yi-bu-bu-si-kao-dong-tai-gui-hua-k15a/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func maxTaxiEarnings2(n int, rides [][]int) int64 {
	type Item struct{ Start, Earn int }
	var groups = make([][]Item, n+1)
	for _, ride := range rides {
		start, end, tip := ride[0], ride[1], ride[2]
		groups[end] = append(groups[end], Item{
			Start: start,
			Earn:  end - start + tip,
		})
	}

	var dp = make([]int, n+1)
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1]
		for _, item := range groups[i] {
			dp[i] = max(dp[i], dp[item.Start]+item.Earn)
		}
	}
	return int64(dp[n])
}
