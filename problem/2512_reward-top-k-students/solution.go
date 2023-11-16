package main

import (
	"container/heap"
	"sort"
	"strings"
)

// 哈希 + 快速排序
// 时间复杂度：O(nlogn)
// 空间复杂度：O(logn)
func topStudents(positiveFeedback []string, negativeFeedback []string, report []string, studentId []int, k int) []int {
	var words = make(map[string]int, len(positiveFeedback)+len(negativeFeedback))
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
	var pairs = make([]Pair, len(studentId))
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

	var ans = make([]int, k)
	for i, score := range pairs[:k] {
		ans[i] = score.id
	}
	return ans
}

// 哈希 + 堆排序（优先队列）
// 时间复杂度：O(klogn)
// 空间复杂度：O(logn)
func topStudents2(positiveFeedback []string, negativeFeedback []string, report []string, studentId []int, k int) []int {
	var words = make(map[string]int, len(positiveFeedback)+len(negativeFeedback))
	for _, w := range positiveFeedback {
		words[w] = 3
	}
	for _, w := range negativeFeedback {
		words[w] = -1
	}
	var pq = new(priorityQueue)
	for i, id := range studentId {
		score := 0
		for _, w := range strings.Split(report[i], " ") {
			score += words[w]
		}
		heap.Push(pq, &item{
			id:    id,
			score: score,
		})
	}

	var ans = make([]int, 0, k)
	for i := 0; i < k; i++ {
		ans = append(ans, heap.Pop(pq).(*item).id)
	}
	return ans
}

type item struct {
	id    int // value
	score int // priority
}

// A PriorityQueue implements heap.Interface and holds Items.
type priorityQueue []*item

func (pq priorityQueue) Len() int { return len(pq) }
func (pq priorityQueue) Less(i, j int) bool {
	if pq[i].score == pq[j].score {
		return pq[i].id < pq[j].id
	}
	return pq[i].score > pq[j].score
}
func (pq priorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *priorityQueue) Push(x any) { *pq = append(*pq, x.(*item)) }

func (pq *priorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}
