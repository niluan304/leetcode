package main

import (
	"math/rand"
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_discuss(t *testing.T) {
	fs := []func(num string, k int) string{
		discuss,
		bruteForce,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, data, 0)
		if err != nil {
			t.Error(err)
		}
	}
}

var data = `
"53133066909114"
2
"53331066909114"

"868112329103983686"
6
"886112329103983686"


`

func Test_validateDiscuss(t *testing.T) {
	fs := []func(num string, k int) string{
		discuss,
	}

	for i := 0; i < 10_000; i++ {
		num, k := generate()
		except := bruteForce(num, k)

		err := tests.Validate2(num, k, except, fs...)
		if err != nil {
			t.Error(err)
		}
	}
}

func generate() (numStr string, k int) {
	var n = rand.Intn(1000) + 1 // rand.Intn(n) 随机数组范围为 [0, n-1]
	var num = make([]byte, n)

	for i := 1; i < n; i++ {
		num[i] = byte(rand.Intn(10)) + '0'
	}
	num[0] = byte(rand.Intn(9)) + '1' // 不应该有前导零

	return string(num), rand.Intn((n+1)/2) + 1
}

func bruteForce(num string, k int) string {
	ans := num
	s, n := []byte(num), len(num)
	for i := 0; i < n; i++ {
		for j := max(0, i-k); j < i; j++ {
			s[i], s[j] = s[j], s[i]
			//v := string(s)
			ans = max(ans, string(s))
			s[i], s[j] = s[j], s[i]
		}
	}
	return ans
}
