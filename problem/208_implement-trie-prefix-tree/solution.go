package main

type Node struct {
	son   map[rune]*Node
	count int
}

func (o *Node) Empty() bool {
	for _, son := range o.son {
		if son != nil {
			return false
		}
	}
	return true
}

type Trie struct {
	root *Node
}

func Constructor() Trie {
	return Trie{
		root: &Node{
			son:   make(map[rune]*Node),
			count: 0,
		},
	}
}

func (t *Trie) Insert(word string) {
	cur := t.root
	for _, b := range word {
		if cur.son[b] == nil {
			cur.son[b] = &Node{
				son:   make(map[rune]*Node),
				count: 0,
			}
		}
		cur = cur.son[b]
	}
	cur.count++
}

func (t *Trie) find(word string) *Node {
	cur := t.root
	for _, w := range word {
		if cur.son[w] == nil {
			return nil
		}
		cur = cur.son[w]
	}
	return cur
}

func (t *Trie) Search(word string) bool {
	o := t.find(word)
	return o != nil && o.count > 0
}

func (t *Trie) StartsWith(prefix string) bool {
	o := t.find(prefix)
	return o != nil
}

func (t *Trie) Delete(elems string) *Node {
	n := len(elems)
	fa := make([]*Node, n)
	cur := t.root
	for i, elem := range elems {
		fa[i] = cur
		cur = cur.son[elem]
		if cur == nil {
			return nil
		}
	}

	cur.count--
	if cur.count == 0 {
		for i := n - 1; i >= 0; i-- {
			f := fa[i]
			f.son[rune(elems[i])] = nil
			if !f.Empty() {
				break
			}
		}
	}
	return cur
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
