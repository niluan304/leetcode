package main

import (
	"slices"

	. "github.com/niluan304/leetcode/container"
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
	// 贪心，先考虑早结束的课程，按课程的结束时间从早到晚排序
	slices.SortFunc(courses, func(a, b []int) int {
		return a[1] - b[1]
	})

	// 课程的最晚结束时间
	n := courses[len(courses)-1][1]
	// 原来是二维的dp数组(课程编号,结束时间),使用状态压缩,只保留结束时间.值是可以修读的课程数
	dp := make([]int, n+1)
	ans := 0
	// 遍历每一个课程
	for _, course := range courses {
		// 对于当前遍历的课程,结束时间的取值范围为[课程的耗时,课程最晚的结束时间]
		for j := course[1]; j >= course[0]; j-- {
			// 更新可以修读的课程数量,并更新答案
			dp[j] = max(dp[j], dp[j-course[0]]+1)
			ans = max(ans, dp[j])
		}
	}
	return ans
}

// 优先队列/堆排序
func scheduleCourse2(courses [][]int) int {
	slices.SortFunc(courses, func(a, b []int) int {
		return a[1] - b[1]
	})

	h := NewMaxIntHeap()
	total := 0 // 已消耗时间
	for _, course := range courses {
		d := course[0]
		if total+d <= course[1] { // 没有超过 lastDay，直接学习
			total += d
			h.Insert(d)
		} else if h.Len() > 0 {
			head := h.Head() // h.Len() == 0 时，执行这个操作会发生数组越界
			if d < head {    // 该课程的时间比之前的最长时间要短
				h.Update(0, d)

				// 节省出来的时间，能在后面上完更多的课程
				// 反悔，撤销之前 duration 最长的课程，改为学习该课程
				total += d - head
			}
		}
	}
	return h.Len()
}

func scheduleCourse3(courses [][]int) int {
	// 贪心，先考虑早结束的课程，按课程的结束时间从早到晚排序
	slices.SortFunc(courses, func(x, y []int) int {
		if x[1] == y[1] {
			return x[0] - y[0]
		}
		return x[1] - y[1]
	})

	day := 0
	h := NewMaxIntHeap()
	for _, course := range courses {
		duration, last := course[0], course[1]
		if last >= day+duration {
			h.Insert(duration)
			day += duration
		} else if h.Len() > 0 {
			head := h.Head() // h.Len() == 0 时，执行这个操作会发生数组越界
			if duration < head {
				day += duration - head
				h.Update(0, duration)
			}
		}
	}
	return h.Len()
}

func scheduleCourse4(courses [][]int) int {
	// 贪心，先考虑早结束的课程，按课程的结束时间从早到晚排序
	slices.SortFunc(courses, func(x, y []int) int {
		if x[1] == y[1] {
			return x[0] - y[0]
		}
		return x[1] - y[1]
	})

	day := 0
	h := NewMaxIntHeap()
	for _, course := range courses {
		duration, last := course[0], course[1]

		h.Insert(duration)
		day += duration

		if last < day {
			day -= h.PopHead()
		}
	}

	return h.Len()
}
