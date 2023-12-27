package main

import (
	"cmp"
	"fmt"
	"math/rand"
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
)

func Test_range_sum_query_mutable(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		Constructor,
		Constructor2,
		Constructor3,
	}

	for _, f := range fs {
		name := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
		i := strings.LastIndex(name, ".")

		t.Run(name[i+1:], func(t *testing.T) {
			err := testutil.RunLeetCodeClassWithFile(t, f, "sample.txt", targetCaseNum)
			if err != nil {
				t.Error(err)
			}
		})
	}
}

func TestSegmentTree_check(t *testing.T) {
	const N = 10000
	nums := make([]int, N)
	nums = rand.Perm(N)

	x, y := Nums(nums), NewSegmentTree(nums)

	for i := 0; i < 100000; i++ {
		value := rand.Intn(10) + 1
		idx := rand.Intn(N) + 1

		switch rand.Intn(3) {
		case 0:
			x.Update(idx, idx, value)
			y.Update(idx, value)
			//fmt.Println("Update", left, right, value, "nums", printInts([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
			//fmt.Println("Update", left, right, value, "nums", printInts(x))
		case 1:
			x.Add(idx, idx, value)
			y.Add(idx, value)
			//fmt.Println("Add", left, right, value, "nums", printInts([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
			//fmt.Println("Add", left, right, value, "nums", printInts(x))
		case 2:
			left, right := lr(N)
			q1, q2 := x.Query(left, right), y.Query(left, right)
			fmt.Println("Query", left, right, "q1:", q1, "q2:", q2)

			if q1 != q2 {
				t.Error("Query Not Equal")
			}
			fmt.Println()
		}
	}
}

type Nums []int

func (n *Nums) Query(left, right int) int {
	return Sum((*n)[left-1 : right])
}

func (n *Nums) Update(left, right int, value int) {
	for i := left - 1; i < right; i++ {
		(*n)[i] = value
	}
}

func (n *Nums) Add(left, right int, add int) {
	for i := left - 1; i < right; i++ {
		(*n)[i] += add
	}
}

func Sum[S ~[]E, E cmp.Ordered](x S) E {
	var m E
	for i := 0; i < len(x); i++ {
		m += x[i]
	}
	return m
}

func lr(n int) (int, int) {
	x, y := rand.Intn(n)+1, rand.Intn(n)+1
	return min(x, y), max(x, y)
}
