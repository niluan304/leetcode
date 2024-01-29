package main

import (
	"math/bits"
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_countCompleteSubstrings(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(word string, k int) int{
		bruteForce,
		countCompleteSubstrings,
		//countCompleteSubstrings2,
		//leetcode,
		//endlessCheng,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
"igigee"
2
3

"aaabbbccc"
3
6

"egiigee"
2
5

`

func bruteForce(word string, k int) int {
	var ans = 0
	var n = len(word)
	for i := 0; i < n; i++ {
		var count = map[uint8]int{}
		x := uint(0)

		for j := i; j < n; j++ {
			b := word[j] - 'a'
			if j > i {
				a := word[j-1] - 'a'
				if Abs(int(b)-int(a)) > 2 {
					break
				}
			}
			count[b]++
			if count[b] == k {
				x |= 1 << b // 二进制置为 1
			} else {
				x &= ^(1 << b) // 二进制置为 0
			}
			if len(count) == bits.OnesCount(x) {
				ans++
			}
		}
	}

	return ans
}
