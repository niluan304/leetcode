package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_power_of_four(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		isPowerOfFour,
		isPowerOfFour2,
		isPowerOfFour3,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
16
true

5
false

1
true

4
true

64
true

256
true

1024
true

4096
true

16384
true

65536
true

262144
true

1048576
true

4194304
true

16777216
true

67108864
true

268435456
true

1073741824
true
`
