package main

import (
	"container/heap"
	"sort"
)

// 动态规划
// n代表课程数量，m代表课程可完成的最晚时间
// 时间复杂度：O(n*m)
// 空间复杂度：O(m)
// 作者：Carter
// 链接：https://leetcode.cn/problems/course-schedule-iii/solutions/1156995
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func scheduleCourse(courses [][]int) int {
	//贪心，先考虑早结束的课程，按课程的结束时间从早到晚排序
	sort.Slice(courses, func(i, j int) bool {
		return courses[i][1] < courses[j][1]
	})

	//课程的最晚结束时间
	var n = courses[len(courses)-1][1]
	//原来是二维的dp数组(课程编号,结束时间),使用状态压缩,只保留结束时间.值是可以修读的课程数
	var dp = make([]int, n+1)
	var ans = 0
	//遍历每一个课程
	for _, course := range courses {
		//对于当前遍历的课程,结束时间的取值范围为[课程的耗时,课程最晚的结束时间]
		for j := course[1]; j >= course[0]; j-- {
			//更新可以修读的课程数量,并更新答案
			dp[j] = max(dp[j], dp[j-course[0]]+1)
			ans = max(ans, dp[j])
		}
	}
	return ans
}

// 优先队列/堆排序
func scheduleCourse2(courses [][]int) int {
	sort.Slice(courses, func(i, j int) bool { return courses[i][1] < courses[j][1] })
	var h = new(hp)
	total := 0 // 已消耗时间
	for _, course := range courses {
		d := course[0]
		if total+d <= course[1] { // 没有超过 lastDay，直接学习
			total += d
			heap.Push(h, d)
		} else if h.Len() > 0 && d < h.IntSlice[0] { // 该课程的时间比之前的最长时间要短
			top := h.IntSlice[0]
			h.IntSlice[0] = d
			heap.Fix(h, 0)

			// 节省出来的时间，能在后面上完更多的课程
			// 反悔，撤销之前 duration 最长的课程，改为学习该课程
			total += d - top
		}
	}
	return h.Len()
}

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] } // 最大堆
func (h *hp) Push(v any)        { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() any          { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }

func scheduleCourse3(courses [][]int) int {
	//贪心，先考虑早结束的课程，按课程的结束时间从早到晚排序
	sort.Slice(courses, func(i, j int) bool {
		x, y := courses[i], courses[j]
		if x[1] == y[1] {
			return x[0] < y[0]
		}
		return x[1] < y[1]
	})
	var day = 0
	var h IntHeap
	for _, course := range courses {
		duration, last := course[0], course[1]
		if last >= day+duration {
			heap.Push(&h, duration)
			day += duration
		} else {
			if h.Len() > 0 && duration < h[0] {
				day += duration - h[0]
				h[0] = duration
				heap.Fix(&h, 0)
			}
		}
	}
	return h.Len()
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] } // 最大堆: h[i] > h[j]，最小堆：h[i] < h[j]
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func scheduleCourse4(courses [][]int) int {
	//贪心，先考虑早结束的课程，按课程的结束时间从早到晚排序
	sort.Slice(courses, func(i, j int) bool {
		x, y := courses[i], courses[j]
		if x[1] == y[1] {
			return x[0] < y[0]
		}
		return x[1] < y[1]
	})
	var day = 0
	var h IntHeap
	for _, course := range courses {
		duration, last := course[0], course[1]

		heap.Push(&h, duration)
		day += duration

		if last < day {
			day -= heap.Pop(&h).(int)
		}
	}

	return h.Len()
}
