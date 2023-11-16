package main

type LockingTree struct {
	parent   *LockingTree
	children []*LockingTree
	num      int
	user     *int
	list     []*LockingTree
}

func Constructor(parent []int) LockingTree {
	var n = len(parent)
	var list = make([]*LockingTree, n)

	for i := 0; i < n; i++ {
		list[i] = new(LockingTree)
	}
	for i := 1; i < n; i++ {
		node := list[parent[i]]
		list[i].parent = node
		list[i].num = i

		node.children = append(node.children, list[i])
	}
	list[0].list = list
	return *list[0]
}

func (t *LockingTree) Lock(num int, user int) bool {
	node := t.list[num]
	if node.user != nil {
		return false
	}
	node.user = &user
	return true
}

func (t *LockingTree) Unlock(num int, user int) bool {
	node := t.list[num]
	if node.user == nil || *node.user != user {
		return false
	}
	node.user = nil
	return true
}

func (t *LockingTree) Upgrade(num int, user int) bool {
	node := t.list[num]
	// 指定节点当前状态为未上锁。
	if node.user != nil {
		return false
	}
	// 指定节点没有任何上锁的祖先节点。
	if node.HasLockParent() {
		return false
	}

	var success = false
	// 指定节点至少有一个上锁状态的子孙节点（可以是 任意 用户上锁的）。
	node.AllChildren(func(child *LockingTree) (stop bool) {
		if child.user == nil {
			return false

		}

		node.AllChildren(func(child *LockingTree) (stop bool) {
			// 并且将该节点的所有子孙节点 解锁 。
			child.user = nil
			return false
		})

		// 不能前置, BFS中会解锁自己
		_ = t.Lock(num, user)
		success = true
		return true

	})
	return success
}

// HasLockParent DFS
func (t *LockingTree) HasLockParent() bool {
	if t == nil {
		return false
	}
	if t.user != nil {
		return true
	}
	return t.parent.HasLockParent()
}

// AllChildren BFS 遍历所有children节点, 包括节点自身
func (t *LockingTree) AllChildren(fn func(child *LockingTree) (stop bool)) {
	var queen = make([]*LockingTree, 0, len(t.list))
	queen = append(queen, t)
	for len(queen) > 0 {
		head := queen[0]
		if fn(head) {
			return
		}
		queen = queen[1:]
		queen = append(queen, head.children...)
	}
}

/**
 * Your LockingTree object will be instantiated and called as such:
 * obj := Constructor(parent);
 * param_1 := obj.Lock(num,user);
 * param_2 := obj.Unlock(num,user);
 * param_3 := obj.Upgrade(num,user);
 */
