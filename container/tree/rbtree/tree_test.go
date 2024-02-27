package rbtree

import (
	"cmp"
	"fmt"
	"math"
	"slices"
	"testing"
)

type TreeWithValid[K cmp.Ordered, V any] struct {
	*Tree[K, V]
}

func NewTreeWithValid[K cmp.Ordered, V any](opts ...Option[K, V]) *TreeWithValid[K, V] {
	return &TreeWithValid[K, V]{
		Tree: NewTree(opts...),
	}
}

func NewIntTreeWithValid[V any](opts ...Option[int, V]) *TreeWithValid[int, V] {
	return &TreeWithValid[int, V]{
		Tree: NewIntTree(opts...),
	}
}

func NewStringTreeWithValid[V any](opts ...Option[string, V]) *TreeWithValid[string, V] {
	return &TreeWithValid[string, V]{
		Tree: NewStringTree(opts...),
	}
}

func (t *TreeWithValid[K, V]) Put(key K, value V) (isNew bool) {
	defer valid(t.Tree)
	return t.Tree.Put(key, value)
}

func (t *TreeWithValid[K, V]) Update(key K, value V) {
	defer valid(t.Tree)
	t.Tree.Update(key, value)
}

func (t *TreeWithValid[K, V]) Delete(key K) {
	defer valid(t.Tree)
	t.Tree.Delete(key)
}

type Pair[K cmp.Ordered, V any] struct {
	key   K
	value V
}

func valid[K cmp.Ordered, V any](t *Tree[K, V]) {
	root := t.Root()
	if root == nil {
		return
	}

	if !isBST(root, t.Min().Key(), t.Max().Key()) {
		panic(fmt.Errorf("%+v is not valid Left-leaning red-black tree", root))
	}
	if !isRedNodesValid(root) {
		panic(fmt.Errorf("%+v is not valid Left-leaning red-black tree", root))
	}
	if ok, _ := isBlackNodesValid(root, 0); !ok {
		panic(fmt.Errorf("%+v is not valid Left-leaning red-black tree", root))
	}
}

func isBST[K cmp.Ordered, V any](node *Node[K, V], left, right K) bool {
	if node == nil {
		return true
	}

	key := node.Key()
	if left <= key && key <= right {
		return isBST(node.Left, left, key) && isBST(node.Right, key, right)
	}
	return false
}

func isRedNodesValid[K cmp.Ordered, V any](node *Node[K, V]) bool {
	if node == nil {
		return !node.IsRed() // 性质3： 叶子节点是黑色的。
	}
	if node.IsRed() {
		if node.Left.IsRed() || node.Right.IsRed() { // 性质4：红色节点的两个子节点都是黑色
			return false
		}
	}
	return isRedNodesValid(node.Left) && isRedNodesValid(node.Right)
}

func isBlackNodesValid[K cmp.Ordered, V any](node *Node[K, V], count int) (bool, int) {
	if node == nil {
		return true, count
	}

	if !node.IsRed() {
		count++
	}
	validLeft, countLeft := isBlackNodesValid(node.Left, count)
	validRight, countRight := isBlackNodesValid(node.Right, count)

	// 性质5：从任一节点到其每个叶子的所有路径都包含相同数目的黑色节点。
	return validLeft && validRight && countLeft == countRight, countLeft
}

func TestTree_Put(t *testing.T) {
	tree := NewTreeWithValid[int, string]()
	pairs := []Pair[int, string]{
		{key: 10, value: "a"},
		{key: 30, value: "c"},
		{key: 40, value: "d"},
		{key: 50, value: "e"},
		{key: 20, value: "b"},
		{key: 60, value: "f"},
	}

	for _, pair := range pairs {
		tree.Put(pair.key, pair.value)
	}

	fmt.Println(tree)

	{
		k, v := 10, "a"
		key, value := tree.Get(k).KeyValue()
		if key != k || value != v {
			t.Errorf("tree.Get(%v).KeyValue() should be: (%v, %v), but got: (%v, %v)", k, k, v, key, value)
		}
	}
	{
		k, v := 60, "f"
		key, value := tree.Get(k).KeyValue()
		if key != k || value != v {
			t.Errorf("tree.Get(%v).KeyValue() should be: (%v, %v), but got: (%v, %v)", k, k, v, key, value)
		}
	}

	{
		k, v, v2 := 22, "bb", "bB"
		isNew := tree.Put(k, v)
		if !isNew {
			t.Errorf("(%v, %v) should be new node, but got: %v", k, v, isNew)
		}

		key, value := tree.Get(k).KeyValue()
		if key != k || value != v {
			t.Errorf("tree.Get(%v).KeyValue() should be: (%v, %v), but got: (%v, %v)", k, k, v, key, value)
		}

		isNew2 := tree.Put(k, v2)
		if isNew2 {
			t.Errorf("(%v, %v) should be old node, but got: %v", k, v2, isNew2)
		}

		key2, value2 := tree.Get(k).KeyValue()
		if key2 != k || value2 != v2 {
			t.Errorf("tree.Get(%v).KeyValue() should be: (%v, %v), but got: (%v, %v)", k, k, v2, key, value2)
		}
	}

	fmt.Println(tree)

	{
		k, v := 34, "__34__"
		isNew := tree.Put(k, v)
		if !isNew {
			t.Errorf("(%v, %v) should be new node, but got: %v", k, v, isNew)
		}

		key, value := tree.Get(k).KeyValue()
		if key != k || value != v {
			t.Errorf("tree.Get(%v).KeyValue() should be: (%v, %v), but got: (%v, %v)", k, k, v, key, value)
		}
	}
	fmt.Println(tree)
}

