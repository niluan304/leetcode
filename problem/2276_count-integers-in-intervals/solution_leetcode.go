package main

type CountIntervals2 struct {
	left, right *CountIntervals2
	l, r, cnt   int
}

func Constructor2() CountIntervals2 { return CountIntervals2{l: 1, r: 1e9} }

func (o *CountIntervals2) Add(l, r int) {
	if o.cnt == o.r-o.l+1 {
		return
	} // o 已被完整覆盖，无需执行任何操作
	if l <= o.l && o.r <= r { // 当前节点已被区间 [l,r] 完整覆盖，不再继续递归
		o.cnt = o.r - o.l + 1
		return
	}
	mid := (o.l + o.r) >> 1
	if o.left == nil {
		o.left = &CountIntervals2{l: o.l, r: mid}
	} // 动态开点
	if o.right == nil {
		o.right = &CountIntervals2{l: mid + 1, r: o.r}
	} // 动态开点
	if l <= mid {
		o.left.Add(l, r)
	}
	if mid < r {
		o.right.Add(l, r)
	}
	o.cnt = o.left.cnt + o.right.cnt
}

func (o *CountIntervals2) Count() int { return o.cnt }
