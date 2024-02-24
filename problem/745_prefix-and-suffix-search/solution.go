package main

type WordFilter struct {
	trie  *Trie[byte, int]
	cache map[string]int
}

// Constructor
// 由于单个单词长度很短，仅 [1,7]，可以暴力穷举出所有后缀+前缀的组合
func Constructor(words []string) WordFilter {
	trie := NewTrie[byte, int]()
	for i, word := range words {
		for j := 0; j <= len(word); j++ {
			elem := word[j:] + sep + word
			trie.InsertWith([]byte(elem), i, func(o *Node[byte, int]) {
				o.Value = i
			})
		}
	}
	return WordFilter{
		trie:  trie,
		cache: map[string]int{},
	}
}

const sep = "{"

func (w *WordFilter) F(prefix, suffix string) (res int) {
	word := suffix + sep + prefix
	if v, ok := w.cache[word]; ok {
		return v
	}
	defer func() { w.cache[word] = res }()

	node := w.trie.Find([]byte(word))
	if node == nil {
		return -1
	}
	return node.Value
}

/**
 * Your WordFilter object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.F(pref,suff);
 */

type Node[E comparable, V any] struct {
	son   map[E]*Node[E, V]
	count int //（子树中）完整字符串的个数
	Value V
}

func (o *Node[E, V]) Count() int {
	return o.count
}

func (o *Node[E, V]) Empty() bool {
	for _, son := range o.son {
		if son != nil {
			return false
		}
	}
	return true
}

type Option[E comparable, V any] func(t *Trie[E, V])

func WithUpsert[E comparable, V any](upsert func(old, new V) V) Option[E, V] {
	return func(t *Trie[E, V]) {
		t.upsert = upsert
	}
}

// Trie 实现泛型的字典树（前缀树）
type Trie[E comparable, V any] struct {
	root   *Node[E, V]
	upsert func(old, new V) V
}

func NewTrie[E comparable, V any](opts ...Option[E, V]) *Trie[E, V] {
	// init with a root (empty string)
	t := &Trie[E, V]{
		root: &Node[E, V]{
			son:   make(map[E]*Node[E, V]),
			count: 0,
		},
		upsert: func(old, new V) V { return new },
	}
	for _, opt := range opts {
		opt(t)
	}
	return t
}

func (t *Trie[E, V]) Insert(elems []E, value V) {
	cur := t.root
	for _, elem := range elems {
		if cur.son[elem] == nil {
			cur.son[elem] = &Node[E, V]{
				son:   make(map[E]*Node[E, V]),
				count: 0,
			}
		}
		cur = cur.son[elem]
	}
	cur.count++
	cur.Value = value
}

func (t *Trie[E, V]) InsertWith(elems []E, value V, f func(o *Node[E, V])) {
	cur := t.root
	for _, elem := range elems {
		if cur.son[elem] == nil {
			cur.son[elem] = &Node[E, V]{
				son:   make(map[E]*Node[E, V]),
				count: 0,
			}
		}
		cur = cur.son[elem]
		f(cur)
	}
	if cur.count == 0 {
		cur.Value = value
	} else {
		cur.Value = t.upsert(cur.Value, value)
	}
	cur.count++
}

func (t *Trie[E, V]) Delete(elems []E) *Node[E, V] {
	n := len(elems)
	fa := make([]*Node[E, V], n)
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
			f.son[elems[i]] = nil
			if !f.Empty() {
				break
			}
		}
	}
	return cur
}

func (t *Trie[E, V]) Find(elems []E) *Node[E, V] {
	cur := t.root
	for _, elem := range elems {
		if cur.son[elem] == nil {
			return nil
		}
		cur = cur.son[elem]
	}
	return cur
}

type WordFilter2 struct {
	data map[string]int
}

// Constructor2
// 注意到单词长度为 [1, 7] 很短，因此可以暴力穷举所有前后缀的组合
func Constructor2(words []string) WordFilter2 {
	data := make(map[string]int, len(words)*49)
	for i, word := range words {
		for j := 0; j <= len(word); j++ {
			elem := word[j:] + sep + word
			for k := 0; k <= len(elem); k++ {
				data[elem[:k]] = i + 1
			}
		}
	}
	return WordFilter2{
		data: data,
	}
}

func (w *WordFilter2) F(prefix, suffix string) int {
	word := suffix + sep + prefix
	return w.data[word] - 1
}
