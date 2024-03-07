package segtree

import (
	"reflect"
	"testing"
)

func TestNewSumLazySegmentTree(t *testing.T) {
	value := func(tree *LazySegmentTree) []int {
		l, r := tree.left, tree.right

		value := make([]int, r-l+1)
		for i := l; i <= r; i++ {
			value[i-l] = tree.Query(i, i)
		}
		return value
	}

	// create a new LazySegmentTree
	left, right := 0, 9
	want := make([]int, right-left+1)
	tree := NewSumLazySegmentTree(left, right)
	if got := value(tree); !reflect.DeepEqual(got, want) {
		t.Errorf("want: %v, but value: %v", want, got)
	}

	// 区间增加
	tree.Add(0, 4, 2)
	// index:   [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
	want = []int{2, 2, 2, 2, 2, 0, 0, 0, 0, 0}
	if got := value(tree); !reflect.DeepEqual(got, want) {
		t.Errorf("want: %v, but value: %v", want, got)
	}

	// 区间覆盖
	tree.Update(3, 7, 5)
	// index:   [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
	want = []int{2, 2, 2, 5, 5, 5, 5, 5, 0, 0}
	if got := value(tree); !reflect.DeepEqual(got, want) {
		t.Errorf("want: %v, but value: %v", want, got)
	}

	tree.Update(2, 3, 1)
	// index:   [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
	want = []int{2, 2, 1, 1, 5, 5, 5, 5, 0, 0}
	if got := value(tree); !reflect.DeepEqual(got, want) {
		t.Errorf("want: %v, but value: %v", want, got)
	}

	tree.Add(1, 2, 2)
	// index:   [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
	want = []int{2, 4, 3, 1, 5, 5, 5, 5, 0, 0}
	if got := value(tree); !reflect.DeepEqual(got, want) {
		t.Errorf("want: %v, but value: %v", want, got)
	}
}

func TestNewSumLazySegmentTreeWithNums(t *testing.T) {
	value := func(tree *LazySegmentTree) []int {
		l, r := tree.left, tree.right

		value := make([]int, r-l+1)
		for i := l; i <= r; i++ {
			value[i-l] = tree.Query(i, i)
		}
		return value
	}

	// create a new LazySegmentTree
	// 	index:   [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	tree := NewSumLazySegmentTreeWithNums(want)
	if got := value(tree); !reflect.DeepEqual(got, want) {
		t.Errorf("want: %v, but value: %v", want, got)
	}

	// 区间增加
	tree.Add(1, 4, 2)
	// index:   [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
	want = []int{3, 4, 5, 6, 5, 6, 7, 8, 9, 10}
	if got := value(tree); !reflect.DeepEqual(got, want) {
		t.Errorf("want: %v, but value: %v", want, got)
	}

	// 区间覆盖
	tree.Update(3, 7, 5)
	// index:   [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
	want = []int{3, 4, 5, 5, 5, 5, 5, 8, 9, 10}
	if got := value(tree); !reflect.DeepEqual(got, want) {
		t.Errorf("want: %v, but value: %v", want, got)
	}

	tree.Update(2, 3, 1)
	// index:   [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
	want = []int{3, 1, 1, 5, 5, 5, 5, 8, 9, 10}
	if got := value(tree); !reflect.DeepEqual(got, want) {
		t.Errorf("want: %v, but value: %v", want, got)
	}

	tree.Add(1, 2, 2)
	// index:   [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
	want = []int{5, 3, 1, 5, 5, 5, 5, 8, 9, 10}
	if got := value(tree); !reflect.DeepEqual(got, want) {
		t.Errorf("want: %v, but value: %v", want, got)
	}
}
