package main

type Stack interface {
	Push(val int)
	Pop()
	Top() int
	GetMin() int
}

type node2 struct {
	data int
	min  int
}

// minStack2 单个栈元素除了保存当前值外, 还额外保存当前最小值
type minStack2 struct {
	nodes []node2
}

func Constructor2() Stack {
	return &minStack2{}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func (s *minStack2) Push(val int) {
	minValue := val
	n := len(s.nodes)
	if n > 0 {
		minValue = min(minValue, s.nodes[n-1].min)
	}

	s.nodes = append(s.nodes, node2{
		data: val,
		min:  minValue,
	})
}

func (s *minStack2) Pop() {
	s.nodes = s.nodes[:s.tailIndex()]
}

func (s *minStack2) Top() int {
	return s.nodes[s.tailIndex()].data
}

func (s *minStack2) GetMin() int {
	return s.nodes[s.tailIndex()].min
}

// alias of len(s.nodes) - 1
func (s *minStack2) tailIndex() int {
	return len(s.nodes) - 1
}
