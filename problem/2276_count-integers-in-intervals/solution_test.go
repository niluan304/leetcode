package main

import (
	"math/rand"
	"testing"

	. "github.com/niluan304/leetcode/container/segtree"
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
	y := NewSumLazySegmentTree(1, N)
	for i := 0; i < 1e5; i++ {
		// value := rand.Intn(10) + 1
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
		// value := rand.Intn(10) + 1
		left, right := lr(N)

		if r := rand.Intn(1000); r > 20 {
			y.Add(left, right)
		} else {
			y.Count()
		}
	}
}

func lr(n int) (int, int) {
	x, y := rand.Intn(n)+1, rand.Intn(n)+1
	return min(x, y), max(x, y)
}
