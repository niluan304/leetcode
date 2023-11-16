package main

import (
	"container/heap"
	"sort"
)

// 贪心 + 二分搜索
// 时间复杂度：O(nlogn)
// 空间复杂度：O(n)
func avoidFlood(rains []int) []int {
	var n = len(rains)
	var ans = make([]int, n)

	var empty, rainDays = make([]int, 0), make(map[int]int)

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

func _max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func _min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func _abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// 贪心 + 优先队列
// 时间复杂度：O(nlogn)
// 空间复杂度：O(n)
func avoidFlood2(rains []int) []int {
	var rainDays = make(map[int]int)
	var n = len(rains)
	var ans = make([]int, n)
	var tasks = make([]*Task, n)

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

	var pq = new(PriorityQueue)
	for i, rain := range rains {
		if tasks[i] != nil {
			heap.Push(pq, tasks[i])
		}
		if rain != 0 {
			ans[i] = -1
			continue
		}
		if pq.Len() == 0 {
			ans[i] = 1
			continue
		}
		task := heap.Pop(pq).(*Task)
		if task.Day < i {
			return nil
		}
		ans[i] = task.Rain
	}

	if pq.Len() > 0 {
		return nil
	}
	return ans
}

type Task struct {
	Day  int
	Rain int
}

type PriorityQueue []*Task

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].Day < pq[j].Day }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x any) { *pq = append(*pq, x.(*Task)) }

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}