func TestTree_Put_IntTree(t *testing.T) {
	tree := NewIntTreeWithValid[string]()
	pairs := []Pair[int, string]{
		{key: 10, value: "a"},
		{key: 30, value: "c"},
		{key: 40, value: "d"},
		{key: 50, value: "e"},
		{key: 20, value: "b"},
		{key: 60, value: "f"},
	}

	for _, pair := range pairs {
		tree.Put(pair.key, pair.value)
	}

	fmt.Println(tree)

	{
		k, v := 10, "a"
		key, value := tree.Get(k).KeyValue()
		if key != k || value != v {
			t.Errorf("tree.Get(%v).KeyValue() should be: (%v, %v), but got: (%v, %v)", k, k, v, key, value)
		}
	}
	{
		k, v := 60, "f"
		key, value := tree.Get(k).KeyValue()
		if key != k || value != v {
			t.Errorf("tree.Get(%v).KeyValue() should be: (%v, %v), but got: (%v, %v)", k, k, v, key, value)
		}
	}

	{
		k, v, v2 := 22, "bb", "BB"
		isNew := tree.Put(k, v)
		if !isNew {
			t.Errorf("(%v, %v) should be new node, but got: %v", k, v, isNew)
		}

		key, value := tree.Get(k).KeyValue()
		if key != k || value != v {
			t.Errorf("tree.Get(%v).KeyValue() should be: (%v, %v), but got: (%v, %v)", k, k, v, key, value)
		}

		isNew2 := tree.Put(k, v2)
		if isNew2 {
			t.Errorf("(%v, %v) should be old node, but got: %v", k, v2, isNew2)
		}

		key2, value2 := tree.Get(k).KeyValue()
		if key2 != k || value2 != v2 {
			t.Errorf("tree.Get(%v).KeyValue() should be: (%v, %v), but got: (%v, %v)", k, k, v2, key, value2)
		}
	}

	fmt.Println(tree)

	{
		k, v := 34, "cf"
		isNew := tree.Put(k, v)
		if !isNew {
			t.Errorf("(%v, %v) should be new node, but got: %v", k, v, isNew)
		}

		key, value := tree.Get(k).KeyValue()
		if key != k || value != v {
			t.Errorf("tree.Get(%v).KeyValue() should be: (%v, %v), but got: (%v, %v)", k, k, v, key, value)
		}
	}
	fmt.Println(tree)
}

