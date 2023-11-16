package main

type node struct {
	data int
	min  int
}

// MinStack 单个栈元素除了保存当前值外, 还额外保存当前最小值
type MinStack struct {
	nodes []node
	len   int
}

func Constructor() MinStack {
	return MinStack{
		nodes: make([]node, 0, 3000),
	}
}

func (s *MinStack) Push(val int) {
	minValue := val
	n := len(s.nodes)
	if n > 0 {
		minValue = min(minValue, s.nodes[n-1].min)
	}

	s.nodes = append(s.nodes, node{
		data: val,
		min:  minValue,
	})
	s.len++
}

func (s *MinStack) Pop() {
	s.nodes = s.nodes[:s.len-1]
	s.len--
}

func (s *MinStack) Top() int {
	return s.nodes[s.len-1].data
}

func (s *MinStack) GetMin() int {
	return s.nodes[s.len-1].min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
