package main

import (
	"math"
	"sort"
)

// timeout
// 时间复杂度：O(n^2)
// 空间复杂度：O(n)
func repairCars(ranks []int, cars int) int64 {
	type Cost struct {
		Car  int64
		Cost int64
	}

	var costs = make([]Cost, len(ranks))
	for i := range costs {
		rank := int64(ranks[i])
		costs[i] = Cost{
			Cost: rank - 100,
		}
	}
	for i := 0; i < cars; i++ {
		var idx = 0
		var cur int64 = math.MaxInt
		for j, cost := range costs {
			next := int64(ranks[j]) * (cost.Car + 1) * (cost.Car + 1)
			if cur > next {
				cur = next
				idx = j
			}
		}
		costs[idx].Car++
		costs[idx].Cost = cur
	}

	var ans int64 = 0
	for _, cost := range costs {
		ans = max(ans, cost.Cost)
	}

	return ans
}

// 二分查找
// 时间复杂度：O(nlog(n))
// 空间复杂度：O(1)
func repairCars2(ranks []int, cars int) int64 {
	// tip1
	// For a predefined fixed time, can all the cars be repaired?
	var enough = func(cost int) bool {
		var count = 0
		for _, rank := range ranks {
			count += int(math.Sqrt(float64(cost) / float64(rank)))
		}
		return count >= cars
	}

	// tip2
	// Try using binary search on the answer.
	maxCost := ranks[0] * cars * cars
	return int64(sort.Search(maxCost, func(i int) bool {
		return enough(i)
	}))
}
