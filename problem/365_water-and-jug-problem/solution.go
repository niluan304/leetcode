package main

func canMeasureWater(jug1 int, jug2 int, target int) bool {
	if target > jug1+jug2 {
		return false
	}

	type Key struct{ x, y int }

	var visit = make(map[Key]bool, jug1+jug2)
	var dfs func(x, y int) bool
	dfs = func(x, y int) (res bool) {
		sum := x + y
		if x == target || y == target || sum == target {
			return true
		}

		key := Key{x: x, y: y}
		if visit[key] {
			return false
		}
		visit[key] = true

		return dfs(0, y) || // 清空水壶 x
			dfs(x, 0) || // 清空水壶 y
			dfs(jug1, y) || // 装满水壶 x
			dfs(x, jug2) || // 装满水壶 y
			dfs(min(jug1, sum), max(sum-jug1, 0)) || // 从水壶 y 向 水壶 x 倒水
			dfs(max(0, sum-jug2), min(sum, jug2)) //从水壶 x 向 水壶 y 倒水
	}

	ans := dfs(0, 0)
	return ans
}
