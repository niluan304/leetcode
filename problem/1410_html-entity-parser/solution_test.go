package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_main(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(text string) string{
		entityParser,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
"&amp; is an HTML entity but &ambassador; is not."
"& is an HTML entity but &ambassador; is not."

"and I quote: &quot;...&quot;"
"and I quote: \"...\""

"asdasdasdasd"
"asdasdasdasd"

"&amp;amp;"
"&amp;"
`
