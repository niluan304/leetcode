package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"testing"

	"github.com/sourcegraph/conc"

	"github.com/niluan304/leetcode/tests"
)

type Func func(nums []int, k int) int64

func Test_checkDiscuss(t *testing.T) {
	fs := []Func{
		discuss,
		discuss2,
		discuss3,
		discuss4,
	}

	for i := 0; i < 1000; i++ {
		err := validate(fs...)
		if err != nil {
			t.Error(err)
		}
	}
}

func validate(fs ...Func) (err error) {
	var n = rand.Intn(1000000) + 1 // rand.Intn(n) 随机数组范围为 [0, n-1]
	var k = rand.Intn(5) + 1
	var nums = make([]int, n)
	for j, _ := range nums {
		nums[j] = rand.Intn(50) + 1
	}

	wg := conc.NewWaitGroup()
	defer wg.Wait()

	except := bruteForce(nums, k)
	for _, f := range fs {
		wg.Go(func() {
			actual := f(nums, k)
			if actual != except {
				e := fmt.Errorf("f:%s, except: %d, actual:%d, k: %d, nums:%v", tests.FuncName(f), except, actual, k, nums)
				err = errors.Join(err, e)
			}
		})
	}

	return err
}

func bruteForce(nums []int, k int) int64 {
	var n, ans = len(nums), 0
	for i := 0; i < n; i++ {
		var count, mx = map[int]int{}, math.MinInt
		for j := i; j < n; j++ {
			count[nums[j]]++
			mx = max(mx, nums[j])
			if count[mx] >= k {
				ans++
			}
		}
	}
	return int64(ans)
}
