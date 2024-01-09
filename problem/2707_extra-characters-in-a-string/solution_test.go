package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_minExtraChar(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(s string, dictionary []string) int{
		minExtraChar,
		minExtraChar2,
		// leetcode,
		endlessCheng,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
"leetscode"
["leet","code","leetcode"]
1

"sayhelloworld"
["hello","world"]
3

"abcabcd"
["abc","abcd"]
0

"aaaaaaaa"
["aaa","aaaa"]
0

"aaaaaaaaa"
["aaa","aaaa"]
0

"aaaabaaa"
["aaaa","aaa"]
1

"aaabaaa"
["aa","ab","aaa"]
0

"dwmodizxvvbosxxw"
["ox","lb","diz","gu","v","ksv","o","nuq","r","txhe","e","wmo","cehy","tskz","ds","kzbu"]
7


`
