package main

import (
	"slices"
	"sort"
)

// 二分答案
// 关键词：最大化最小值
// 挨个判断每台机器最多可以制造多少份合金。
// 假设要制造 $num$ 份合金，由于 $num$ 越小，花费的钱越少， $num$ 越多，花费的钱越多，有单调性，可以二分。
//
// - 时间复杂度：$\mathcal{O}(k \cdot n \log U)$。其中 $U = \min{s} + budget$。
// - 空间复杂度：$\mathcal{O}(1)$。
func maxNumberOfAlloys(_ int, _ int, budget int, composition [][]int, stock []int, cost []int) int {
	var ans = 0 // math.MaxInt32 // math.MinInt32

	// 二分的右边界查找
	var mx int = 2e8 + 1 // 寻找超支预算的最小值，2e8 是不会超支的
	mx = slices.Min(stock) + budget + 1

	for _, com := range composition {
		// 二分的右边界应当为 (stock + budget) / min(composition)
		num := sort.Search(mx, func(num int) bool {
			money := 0
			for i, v := range com {
				d := num*v - stock[i]
				if d > 0 { // 超出库存了，需要使用预算
					money += d * cost[i]
				}
			}
			return money >= budget+1 // 生成 num 块合金的预算是超支预算的最小值，那生产 num-1 块合金一定不会超支预算
		}) - 1

		ans = max(ans, num)
	}

	return ans
}
