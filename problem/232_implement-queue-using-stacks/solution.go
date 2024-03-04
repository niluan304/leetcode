package main

type MyQueue struct {
	data Stack
	tmp  Stack
}

func Constructor() MyQueue {
	return MyQueue{
		data: make(Stack, 0),
		tmp:  make(Stack, 0),
	}
}

func (q *MyQueue) Push(x int) {
	for !q.data.Empty() {
		q.tmp.Push(*q.data.Pop())
	}
	q.data.Push(x)
	for !q.tmp.Empty() {
		q.data.Push(*(q.tmp.Pop()))
	}
}

func (q *MyQueue) Pop() int    { return *(q.data.Pop()) }
func (q *MyQueue) Peek() int   { return *(q.data.Top()) }
func (q *MyQueue) Empty() bool { return q.data.Empty() }

type MyQueue2 struct {
	in, out Stack
}

func Constructor2() MyQueue2 {
	return MyQueue2{
		in:  make(Stack, 0),
		out: make(Stack, 0),
	}
}

func (q *MyQueue2) Push(x int) {
	q.in.Push(x)
}

func (q *MyQueue2) Pop() int {
	if q.out.Empty() {
		q.in2out()
	}
	return *q.out.Pop()
}

func (q *MyQueue2) Peek() int {
	if q.out.Empty() {
		q.in2out()
	}
	return *q.out.Top()
}

func (q *MyQueue2) Empty() bool {
	return q.in.Empty() && q.out.Empty()
}

func (q *MyQueue2) in2out() {
	for !q.in.Empty() {
		q.out.Push(*(q.in.Pop()))
	}
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

type Stack []int

func (s *Stack) Push(v int) {
	*s = append(*s, v)
}

func (s *Stack) Pop() *int {
	n := len(*s)
	if n == 0 {
		return nil
	}
	x := (*s)[n-1]
	*s = (*s)[:n-1]
	return &x
}

func (s Stack) Size() int {
	return len(s)
}

func (s Stack) Empty() bool {
	return len(s) == 0
}

func (s *Stack) Top() *int {
	n := len(*s)
	if n == 0 {
		return nil
	}
	return &(*s)[n-1]
}
