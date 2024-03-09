package main

import (
	"sort"
	"strings"

	. "github.com/niluan304/leetcode/container"
)

// 哈希 + 快速排序
// 时间复杂度：O(nlogn)
// 空间复杂度：O(logn)
func topStudents(positiveFeedback []string, negativeFeedback []string, report []string, studentId []int, k int) []int {
	words := make(map[string]int, len(positiveFeedback)+len(negativeFeedback))
	for _, w := range positiveFeedback {
		words[w] = 3
	}
	for _, w := range negativeFeedback {
		words[w] = -1
	}

	type Pair struct {
		id    int
		score int
	}
	pairs := make([]Pair, len(studentId))
	for i, id := range studentId {
		score := 0
		for _, w := range strings.Split(report[i], " ") {
			score += words[w]
		}
		pairs[i] = Pair{
			id:    id,
			score: score,
		}
	}
	sort.Slice(pairs, func(i, j int) bool {
		x, y := pairs[i], pairs[j]
		if x.score == y.score {
			return x.id < y.id
		}
		return x.score > y.score
	})

	ans := make([]int, k)
	for i, score := range pairs[:k] {
		ans[i] = score.id
	}
	return ans
}

// 哈希 + 堆排序（优先队列）
// 时间复杂度：O(klogn)
// 空间复杂度：O(logn)
func topStudents2(positiveFeedback []string, negativeFeedback []string, report []string, studentId []int, k int) []int {
	words := make(map[string]int, len(positiveFeedback)+len(negativeFeedback))
	for _, w := range positiveFeedback {
		words[w] = 3
	}
	for _, w := range negativeFeedback {
		words[w] = -1
	}

	type item struct {
		id    int // value
		score int // priority
	}
	hp := NewEmptyHeap(func(x, y item) bool {
		if x.score == y.score {
			return x.id < y.id
		}
		return x.score > y.score
	})

	for i, id := range studentId {
		score := 0
		for _, w := range strings.Split(report[i], " ") {
			score += words[w]
		}
		hp.Push(item{id: id, score: score})
	}

	ans := make([]int, 0, k)
	for i := 0; i < k; i++ {
		ans = append(ans, hp.Pop().id)
	}
	return ans
}
