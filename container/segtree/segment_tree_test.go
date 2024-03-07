package segtree

import (
	"math"
	"reflect"
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func TestMid(t *testing.T) {
	fs := []func(left, right int) int{
		func(left, right int) int { return (left + right) / 2 },         // 无法处理溢出 和 下取整
		func(left, right int) int { return int(uint(left+right) >> 1) }, // 无法处理 (-99 + 50)/2
		func(left, right int) int { return (left + right) >> 1 },        // 无法处理 (50 + math.MaxInt)/2
		func(left, right int) int { return left + (right-left)/2 },      // 无法处理 (-99 + math.MaxInt)/2
	}

	args := []struct{ left, right, want int }{
		{left: -99, right: 50, want: -25},
		{left: -99, right: math.MaxInt, want: (-99 + math.MaxInt) / 2},
		{left: 50, right: math.MaxInt, want: 50 + (math.MaxInt-50)/2},
	}

	for _, f := range fs {
		name := tests.FuncName(f)
		t.Run(name, func(t *testing.T) {
			for _, arg := range args {
				left, right := arg.left, arg.right
				if got := f(left, right); got != arg.want {
					t.Errorf("%v(%v, %v) = %v, want %v", name, left, right, got, arg.want)
				}
			}
		})
	}
}

func TestNewSumSegmentTreeWithNums(t *testing.T) {
	value := func(tree *SegmentTree) []int {
		l, r := tree.left, tree.right

		value := make([]int, r-l+1)
		for i := l; i <= r; i++ {
			value[i-l] = tree.Query(i, i)
		}
		return value
	}

	// create a new LazySegmentTree

	var want []int
	// index:   [  1,  2,  3,   4,  5,   6,   7,   8,   9, 10]
	want = []int{-28, -39, 53, 65, 11, -56, -65, -39, -43, 97}
	tree := NewSumSegmentTreeWithNums(want)
	if got := value(tree); !reflect.DeepEqual(got, want) {
		t.Errorf("want: %v, but value: %v", want, got)
	}

	if got, want := tree.Query(5+1, 6+1), -121; !reflect.DeepEqual(got, want) {
		t.Errorf("want: %v, but value: %v", want, got)
	}

	tree.Set(10, func(old int) int { return 27 })
	// index:   [  1,  2,  3,   4,  5,   6,   7,   8,   9, 10]
	want = []int{-28, -39, 53, 65, 11, -56, -65, -39, -43, 27}
	if got := value(tree); !reflect.DeepEqual(got, want) {
		t.Errorf("want: %v, but value: %v", want, got)
	}

	if got, want := tree.Query(2+1, 3+1), 118; !reflect.DeepEqual(got, want) {
		t.Errorf("want: %v, but value: %v", want, got)
	}

	if got, want := tree.Query(6+1, 7+1), -104; !reflect.DeepEqual(got, want) {
		t.Errorf("want: %v, but value: %v", want, got)
	}

	tree.Set(5, func(old int) int { return 6 })
	// index:   [  1,  2,  3,   4, 5,   6,   7,   8,   9, 10]
	want = []int{-28, -39, 53, 65, 6, -56, -65, -39, -43, 27}
	if got := value(tree); !reflect.DeepEqual(got, want) {
		t.Errorf("want: %v, but value: %v", want, got)
	}
}

func TestNewSumSegmentTree(t *testing.T) {
	value := func(tree *SegmentTree) []int {
		l, r := tree.left, tree.right

		value := make([]int, r-l+1)
		for i := l; i <= r; i++ {
			value[i-l] = tree.Query(i, i)
		}
		return value
	}

	// create a new LazySegmentTree
	tree := NewSumSegmentTree(0, 9)
	want := make([]int, tree.right-tree.left+1) // want := make([]int, 10)
	if got := value(tree); !reflect.DeepEqual(got, want) {
		t.Errorf("want: %v, but value: %v", want, got)
	}

	// set want[4], want[8] = 2, 5
	tree.Set(4, func(old int) int { return 2 })
	tree.Set(8, func(old int) int { return 5 })
	// index:   [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
	want = []int{0, 0, 0, 0, 2, 0, 0, 0, 5, 0}
	if got := value(tree); !reflect.DeepEqual(got, want) {
		t.Errorf("want: %v, but value: %v", want, got)
	}

	// set want[4] += 6
	tree.Set(4, func(old int) int { return old + 6 })
	// index:   [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
	want = []int{0, 0, 0, 0, 8, 0, 0, 0, 5, 0}
	if got := value(tree); !reflect.DeepEqual(got, want) {
		t.Errorf("want: %v, but value: %v", want, got)
	}

	// set want[8] /= 2
	tree.Set(8, func(old int) int { return old / 2 })
	// index:   [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
	want = []int{0, 0, 0, 0, 8, 0, 0, 0, 2, 0}
	if got := value(tree); !reflect.DeepEqual(got, want) {
		t.Errorf("want: %v, but value: %v", want, got)
	}

	// set want[8] = want[8]^2
	tree.Set(8, func(old int) int { return old * old })
	// index:   [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
	want = []int{0, 0, 0, 0, 8, 0, 0, 0, 4, 0}
	if got := value(tree); !reflect.DeepEqual(got, want) {
		t.Errorf("want: %v, but value: %v", want, got)
	}
}
