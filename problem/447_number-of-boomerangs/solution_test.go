package main

import (
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
)

func Test_main(t *testing.T) {
	targetCaseNum := 0 // -1
	inputs, outputs := sample(samples)

	fs := []func(points [][]int) int{
		numberOfBoomerangs,
		numberOfBoomerangs2,
	}

	for _, f := range fs {
		name := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
		i := strings.LastIndex(name, ".")

		t.Run(name[i+1:], func(t *testing.T) {
			err := testutil.RunLeetCodeFuncWithCase(t, f, inputs, outputs, targetCaseNum)
			if err != nil {
				t.Error(err)
			}
		})
	}
}

var samples = `
[[0,0],[1,0],[2,0]]
2

[[1,1],[2,2],[3,3]]
2

[[1,1]]
0


`

func sample(samples string) (inputs, outputs [][]string) {
	for _, s := range strings.Split(samples, "\n\n") {
		s = strings.TrimSpace(s)
		if s == "" {
			continue
		}

		raw := strings.Split(s, "\n")
		n := len(raw)
		inputs = append(inputs, raw[:n-1])
		outputs = append(outputs, []string{raw[n-1]})
	}
	return
}
