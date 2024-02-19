package rbtree

import (
	"bytes"
	"cmp"
)

type color bool

const (
	red   color = true
	black color = false
)

func (c color) String() string {
	if c == red {
		return "🔴"
	}
	return "🔳"
}

// Option 用于修改 Tree 的配置
type Option[K cmp.Ordered, V any] func(t *Tree[K, V])

// WithUpsert 配置 Put 操作时，如何更新已存在节点 Node.Value
func WithUpsert[K cmp.Ordered, V any](upsert func(old, new V) V) Option[K, V] {
	return func(t *Tree[K, V]) {
		t.upsert = upsert
	}
}

// Tree 实现了左倾红黑树，左倾红黑树是 2-3树在二叉树上的转换，因此 2-3 树与左倾红黑树是等价的。
type Tree[K cmp.Ordered, V any] struct {
	root   *Node[K, V]
	upsert func(old, new V) V
}

// NewTree 初始化左倾红黑树，参数 opts 用于修改红黑树的配置
//
// 目前可以修改的配置有
// 1. upsert 字段用于 Put 操作时，如何更新已存在节点的 Value，默认情况下会替换为输入的 value
func NewTree[K cmp.Ordered, V any](opts ...Option[K, V]) *Tree[K, V] {
	tree := &Tree[K, V]{
		root: nil,
		upsert: func(old, new V) V {
			return new
		},
	}
	for _, opt := range opts {
		opt(tree)
	}
	return tree
}

// NewIntTree 初始化 int 类型的左倾红黑树，参数 opts 用于修改红黑树的配置
//
// 目前可以修改的配置有
// 1. upsert 字段用于 Put 操作时，如何更新已存在节点的 Value，默认情况下会替换为输入的 value
func NewIntTree[V any](opts ...Option[int, V]) *Tree[int, V] {
	return NewTree[int, V](opts...)
}

// NewStringTree 初始化 string 类型的左倾红黑树，参数 opts 用于修改红黑树的配置
//
// 目前可以修改的配置有
// 1. upsert 字段用于 Put 操作时，如何更新已存在节点的 Value，默认情况下会替换为输入的 value
func NewStringTree[V any](opts ...Option[string, V]) *Tree[string, V] {
	return NewTree[string, V](opts...)
}

func (t *Tree[K, V]) put(o *Node[K, V], key K, value V) (node *Node[K, V]) {
	if o == nil {
		return &Node[K, V]{
			Left:  nil,
			Right: nil,
			color: red,
			size:  1,
			key:   key,
			Value: value,
		}
	}

	switch {
	case key == o.key:
		o.Value = t.upsert(o.Value, value)
		return o
	case key > o.key:
		o.Right = t.put(o.Right, key, value)
	case key < o.key:
		o.Left = t.put(o.Left, key, value)
	}

	if o.Right.IsRed() && !o.Left.IsRed() { // different from o.balance()
		o = o.rotateLeft()
	}
	if o.Left.IsRed() && o.Left.Left.IsRed() {
		o = o.rotateRight()
	}
	if o.Left.IsRed() && o.Right.IsRed() {
		o.flipColors()
	}

	o.pushUp()
	return o
}

// Put 向树中插入新的节点，如果节点已存在，则按照 t.upsert 更新节点值
// t.upsert 默认情况，节点 *Node.Value 会替换为输入的 value
func (t *Tree[K, V]) Put(key K, value V) (isNew bool) {
	preSize := t.Size()

	t.root = t.put(t.root, key, value)
	t.root.color = black

	return preSize != t.Size()
}

// Update 更新树中指定 key 节点的 value，如果节点不存在，则什么都不做
func (t *Tree[K, V]) Update(key K, value V) {
	get := t.Get(key)
	if get != nil {
		get.Value = value
	}
}

// Get 返回树中寻找指定 key 的节点，如果返回的是 nil, 则代表未找到
func (t *Tree[K, V]) Get(key K) *Node[K, V] {
	for o := t.root; o != nil; {
		switch {
		case key == o.key:
			return o
		case key > o.key:
			o = o.Right
		case key < o.key:
			o = o.Left
		}
	}
	return nil
}

