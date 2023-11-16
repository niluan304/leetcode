package main

import (
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
)

func Test_earliest_possible_day_of_full_bloom(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		earliestFullBloom,
	}

	for _, f := range fs {
		name := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
		i := strings.LastIndex(name, ".")

		t.Run(name[i+1:], func(t *testing.T) {
			err := testutil.RunLeetCodeFuncWithFile(t, f, "sample.txt", targetCaseNum)
			if err != nil {
				t.Error(err)
			}
		})
	}
}
