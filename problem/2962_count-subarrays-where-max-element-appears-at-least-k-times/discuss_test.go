package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

func Test_checkDiscuss(t *testing.T) {
	fs := []func(nums []int, k int) int64{
		discuss,
		discuss2,
		discuss3,
	}

	for i := 0; i < 10000; i++ {
		err := check(fs...)
		if err != nil {
			t.Error(err)
		}
	}
}

func check(fs ...func(nums []int, k int) int64) error {
	var n = rand.Intn(10000) + 1 // rand.Intn(n) 随机数组范围为 [0, n-1]
	var k = rand.Intn(5) + 1
	var nums = make([]int, n)
	for j, _ := range nums {
		nums[j] = rand.Intn(50) + 1
	}

	except := violent(nums, k)
	for _, f := range fs {
		actual := f(nums, k)
		if actual != except {
			return fmt.Errorf("f:%s, except: %d, actual:%d, k: %d, nums:%v", funcName(f), except, actual, k, nums)
		}
	}

	return nil
}

func funcName(f any) string {
	name := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	i := strings.LastIndex(name, ".")
	return name[i+1:]
}
