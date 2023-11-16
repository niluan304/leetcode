package main

// minStack3 辅助栈保存最小值的顺序
type minStack3 struct {
	min   []*int
	stack []*int
}

func Constructor3() Stack {
	return &minStack3{
		min:   make([]*int, 0, 10),
		stack: make([]*int, 0, 10),
	}
}

func (s *minStack3) Push(val int) {
	s.stack = append(s.stack, &val)

	if len(s.min) == 0 || *s.min[len(s.min)-1] > val {
		s.min = append(s.min, &val)
	}
}

func (s *minStack3) Pop() {
	last := len(s.stack) - 1
	if s.stack[last] == s.min[len(s.min)-1] {
		s.min = s.min[:len(s.min)-1]
	}
	s.stack = s.stack[:last]

}

func (s *minStack3) Top() int {
	return *s.stack[len(s.stack)-1]
}

func (s *minStack3) GetMin() int {
	return *s.min[len(s.min)-1]
}
