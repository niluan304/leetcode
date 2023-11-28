package main

import "container/list"

// FrontMiddleBackQueue2 第一种写法：链表
type FrontMiddleBackQueue2 struct {
	left  *list.List
	right *list.List
}

// Constructor2
// ![LC1670-c.png](https://pic.leetcode.cn/1701084361-JejjZq-LC1670-c.png)
// ### 复杂度分析
//
// - 时间复杂度：所有操作均为 $\mathcal{O}(1)$。
// - 空间复杂度：$\mathcal{O}(q)$，其中 $q$ 为操作次数。
//
// 欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
//
// 更多精彩题解，请看 [往期题解精选（按 tag 分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
func Constructor2() FrontMiddleBackQueue2 {
	return FrontMiddleBackQueue2{
		left:  list.New(),
		right: list.New(),
	}
}

// 调整长度，保证 0 <= right.Len() - left.Len() <= 1
// 从而保证可以在正中间插入删除元素
func (q *FrontMiddleBackQueue2) balance() {
	if q.left.Len() > q.right.Len() {
		q.right.PushFront(q.left.Remove(q.left.Back()))
	} else if q.right.Len() > q.left.Len()+1 {
		q.left.PushBack(q.right.Remove(q.right.Front()))
	}
}

func (q *FrontMiddleBackQueue2) PushFront(val int) {
	q.left.PushFront(val)
	q.balance()
}

func (q *FrontMiddleBackQueue2) PushMiddle(val int) {
	if q.left.Len() < q.right.Len() {
		q.left.PushBack(val)
	} else {
		q.right.PushFront(val)
	}
}

func (q *FrontMiddleBackQueue2) PushBack(val int) {
	q.right.PushBack(val)
	q.balance()
}

func (q *FrontMiddleBackQueue2) PopFront() (val int) {
	if q.right.Len() == 0 { // 整个队列为空
		return -1
	}
	if q.left.Len() > 0 {
		val = q.left.Remove(q.left.Front()).(int)
	} else {
		val = q.right.Remove(q.right.Front()).(int)
	}
	q.balance()
	return
}

func (q *FrontMiddleBackQueue2) PopMiddle() int {
	if q.right.Len() == 0 { // 整个队列为空
		return -1
	}
	if q.left.Len() == q.right.Len() {
		return q.left.Remove(q.left.Back()).(int)
	}
	return q.right.Remove(q.right.Front()).(int)
}

func (q *FrontMiddleBackQueue2) PopBack() int {
	if q.right.Len() == 0 { // 整个队列为空
		return -1
	}
	val := q.right.Remove(q.right.Back()).(int)
	q.balance()
	return val
}
