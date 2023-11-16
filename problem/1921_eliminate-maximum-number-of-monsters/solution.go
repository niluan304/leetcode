package main

import (
	"sort"
)

// 先计算出怪物的到达时间, 然后再排序
// 时间复杂度：O(n) / 排序的 O(nlogn)
// 空间复杂度：O(n)
func eliminateMaximum(dist []int, speed []int) int {
	var n = len(dist)
	var cost = make([]int, 0, n)
	for i := range dist {
		// 题目要求整除结果向上取整, 类似 math.Ceil()
		cost = append(cost, (dist[i]-1)/speed[i]+1)
	}
	sort.Ints(cost)

	// debug
	//fmt.Println(dist, "dist")
	//fmt.Println(speed, "speed")
	//fmt.Println(cost, "cost")
	//fmt.Println()

	for i, c := range cost {
		if c <= i {
			return i
		}
	}

	return n
}
