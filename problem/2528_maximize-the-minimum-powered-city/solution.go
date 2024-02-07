package main

import (
	"sort"
)

// 综合性题目：
// 1. 前缀和计算城市的实际电量
// 2. 二分搜索寻找最小的 minPower
// 3. 差分数组更新供电站的数组
//
// 二分答案 minPower，从左到右遍历 stations，如果 stations[i] 电量小于 minPower，那么需要建供电站来补足。
func maxPower(stations []int, r int, k int) int64 {
	var n = len(stations)
	var sum = make([]int, n+1) // 前缀和
	for i, x := range stations {
		sum[i+1] = sum[i] + x
	}

	var powers, mm = make([]int, n+1), 0
	for i := range stations {
		powers[i+1] = sum[min(i+r+1, n)] - sum[max(i-r, 0)] // 电量
		mm = max(mm, powers[i+1])
	}

	// 二分搜索：在 nums 中，找到第一个 i，使得 nums[i] <= x
	var ans = sort.Search(mm+k, func(minPower int) bool {
		var diff = make([]int, n+1)
		var need, power = 0, 0
		for i := 0; i < n; i++ {
			// 获取 powers[i] 的写法 1：
			diff[i] += powers[i+1] - powers[i] // 同时构建和更新差分数组
			power += diff[i]                   // 差分数组 [0:i] 的前缀和即为还原后的 powers[i]

			//// 获取 powers[i] 的写法 2：
			//// 由于 diff[i] 不再用到，可以等价写为：
			//power += powers[i+1] - powers[i] + diff[i] // 还原后的 powers[i]

			x := (minPower + 1) - power // +1，改为二分最小的不满足要求的值，这样 sort.Search 返回的就是最大的满足要求的值
			// 需要 x 个供电站
			if x <= 0 {
				continue
			}

			need += x

			// 差分更新
			power += x // powers[i] 也应当更新
			//diff[i] += x // diff[i] 不再使用，可不更新
			if i+r*2+1 < n {
				diff[i+r*2+1] -= x
			}
		}
		return need > k
	})

	return int64(ans)
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// 1. 使用差分数组计算和更新城市的电量
// 2. 二分答案：假设最小电量为 minPower，贪心的计算需要新建多少个电站，然后再和预算比较
func maxPower2(stations []int, r int, k int) int64 {
	var n = len(stations)
	var origin, diff = make([]int, n), make([]int, n)

	for i, station := range stations {
		left := max(0, i-r)
		right := i + r // [i-r, i+r] 区间上加上 发电站数量station

		origin[left] += station
		if right+1 < n {
			origin[right+1] -= station
		}
	}

	// 因为边界取值，喜提 WA*2
	ans := sort.Search(1e5*1e5+1e9+1, func(minPower int) bool {
		var cnt = 0
		copy(diff, origin)

		for i := 0; i < n; i++ {
			if i > 0 {
				diff[i] += diff[i-1] // 还原差分数组，得到 i 有多少发电站数量
			}

			x := minPower - diff[i]
			if x > 0 {
				diff[i] += x
				right := i + 2*r
				if right+1 < n {
					diff[right+1] -= x
				}
				cnt += x
			}
		}
		return cnt > k
	})
	return int64(ans - 1)
}
