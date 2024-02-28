package main

// 思考题：如果操作改为可以加一或减一呢？
func discuss(n int, cost []int) int {
	ans := 0
	var dfs func(i int) (low, high int)
	dfs = func(i int) (low, high int) {
		if i > n {
			return 0, 0
		}

		leftLow, lowHigh := dfs(2 * i)      // 左儿子的取值范围
		rightLow, rightHigh := dfs(2*i + 1) // 右儿子的取值范围

		z0, z1, c := IntervalIntersect(leftLow, lowHigh, rightLow, rightHigh) // 让左右儿子平衡的取值范围，以及操作代价
		ans += c
		return cost[i-1] + z0, cost[i-1] + z1 // 当前节点的取值范围
	}
	dfs(1)
	return ans
}

func IntervalIntersect(x0, x1 int, y0, y1 int) (z0, z1 int, cost int) {
	// 不相交的两种情况
	if y0 > x1 {
		return x1, y0, y0 - x1
	}
	if x0 > y1 {
		return y1, x0, x0 - y1
	}

	// 相交的情况
	return max(x0, y0), min(x1, y1), 0
}
