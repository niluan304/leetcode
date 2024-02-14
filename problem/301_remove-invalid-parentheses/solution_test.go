package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_removeInvalidParentheses(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(s string) []string{
		//bruteForce,
		removeInvalidParentheses,
		//removeInvalidParentheses2,
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
"()())()"
["(())()","()()()"]

"(a)())()"
["(a())()","(a)()()"]

")("
[""]

"(a)(b)c()"
["(a)(b)c()"]

"("
[""]

"(()"
["()"]

")(f"
["f"]

`