func TestTree_Put_StringTree(t *testing.T) {
	tree := NewStringTreeWithValid[int]()
	pairs := []Pair[string, int]{
		{key: "a", value: 10},
		{key: "c", value: 30},
		{key: "d", value: 40},
		{key: "e", value: 50},
		{key: "b", value: 20},
		{key: "f", value: 60},
	}

	for _, pair := range pairs {
		tree.Put(pair.key, pair.value)
	}

	fmt.Println(tree)

	{
		k, v := "a", 10
		key, value := tree.Get(k).KeyValue()
		if key != k || value != v {
			t.Errorf("tree.Get(%v).KeyValue() should be: (%v, %v), but got: (%v, %v)", k, k, v, key, value)
		}
	}
	{
		k, v := "f", 60
		key, value := tree.Get(k).KeyValue()
		if key != k || value != v {
			t.Errorf("tree.Get(%v).KeyValue() should be: (%v, %v), but got: (%v, %v)", k, k, v, key, value)
		}
	}

	{
		k, v, v2 := "bb", 22, 24
		isNew := tree.Put(k, v)
		if !isNew {
			t.Errorf("(%v, %v) should be new node, but got: %v", k, v, isNew)
		}

		key, value := tree.Get(k).KeyValue()
		if key != k || value != v {
			t.Errorf("tree.Get(%v).KeyValue() should be: (%v, %v), but got: (%v, %v)", k, k, v, key, value)
		}

		isNew2 := tree.Put(k, v2)
		if isNew2 {
			t.Errorf("(%v, %v) should be old node, but got: %v", k, v2, isNew2)
		}

		key2, value2 := tree.Get(k).KeyValue()
		if key2 != k || value2 != v2 {
			t.Errorf("tree.Get(%v).KeyValue() should be: (%v, %v), but got: (%v, %v)", k, k, v2, key, value2)
		}
	}

	fmt.Println(tree)

	{
		k, v := "cd", 34
		isNew := tree.Put(k, v)
		if !isNew {
			t.Errorf("(%v, %v) should be new node, but got: %v", k, v, isNew)
		}

		key, value := tree.Get(k).KeyValue()
		if key != k || value != v {
			t.Errorf("tree.Get(%v).KeyValue() should be: (%v, %v), but got: (%v, %v)", k, k, v, key, value)
		}
	}
	fmt.Println(tree)
}

func TestTree_Put_Slices(t *testing.T) {
	tree := NewTreeWithValid[int, []string](WithUpsert[int, []string](func(old, new []string) []string {
		return append(old, new...)
	}))
	pairs := []Pair[int, []string]{
		{key: 10, value: []string{"a"}},
		{key: 30, value: []string{"c", "_31_"}},
		{key: 40, value: []string{"d", "_41_", "_42_"}},
		{key: 50, value: []string{"e", "_51_"}},
		{key: 20, value: []string{"b"}},
		{key: 60, value: []string{}},
	}

	for _, pair := range pairs {
		tree.Put(pair.key, pair.value)
	}
	fmt.Println(tree)

	{
		k, v := 10, []string{"a"}
		key, value := tree.Get(k).KeyValue()
		if key != k || !slices.Equal(value, v) {
			t.Errorf("tree.Get(%v).KeyValue() should be: (%v, %v), but got: (%v, %v)", k, k, v, key, value)
		}

		v2 := []string{"aa", "ab", "ac"}
		tree.Put(k, v2)

		v = append(v, v2...)
		key2, value2 := tree.Get(k).KeyValue()
		if key2 != k || !slices.Equal(value2, v) {
			t.Errorf("tree.Get(%v).KeyValue() should be: (%v, %v), but got: (%v, %v)", k, k, v, key2, value2)
		}
		fmt.Println(tree)
	}

	{
		k, v := 60, []string{}
		key, value := tree.Get(k).KeyValue()
		if key != k || !slices.Equal(value, v) {
			t.Errorf("tree.Get(%v).KeyValue() should be: (%v, %v), but got: (%v, %v)", k, k, v, key, value)
		}

		v2 := []string{"fa", "fb", "fc"}
		tree.Put(k, v2)

		v = append(v, v2...)
		key2, value2 := tree.Get(k).KeyValue()
		if key2 != k || !slices.Equal(value2, v) {
			t.Errorf("tree.Get(%v).KeyValue() should be: (%v, %v), but got: (%v, %v)", k, k, v, key2, value2)
		}
		fmt.Println(tree)
	}
}

func TestTree_Get(t *testing.T) {
	tree := NewTreeWithValid[int, string]()
	pairs := []struct {
		key   int
		value string
	}{
		{key: 10, value: "a"},
		{key: 30, value: "c"},
		{key: 40, value: "d"},
		{key: 50, value: "e"},
		{key: 20, value: "b"},
		{key: 60, value: "f"},
	}
	for _, pair := range pairs {
		tree.Put(pair.key, pair.value)
	}

	{
		k, v := 10, "a"
		key, value := tree.Get(k).KeyValue()
		if key != k || value != v {
			t.Errorf("tree.Get(%v).KeyValue() should be: (%v, %v), but got: (%v, %v)", k, k, v, key, value)
		}
	}
	{
		k := 15
		get := tree.Get(k)
		if get != nil {
			t.Errorf("tree.Get(%v) should be nil, but got: %+v", k, get)
		}
	}

	{
		if get := tree.Get(math.MaxInt32); get != nil {
			t.Errorf("tree.Get(math.MaxInt32) should be nil, but got: %+v", get)
		}

		if get := tree.Get(math.MinInt32); get != nil {
			t.Errorf("tree.Get(math.MinInt32) should be nil, but got: %+v", get)
		}
	}
}

