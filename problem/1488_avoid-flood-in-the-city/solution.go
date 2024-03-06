package main

import (
	"sort"

	. "github.com/niluan304/leetcode/container"
)

// 贪心 + 二分搜索
// 时间复杂度：O(nlogn)
// 空间复杂度：O(n)
func avoidFlood(rains []int) []int {
	n := len(rains)
	ans := make([]int, n)

	empty, rainDays := make([]int, 0), make(map[int]int)

	for i, rain := range rains {
		if rain == 0 {
			empty = append(empty, i)
			ans[i] = 1
			continue
		}
		if day, ok := rainDays[rain]; ok {
			j := sort.Search(len(empty), func(i int) bool {
				return empty[i] > day
			})
			if j == len(empty) {
				return nil
			}
			ans[empty[j]] = rain
			empty = append(empty[:j], empty[j+1:]...)
		}
		rainDays[rain] = i
		ans[i] = -1
	}

	return ans
}

// 贪心 + 优先队列
// 时间复杂度：O(nlogn)
// 空间复杂度：O(n)
func avoidFlood2(rains []int) []int {
	type Task struct {
		Day  int
		Rain int
	}

	rainDays := make(map[int]int)
	n := len(rains)
	ans := make([]int, n)
	tasks := make([]*Task, n)

	for i, rain := range rains {
		if rain == 0 {
			continue
		}

		if j, ok := rainDays[rain]; ok {
			tasks[j] = &Task{
				Day:  i,
				Rain: rain,
			}
		}
		rainDays[rain] = i
	}

	hp := NewEmptyHeap(func(x, y Task) bool {
		return x.Day < y.Day
	})
	for i, rain := range rains {
		if tasks[i] != nil {
			hp.Insert(*tasks[i])
		}
		if rain != 0 {
			ans[i] = -1
			continue
		}
		if hp.Len() == 0 {
			ans[i] = 1
			continue
		}
		task := hp.PopHead()
		if task.Day < i {
			return nil
		}
		ans[i] = task.Rain
	}

	if hp.Len() > 0 {
		return nil
	}
	return ans
}
