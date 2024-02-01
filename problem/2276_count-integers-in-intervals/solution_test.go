package main

import (
	"cmp"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/niluan304/leetcode/tests"
)

func Test_main(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		Constructor,
		Constructor2,
		Constructor3,
		Constructor4,
	}

	for _, f := range fs {
		err := tests.RunClass(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
["CountIntervals", "add", "add", "count", "add", "count"]
[[], [2, 3], [7, 10], [], [5, 8], []]
[null, null, null, 6, null, 8]

["CountIntervals","count","add","add","count","count","add"]
[[],[],[39,44],[13,49],[],[],[47,50]]
[null,0,null,null,37,37,null]

["CountIntervals","count","add","add","add","add","add","count","add","add"]
[[],[],[8,43],[13,16],[26,33],[28,36],[29,37],[],[34,46],[10,23]]
[null,0,null,null,null,null,null,36,null,null]

["CountIntervals","add","add","add","count","count","count","add","add","add","count","add","count","add","count","count","count","count","count","add","add","add","add","count","add","count","add","count","count","add","count","count","add","count","count","count","add","add","add","count","add","add","add","add","count","add","count","count","add","add","add","add","add","count","count","add","add","count","count","count","add","add","add","add","count","add","add","add","count","count","count","count","add","count","count","add","add","count","add","count","count","add","add","count","add","add","add","count","add","add","add","add","count","count","count","add","count","add","add","count","add","add","count","add","add","add","count","add","add","add","count","count","count","add","count","add","add","add","count","add","count","count","count","add","count","count","count","add","count","count","add","count","add"]
[[],[365,897],[261,627],[781,884],[],[],[],[328,495],[224,925],[228,464],[],[416,451],[],[747,749],[],[],[],[],[],[740,757],[51,552],[20,896],[459,712],[],[383,670],[],[701,924],[],[],[392,591],[],[],[935,994],[],[],[],[398,525],[335,881],[243,517],[],[995,1000],[15,335],[430,490],[376,681],[],[733,841],[],[],[603,633],[974,978],[466,786],[241,753],[259,887],[],[],[410,514],[173,300],[],[],[],[805,957],[272,805],[723,858],[113,118],[],[426,987],[318,997],[741,978],[],[],[],[],[701,992],[],[],[562,766],[987,1000],[],[929,929],[],[],[926,931],[913,975],[],[930,962],[707,914],[688,757],[],[430,433],[452,683],[794,919],[799,991],[],[],[],[658,731],[],[328,686],[998,999],[],[455,938],[981,988],[],[92,699],[311,690],[916,920],[],[213,339],[605,961],[719,902],[],[],[],[129,833],[],[844,926],[940,956],[148,182],[],[163,885],[],[],[],[532,886],[],[],[],[306,906],[],[],[948,963],[],[116,853]]
[null,null,null,null,637,637,637,null,null,null,702,null,702,null,702,702,702,702,702,null,null,null,null,906,null,906,null,906,906,null,906,906,null,966,966,966,null,null,null,966,null,null,null,null,977,null,977,977,null,null,null,null,null,977,977,null,null,977,977,977,null,null,null,null,986,null,null,null,986,986,986,986,null,986,986,null,null,986,null,986,986,null,null,986,null,null,null,986,null,null,null,null,986,986,986,null,986,null,null,986,null,null,986,null,null,null,986,null,null,null,986,986,986,null,986,null,null,null,986,null,986,986,986,null,986,986,986,null,986,986,null,986,null]


`

func TestLazySegmentTree_unit(t *testing.T) {
	const N = 1e9
	y := &LazySegmentTree{left: 1, right: N}
	for i := 0; i < 1e5; i++ {
		//value := rand.Intn(10) + 1
		left, right := lr(N)

		if r := rand.Intn(1000); r > 20 {
			y.Add(left, right, 1)
		} else {
			y.Query(1, N)
		}
	}
}

func TestChthollyTree_unit(t *testing.T) {
	const N = 1e9
	y := Constructor3()
	for i := 0; i < 1e5; i++ {
		//value := rand.Intn(10) + 1
		left, right := lr(N)

		if r := rand.Intn(1000); r > 20 {
			y.Add(left, right)
		} else {
			y.Count()
		}
	}
}

var w io.Writer = os.Stdout

func TestLazySegmentTree_check(t *testing.T) {
	const N = 10000
	nums := make([]int, N)
	//nums = rand.Perm(N)

	w = io.Discard

	x, y := Nums(nums), &LazySegmentTree{left: 1, right: N}

	now := time.Now()

	for i := 0; i < 10000; i++ {
		value := rand.Intn(10) + 1
		left, right := lr(N)

		switch rand.Intn(3) {
		case 0:
			x.Update(left, right, value)
			fmt.Fprintln(w, "Update", left, right, value, "nums", printInts([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
			fmt.Fprintln(w, "Update", left, right, value, "nums", printInts(x))
			y.Update(left, right, value)
		case 1:
			x.Add(left, right, value)
			fmt.Fprintln(w, "Add", left, right, value, "nums", printInts([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
			fmt.Fprintln(w, "Add", left, right, value, "nums", printInts(x))
			y.Add(left, right, value)
		case 2:
			left, right = lr(N)
			q1, q2 := x.Query(left, right), y.Query(left, right)
			//fmt.Fprintln(w,"Query", left, right, "q1:", q1, "q2:", q2)

			if q1 != q2 {
				fmt.Fprintln(os.Stdout, "Query", left, right, "q1:", q1, "q2:", q2)
				//bfs(y)
				t.Error("Query Not Equal")
			}
		}
		//bfs(y)
		fmt.Fprintln(w)

		if i > 0 && i%1000 == 0 {
			fmt.Fprintln(os.Stdout, i, time.Since(now))
		}
	}
}

func TestLazySegmentTree_fix(t *testing.T) {
	const N = 10
	nums := make([]int, N)
	//nums = rand.Perm(N)
	w = io.Discard

	x, y := Nums(nums), &LazySegmentTree{left: 1, right: N}

	x.Add(2, 4, 10)
	fmt.Fprintln(w, "x.Add(2, 4, 10)")
	fmt.Fprintln(w, "", printInts([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}), "\n", printInts(x))
	y.Add(2, 4, 10)
	bfs(y)
	x.Update(1, 9, 9)
	fmt.Fprintln(w, "x.Update(1, 9, 9)")
	fmt.Fprintln(w, "", printInts([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}), "\n", printInts(x))
	y.Update(1, 9, 9)
	bfs(y)
	x.Update(5, 9, 1)
	fmt.Fprintln(w, "x.Update(5, 9, 1)")
	fmt.Fprintln(w, "", printInts([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}), "\n", printInts(x))
	y.Update(5, 9, 1)
	bfs(y)
	x.Add(7, 10, 8)
	fmt.Fprintln(w, "x.Add(7, 10, 8)")
	fmt.Fprintln(w, "", printInts([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}), "\n", printInts(x))
	y.Add(7, 10, 8)
	bfs(y)
	x.Add(5, 6, 4)
	fmt.Fprintln(w, "x.Add(5, 6, 4)")
	fmt.Fprintln(w, "", printInts([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}), "\n", printInts(x))
	y.Add(5, 6, 4)
	bfs(y)
	x.Add(6, 10, 1)
	fmt.Fprintln(w, "x.Add(6, 10, 1)")
	fmt.Fprintln(w, "", printInts([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}), "\n", printInts(x))
	y.Add(6, 10, 1)
	bfs(y)
	x.Add(6, 10, 3)
	fmt.Fprintln(w, "x.Add(6, 10, 3)")
	fmt.Fprintln(w, "", printInts([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}), "\n", printInts(x))
	y.Add(6, 10, 3)
	bfs(y)
	x.Add(3, 5, 2)
	fmt.Fprintln(w, "x.Add(3, 5, 2)")
	fmt.Fprintln(w, "", printInts([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}), "\n", printInts(x))
	y.Add(3, 5, 2)
	bfs(y)
	x.Add(5, 8, 6)
	fmt.Fprintln(w, "x.Add(5, 8, 6)")
	fmt.Fprintln(w, "", printInts([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}), "\n", printInts(x))
	y.Add(5, 8, 6)
	bfs(y)
	x.Add(4, 8, 6)
	fmt.Fprintln(w, "x.Add(4, 8, 6)")
	fmt.Fprintln(w, "", printInts([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}), "\n", printInts(x))
	y.Add(4, 8, 6)
	bfs(y)
	x.Add(3, 10, 7)
	fmt.Fprintln(w, "x.Add(3, 10, 7)")
	fmt.Fprintln(w, "", printInts([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}), "\n", printInts(x))
	y.Add(3, 10, 7)
	bfs(y)
	x.Update(4, 8, 4)
	fmt.Fprintln(w, "x.Update(4, 8, 4)")
	fmt.Fprintln(w, "", printInts([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}), "\n", printInts(x))
	y.Update(4, 8, 4)
	bfs(y)
	x.Update(3, 6, 9)
	fmt.Fprintln(w, "x.Update(3, 6, 9)")
	fmt.Fprintln(w, "", printInts([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}), "\n", printInts(x))
	y.Update(3, 6, 9)
	bfs(y)
	x.Add(5, 8, 7)
	fmt.Fprintln(w, "x.Add(5, 8, 7)")
	fmt.Fprintln(w, "", printInts([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}), "\n", printInts(x))
	y.Add(5, 8, 7)
	bfs(y)
	x.Add(2, 5, 1)
	fmt.Fprintln(w, "x.Add(2, 5, 1)")
	fmt.Fprintln(w, "", printInts([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}), "\n", printInts(x))
	y.Add(2, 5, 1)
	bfs(y)
	x.Add(1, 10, 6)
	fmt.Fprintln(w, "x.Add(1, 10, 6)")
	fmt.Fprintln(w, "", printInts([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}), "\n", printInts(x))
	y.Add(1, 10, 6)
	bfs(y)
	x.Add(1, 4, 3)
	fmt.Fprintln(w, "x.Add(1, 4, 3)")
	fmt.Fprintln(w, "", printInts([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}), "\n", printInts(x))
	y.Add(1, 4, 3)
	bfs(y)

	if q1 := y.Query(1, 3); q1 != 56 {
		t.Error(q1)
	}
}

func printInts(x []int) string {
	s := ""
	for _, v := range x {
		s += fmt.Sprintf("%2d, ", v)
	}
	return s
}

func bfs(st *LazySegmentTree) {
	queen := []*LazySegmentTree{st}
	for len(queen) > 0 {
		var q []*LazySegmentTree
		for _, tree := range queen {
			if tree == nil || tree.left == 0 {
				continue
			}
			fmt.Printf("[%2d,%2d] %2d\t\t", tree.left, tree.right, tree.sum)
		}
		fmt.Fprintln(w)
		for _, tree := range queen {
			if tree == nil || tree.left == 0 {
				continue
			}
			fmt.Printf("todo:%2d,%2s\t\t", tree.todoAdd, toStr(tree.todoUpdate))
		}

		for _, tree := range queen {
			if tree == nil || tree.left == 0 {
				continue
			}
			q = append(q, tree.lChild, tree.rChild)
		}

		fmt.Fprintln(w)
		queen = q
	}
}

func toStr(v *int) string {
	if v == nil {
		return "_"
	}
	return strconv.Itoa(*v)
}

func lr(n int) (int, int) {
	x, y := rand.Intn(n)+1, rand.Intn(n)+1
	return min(x, y), max(x, y)
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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
