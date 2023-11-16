package main

// minStack5 slice本身就是引用数据类型, 每次操作都需要从地址取值(*s), 实现不如 minStack2 优雅, 且容易埋坑
type minStack5 []node

func Constructor5() Stack {
	return &minStack5{}
}

func (s *minStack5) Push(val int) {
	minValue := val
	n := len(*s)
	if n > 0 {
		minValue = min(minValue, (*s)[n-1].min)
	}

	*s = append(*s, node{
		data: val,
		min:  minValue,
	})
}

func (s *minStack5) Pop() {
	*s = (*s)[:len(*s)-1]
}

func (s *minStack5) Top() int {
	return (*s)[len(*s)-1].data
}

func (s *minStack5) GetMin() int {
	return (*s)[len(*s)-1].min
}