// Floor 寻找给定 key 的 floor 节点，floor 节点定义为：在小于等于 key 节点中的最大值，
// 如果 floor 为 nil，则意味着树是空的，或者树中所有元素都大于 key。
func (t *Tree[K, V]) Floor(key K) (floor *Node[K, V]) {
	for o := t.root; o != nil; {
		switch {
		case key == o.key:
			return o
		case key > o.key:
			floor = o
			o = o.Right
		case key < o.key:
			o = o.Left
		}
	}
	return floor
}

// Ceil 寻找给定 key的 ceil 节点，ceil 节点定义为：在大于等于 key 节点中的最小值
// 如果 ceil 为 nil，则意味着树是空的，或者树中所有元素都小于 key
func (t *Tree[K, V]) Ceil(key K) (ceil *Node[K, V]) {
	for o := t.root; o != nil; {
		switch {
		case key == o.key:
			return o
		case key > o.key:
			o = o.Right
		case key < o.key:
			ceil = o
			o = o.Left
		}
	}
	return ceil
}

func (t *Tree[K, V]) delete(o *Node[K, V], key K) *Node[K, V] {
	if key < o.key {
		// key < o.key, 但是 o.left == nil, 因此删除的 key 不存在
		if o.Left == nil {
			return o.balance()
		}

		// 如果往左走，当前节点是 h，那么需要保证 h 是红色，或者 h->lc 是红色；
		if !o.Left.IsRed() && !o.Left.Left.IsRed() {
			o = o.moveRedLeft()
		}
		o.Left = t.delete(o.Left, key)
	} else {
		if o.Left.IsRed() {
			o = o.rotateRight()
		}
		if key == o.key && o.Right == nil {
			return nil
		}

		// key > o.key, 但是 o.right == nil, 因此删除的 key 不存在
		if o.Right == nil {
			return o.balance()
		}

		// 如果往右走，当前节点是 h，那么需要保证 h 是红色，或者 h->rc 是红色。
		if !o.Right.IsRed() && !o.Right.Left.IsRed() {
			o = o.moveRedRight()
		}
		if key == o.key {
			x := o.Right.Min()
			o.key = x.key
			o.Value = x.Value
			o.Right = o.Right.deleteMin()
		} else {
			o.Right = t.delete(o.Right, key)
		}
	}
	return o.balance()
}

// Delete 从树中删除指定 key 的节点，如果节点不存在，那么什么也不做
func (t *Tree[K, V]) Delete(key K) {
	if t.root == nil {
		return
	}

	if !t.root.Left.IsRed() && !t.root.Right.IsRed() {
		t.root.color = red
	}

	t.root = t.delete(t.root, key)
	if t.root != nil {
		t.root.color = black
	}
}

func (t *Tree[K, V]) Root() *Node[K, V] {
	return t.root
}

func (t *Tree[K, V]) Size() int {
	return t.root.Size()
}

func (t *Tree[K, V]) Min() *Node[K, V] {
	return t.root.Min()
}

func (t *Tree[K, V]) Max() *Node[K, V] {
	return t.root.Max()
}

// Range 按照升序，遍历树中的节点
func (t *Tree[K, V]) Range(f func(o *Node[K, V]) bool) {
	t.root.Range(f)
	return
}

func (t *Tree[K, V]) Keys() []K {
	keys := make([]K, 0, t.root.Size())
	t.Range(func(o *Node[K, V]) bool {
		keys = append(keys, o.key)
		return true
	})
	return keys
}

func (t *Tree[K, V]) Values() []V {
	values := make([]V, 0, t.root.Size())
	t.Range(func(o *Node[K, V]) bool {
		values = append(values, o.Value)
		return true
	})
	return values
}

func (t *Tree[K, V]) Between(left, right K, f func(o *Node[K, V])) {
	t.Range(func(o *Node[K, V]) bool {
		if left <= o.key && o.key <= right {
			f(o)
			return true
		}
		return false
	})
}

func (t *Tree[K, V]) String() string {
	if t.root == nil {
		return "BST (empty)\n"
	}

	buf := bytes.NewBufferString("BST\n")
	t.root.draw("", true, buf)
	return buf.String()
}
