package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_countDigitOne(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(n int) int{
		countDigitOne,
		countDigitOne2,
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
13
6

0
0

20
12

29
13

31
14

34
14

45
15

56
16

77
18

89
19

100
21

3432423
3123793

123214
93262

4546
2415

78798912
65299685

1657765
1691123

123673423
130913517


`
