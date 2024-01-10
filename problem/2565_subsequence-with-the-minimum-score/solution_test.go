package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_minimumScore(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(s string, t string) int{
		minimumScore,
		minimumScore2,
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
"abacaba"
"bzaa"
1

"cde"
"xyz"
3

"abacaba"
"baa"
0

"abacaba"
"zbaa"
1

"abecdebe"
"eaebceae"
6

"dabbbeddeabbaccecaee"
"bcbbaabdbebecbebded"
16

"dcadebdecbeaedd"
"dcdadeb"
1

"abacaba"
"baax"
1

"gbjbacdiiiecgceeafdcdhjhhcjfchjbejibhejgjhhhjhifahfbbcfcajehjgcjgffjdejbhjahgffgdaifhhgaadjiabfdf"
"hidefgbgjghceagdaaib"
5


`
