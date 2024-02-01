package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_main(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(n int) int{
		nextBeautifulNumber,
		leetcode1,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
1
22

1000
1333

3000
3133

80
122

224160
224444

1000000
1224444

0
1
`
