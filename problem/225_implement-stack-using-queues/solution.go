package main

type MyStack struct {
	data Queen
	tmp  Queen
}

func Constructor() MyStack {
	return MyStack{
		data: make(Queen, 0),
		tmp:  make(Queen, 0),
	}
}

func (s *MyStack) Push(x int) {
	s.tmp.Push(x)
	for s.data.Size() > 0 {
		s.tmp.Push(*s.data.Pop())
	}
	s.tmp, s.data = s.data, s.tmp // 置空 s.tmp
}

func (s *MyStack) Pop() int {
	return *s.data.Pop()
}

func (s *MyStack) Top() int {
	return *s.data.Head()
}

func (s *MyStack) Empty() bool {
	return s.data.Size() == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

type Queen []int

func (q Queen) Size() int {
	return len(q)
}

func (q *Queen) Push(v int) {
	*q = append(*q, v)
}

func (q *Queen) Pop() *int {
	if q.Size() == 0 {
		return nil
	}
	x := (*q)[0]
	*q = (*q)[1:]
	return &x
}

func (q *Queen) Head() *int {
	if q.Size() == 0 {
		return nil
	}
	return &(*q)[0]
}