func TestTree_Update(t *testing.T) {
	tree := NewTreeWithValid[int, string]()
	pairs := []struct {
		key   int
		value string
	}{
		{key: 10, value: "a"},
		{key: 30, value: "c"},
		{key: 40, value: "d"},
		{key: 50, value: "e"},
		{key: 20, value: "b"},
		{key: 60, value: "f"},
	}
	for _, pair := range pairs {
		tree.Put(pair.key, pair.value)
	}

	{
		k, preV, curV := 10, "a", "A"
		key, value := tree.Get(k).KeyValue()
		if key != k || value != preV {
			t.Errorf("tree.Get(%v).KeyValue() should be: (%v, %v), but got: (%v, %v)", k, k, preV, key, value)
		}

		tree.Update(key, curV)
		key2, value2 := tree.Get(k).KeyValue()
		if key2 != k || value2 != curV {
			t.Errorf("tree.Get(%v).KeyValue() should be: (%v, %v), but got: (%v, %v)", k, k, curV, key, value)
		}
	}
}

func TestTree_Delete(t *testing.T) {
	tree := NewTreeWithValid[int, string]()
	{
		tree.Delete(0)
		fmt.Println(tree)
	}

	pairs := []Pair[int, string]{
		{key: 10, value: "a"},
		{key: 30, value: "c"},
		{key: 40, value: "d"},
		{key: 50, value: "e"},
		{key: 20, value: "b"},
		{key: 60, value: "f"},
		{key: 70, value: "g"},
		{key: 80, value: "h"},
		{key: 90, value: "i"},
	}
	for _, d := range pairs {
		tree.Put(d.key, d.value)
	}

	{
		k, v := 40, "d"
		key, value := tree.Get(k).KeyValue()
		if key != k || value != v {
			t.Errorf("tree.Get(%v).KeyValue() should be: (%v, %v), but got: (%v, %v)", k, k, v, key, value)
		}
		fmt.Println(tree)

		tree.Delete(k) // delete key
		if get := tree.Get(k); get != nil {
			t.Errorf("tree.Get(%v) should be nil, but got: %+v", k, get)
		}
		fmt.Println(tree)

		preSize := tree.Size()
		tree.Delete(k) // delete not exist key
		curSize := tree.Size()
		if preSize != curSize {
			t.Errorf("tree.Size() should be (%v), but got: %+v", preSize, curSize)
		}
	}
	{
		k, v := 50, "e"
		key, value := tree.Get(k).KeyValue()
		if key != k || value != v {
			t.Errorf("tree.Get(%v).KeyValue() should be: (%v, %v), but got: (%v, %v)", k, k, v, key, value)
		}

		tree.Delete(k) // delete key 20
		if get := tree.Get(k); get != nil {
			t.Errorf("tree.Get(%v) should be nil, but got: %+v", k, get)
		}
		fmt.Println(tree)

		preSize := tree.Size()
		tree.Delete(k) // delete not exist key
		curSize := tree.Size()
		if preSize != curSize {
			t.Errorf("tree.Size() should be (%v), but got: %+v", preSize, curSize)
		}
	}
	{
		k, v := 10, "a"
		key, value := tree.Get(k).KeyValue()
		if key != k || value != v {
			t.Errorf("tree.Get(%v).KeyValue() should be: (%v, %v), but got: (%v, %v)", k, k, v, key, value)
		}

		tree.Delete(k) // delete key 20
		if get := tree.Get(k); get != nil {
			t.Errorf("tree.Get(%v) should be nil, but got: %+v", k, get)
		}
		fmt.Println(tree)

		preSize := tree.Size()
		tree.Delete(k) // delete not exist key
		curSize := tree.Size()
		if preSize != curSize {
			t.Errorf("tree.Size() should be (%v), but got: %+v", preSize, curSize)
		}
	}
	{
		k, v := 30, "c"
		key, value := tree.Get(k).KeyValue()
		if key != k || value != v {
			t.Errorf("tree.Get(%v).KeyValue() should be: (%v, %v), but got: (%v, %v)", k, k, v, key, value)
		}

		tree.Delete(k) // delete key 20
		if get := tree.Get(k); get != nil {
			t.Errorf("tree.Get(%v) should be nil, but got: %+v", k, get)
		}
		fmt.Println(tree)

		preSize := tree.Size()
		tree.Delete(k) // delete not exist key
		curSize := tree.Size()
		if preSize != curSize {
			t.Errorf("tree.Size() should be (%v), but got: %+v", preSize, curSize)
		}
	}
}

