package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_count(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(num1 string, num2 string, minSum int, maxSum int) int{
		count,
		count2,
		count3,
		// leetcode,
		// endlessCheng,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
"1"
"12"
1
8
11

"1"
"5"
1
5
5

"1"
"1000000000000000000"
1
400
49

"1"
"9999"
4
8
460

"123"
"9999"
4
8
413

"1"
"9"
1
8
8

"10"
"12"
1
8
3

"1"
"4354353463643"
1
400
353433165

"134"
"993423499"
4
8
24036

"1"
"993423499"
4
8
24090

"234123"
"999342349"
4
8
21653

"1"
"999342349"
4
8
24090

"1"
"4354353463643"
34
324
386160468

"1"
"993423499"
23
234
975906816

"1"
"999342349"
14
78
998846883

"4179205230"
"7748704426"
8
46
883045899

`
