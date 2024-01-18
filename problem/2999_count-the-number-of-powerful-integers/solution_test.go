package main

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_numberOfPowerfulInt(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(start int64, finish int64, limit int, s string) int64{
		numberOfPowerfulInt,
		numberOfPowerfulInt2,
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
1
1000
4
"3"
25

343
60050
4
"124"
24

1
6000
4
"124"
5

15
215
6
"10"
2

1000
2000
4
"3000"
0

1000
2000
5
"1500"
1

1534
215567
6
"103"
109


200000006000
1000000000000
5
"354342"
31104

1000000
10000000
5
"3"
38880

100000000000
1000000000000
5
"354342"
38880

1
1000000000000
5
"354342"
46656

1
100000000000
5
"354342"
7776

1000345
2000234325
4
"3000"
6239

1000
2000234
4
"3"
6239

`

func TestName(t *testing.T) {
Loop:
	for i := 0; i < 1000; i++ {
		s := strconv.Itoa(i)
		if !strings.HasSuffix(s, "3") {
			continue
		}
		for _, b := range s {
			if b > '4' {
				continue Loop
			}
		}
		fmt.Println(s)
	}
}
