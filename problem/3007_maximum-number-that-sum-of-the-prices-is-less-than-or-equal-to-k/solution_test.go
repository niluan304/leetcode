package main

import (
	"strconv"
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_findMaximumNumber(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(k int64, x int) int64{
		findMaximumNumber,
		findMaximumNumber2,
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
9
1
6

7
2
9
` +
	`
6271836263728
1
331018054007

6271836263728
2
656450237611

6271836263728
3
975760112388

6271836263728
4
1275042756831

6271836263728
5
1636721298143

6271836263728
6
2096339854863

6271836263728
7
2422661505403

6271836263728
8
2586360068311

`

func Test_digitDp(t *testing.T) {
	type args struct {
		low, high string
		x         int
	}

	ts := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{low: "0", high: strconv.FormatInt(int64(6), 2), x: 1},
			want: 9,
		},
		{
			name: "",
			args: args{low: "0", high: strconv.FormatInt(int64(9), 2), x: 2},
			want: 6,
		},
		{
			name: "",
			args: args{low: "0", high: strconv.FormatInt(int64(10), 2), x: 2},
			want: 8,
		},
	}
	for _, tt := range ts {
		t.Run(tt.name, func(t *testing.T) {
			if got := digitDp(tt.args.low, tt.args.high, tt.args.x); got != tt.want {
				t.Errorf("digitDp(%+v) = %v, want %v", tt.args, got, tt.want)
			}
		})
	}
}