func TestTree_KeysAndValues(t *testing.T) {
	tree := NewTreeWithValid[int, string]()
	{
		size := tree.Size()
		if size != 0 {
			t.Errorf("len(values) should be %v, bug got: %v", 0, size)
		}

		keys := tree.Keys()
		if len(keys) != 0 {
			t.Errorf("len(keys) should be %v, bug got: %v", 0, keys)
		}
	}

	pairs := []Pair[int, string]{
		{key: 10, value: "a"},
		{key: 30, value: "c"},
		{key: 40, value: "d"},
		{key: 50, value: "e"},
		{key: 20, value: "b"},
		{key: 60, value: "f"},
	}

	for _, pair := range pairs {
		tree.Put(pair.key, pair.value)
	}

	slices.SortFunc(pairs, func(a, b Pair[int, string]) int {
		return a.key - b.key
	})

	{
		keys := tree.Keys()
		equal := slices.EqualFunc(pairs, keys, func(elem Pair[int, string], i int) bool {
			return elem.key == i
		})
		if !equal {
			t.Errorf("tree.Keys() should be %v, bug got: %v", pairs, keys)
		}
	}
	{
		values := tree.Values()
		equal := slices.EqualFunc(pairs, values, func(e Pair[int, string], s string) bool {
			return e.value == s
		})
		if !equal {
			t.Errorf("tree.Values() should be %v, bug got: %v", pairs, values)
		}
	}

	fmt.Println(tree)
}

func TestTree_FloorAndCeil(t *testing.T) {
	tree := NewTreeWithValid[int, string]()
	pairs := []Pair[int, string]{
		{key: 10, value: "a"},
		{key: 30, value: "c"},
		{key: 40, value: "d"},
		{key: 50, value: "e"},
		{key: 20, value: "b"},
		{key: 60, value: "f"},
	}
	for _, pair := range pairs {
		tree.Put(pair.key, pair.value)
	}

	{
		k := 9 // k is less than min key
		floor := tree.Floor(k)
		if floor != nil {
			t.Errorf("tree.Floor(%v) should be nil, but got: %+v", k, floor)
		}
	}
	{
		k := 66 // k is greater than max key
		ceil := tree.Ceil(k)
		if ceil != nil {
			t.Errorf("tree.Ceil(%v) should be nil, but got: %+v", k, ceil)
		}
	}

	{
		keys := tree.Keys()
		for _, key := range keys {
			get := tree.Get(key)

			if floor := tree.Floor(key); floor != get {
				t.Errorf("tree.Floor(%v) should be tree.Get(%v), but got: %+v", key, get, floor)
			}
			if floor := tree.Floor(key + 3); floor != get {
				t.Errorf("tree.Floor(%v) should be tree.Get(%v), but got: %+v", key, get, floor)
			}

			if ceil := tree.Ceil(key); ceil != get {
				t.Errorf("tree.Floor(%v) should be tree.Get(%v), but got: %+v", key, get, ceil)
			}
			if ceil := tree.Ceil(key - 3); ceil != get {
				t.Errorf("tree.Floor(%v) should be tree.Get(%v), but got: %+v", key, get, ceil)
			}
		}
	}
}

func TestTree_Between(t *testing.T) {
	tree := NewTreeWithValid[int, string]()
	pairs := []Pair[int, string]{
		{key: 10, value: "a"},
		{key: 30, value: "c"},
		{key: 40, value: "d"},
		{key: 50, value: "e"},
		{key: 20, value: "b"},
		{key: 60, value: "f"},
	}
	for _, pair := range pairs {
		tree.Put(pair.key, pair.value)
	}

	{
		keys := tree.Keys()
		values := tree.Values()

		var keys2 []int
		var values2 []string

		left, right := 20, 50
		tree.Between(left, right, func(o *Node[int, string]) {
			k, v := o.KeyValue()
			keys2 = append(keys2, k)
			values2 = append(values2, v)
		})

		i, _ := slices.BinarySearch(keys, left+1)
		j, _ := slices.BinarySearch(keys, right)
		if !slices.Equal(keys[i-1:j+1], keys2) {
			t.Errorf("tree.Between(%v, %v, f) should be: %v, but got: %v", left, right, keys[i-1:j+1], keys2)
		}
		if !slices.Equal(values[i-1:j+1], values2) {
			t.Errorf("tree.Between(%v, %v, f) should be: %v, but got: %v", left, right, values[i-1:j+1], values2)
		}
	}
}
