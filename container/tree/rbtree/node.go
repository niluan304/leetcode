package rbtree

import (
	"bytes"
	"cmp"
	"fmt"
)

type Node[K cmp.Ordered, V any] struct {
	Left  *Node[K, V]
	Right *Node[K, V]
	color color // 节点颜色， red or black
	size  int   // 子树中树的个数
	key   K
	Value V
}

// IsRed 返回节点是否为红色
// 如果结点为 nil，代表是叶子结点，那么为黑色，返回 false
func (o *Node[K, V]) IsRed() bool {
	return o != nil && o.color == red
}

func (o *Node[K, V]) rotateLeft() *Node[K, V] {
	//          |                                  |
	//          O                                  X
	//         / \         left rotate            / \
	//        α   X       ------------->         O   γ
	//           / \                            / \
	//          β   γ                          α   β

	x := o.Right
	o.Right = x.Left  // 当前节点的右儿子变为 x 的左儿子
	x.Left = o        // 调转红箭头
	x.color = o.color // 更新节点颜色
	o.color = red     // 新节点一定为红色

	o.pushUp()
	x.pushUp()
	return x
}

func (o *Node[K, V]) rotateRight() *Node[K, V] {
	//          |                                  |
	//          O                                  X
	//         / \         right rotate           / \
	//        X   γ       ------------->         α   O
	//       / \                                    / \
	//      α   β                                  β   γ

	x := o.Left
	o.Left = x.Right  // 当前节点的左儿子变为 x 的右儿子
	x.Right = o       // 调转红箭头
	x.color = o.color // 更新节点颜色
	o.color = red     // 新节点一定为红色

	o.pushUp()
	x.pushUp()
	return x
}

func (o *Node[K, V]) flipColors() {
	o.color = !o.color
	o.Left.color = !o.Left.color
	o.Right.color = !o.Right.color
}

func (o *Node[K, V]) balance() *Node[K, V] {
	if o.Right.IsRed() { // case 1
		o = o.rotateLeft()
	}
	if o.Left.IsRed() && o.Left.Left.IsRed() { // case 2
		o = o.rotateRight()
	}
	if o.Left.IsRed() && o.Right.IsRed() { // case 3
		o.flipColors()
	}

	o.pushUp()
	return o
}

func (o *Node[K, V]) pushUp() {
	o.size = 1 + o.Left.Size() + o.Right.Size()
}

func (o *Node[K, V]) moveRedLeft() *Node[K, V] {
	o.flipColors()
	if o.Right.Left.IsRed() {
		o.Right = o.Right.rotateRight()
		o = o.rotateLeft()
		o.flipColors()
	}
	return o
}

func (o *Node[K, V]) moveRedRight() *Node[K, V] {
	o.flipColors()
	if o.Left.Left.IsRed() {
		o = o.rotateRight()
		o.flipColors()
	}
	return o
}

func (o *Node[K, V]) deleteMin() *Node[K, V] {
	if o.Left == nil {
		return nil
	}
	if !o.Left.IsRed() && !o.Left.Left.IsRed() { // 23 树视角下，当前节点容量只有 1
		o = o.moveRedLeft()
	}
	o.Left = o.Left.deleteMin()
	return o.balance()
}

func (o *Node[K, V]) Size() int {
	if o == nil {
		return 0
	}
	return o.size
}

func (o *Node[K, V]) Max() (node *Node[K, V]) {
	for o.Right != nil {
		o = o.Right
	}
	return o
}

func (o *Node[K, V]) Min() (node *Node[K, V]) {
	for o.Left != nil {
		o = o.Left
	}
	return o
}

// Range 中序遍历红黑树，也就是升序遍历
func (o *Node[K, V]) Range(f func(o *Node[K, V]) bool) {
	if o == nil {
		return
	}

	o.Left.Range(f)
	if f(o) {
		o.Right.Range(f)
	}
}

func (o *Node[K, V]) Key() K {
	return o.key
}

func (o *Node[K, V]) KeyValue() (K, V) {
	return o.key, o.Value
}

func (o *Node[K, V]) String() string {
	return fmt.Sprintf("%v,size:%d | %v(%v)", o.color, o.size, o.key, o.Value)
}

func (o *Node[K, V]) draw(prefix string, isTail bool, buf *bytes.Buffer) {
	if o.Right != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}
		o.Right.draw(newPrefix, false, buf)
	}

	buf.WriteString(prefix)
	if isTail {
		buf.WriteString("└── ")
	} else {
		buf.WriteString("┌── ")
	}
	buf.WriteString(o.String() + "\n")

	if o.Left != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		o.Left.draw(newPrefix, true, buf)
	}
}
