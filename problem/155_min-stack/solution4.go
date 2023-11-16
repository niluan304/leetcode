package main

type Node struct {
	val  int
	min  int
	next *Node
}

// minStack4 用链表的方式实现 minStack2 操作
type minStack4 struct {
	head *Node
}

func Constructor4() Stack {
	return &minStack4{}
}

func (s *minStack4) Push(val int) {
	minValue := val
	if s.head != nil {
		minValue = min(s.head.min, val)
	}

	s.head = &Node{
		val:  val,
		min:  minValue,
		next: s.head,
	}
}

func (s *minStack4) Pop() {
	s.head = s.head.next
}

func (s *minStack4) Top() int {
	return s.head.val
}

func (s *minStack4) GetMin() int {
	return s.head.min
}
