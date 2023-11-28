package main

type FrontMiddleBackQueue struct {
	queen []int
}

func Constructor() FrontMiddleBackQueue {
	return FrontMiddleBackQueue{
		queen: make([]int, 0),
	}
}

func (q *FrontMiddleBackQueue) PushFront(val int) {
	q.queen = append([]int{val}, q.queen...)
}

func (q *FrontMiddleBackQueue) PushMiddle(val int) {
	n := q.Len()
	i := n / 2

	var queen = append([]int{}, q.queen[:i]...)
	queen = append(queen, val)
	q.queen = append(queen, q.queen[i:]...)
}

func (q *FrontMiddleBackQueue) PushBack(val int) {
	q.queen = append(q.queen, val)
}

func (q *FrontMiddleBackQueue) PopFront() int {
	n := q.Len()
	if n == 0 {
		return -1
	}
	ans := q.queen[0]
	q.queen = q.queen[1:]
	return ans
}

func (q *FrontMiddleBackQueue) PopMiddle() int {
	n := q.Len()
	if n == 0 {
		return -1
	}
	i := (n - 1) / 2
	ans := q.queen[i]
	q.queen = append(q.queen[:i], q.queen[i+1:]...)
	return ans
}

func (q *FrontMiddleBackQueue) PopBack() int {
	n := q.Len()
	if n == 0 {
		return -1
	}
	ans := q.queen[n-1]
	q.queen = q.queen[:n-1]
	return ans
}

func (q *FrontMiddleBackQueue) Len() int {
	return len(q.queen)
}

/**
 * Your FrontMiddleBackQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PushFront(val);
 * obj.PushMiddle(val);
 * obj.PushBack(val);
 * param_4 := obj.PopFront();
 * param_5 := obj.PopMiddle();
 * param_6 := obj.PopBack();
 */
