package main

import (
	"math/rand"
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_beautifulSubstrings(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(s string, k int) int64{
		bruteForce,
		beautifulSubstrings,
		beautifulSubstrings2,
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
"baeyh"
2
2

"abba"
1
3

"bcdf"
1
0

"ababaababbbb"
8
2

`

func Test_validate(t *testing.T) {
	for i := 0; i < 1000000; i++ {
		s, k := generate()
		except := bruteForce(s, k)
		err := tests.Validate2(s, k, except,
			beautifulSubstrings,
			beautifulSubstrings2,
			beautifulSubstrings3,
		)
		if err != nil {
			t.Error(err)
		}
	}
}

func generate() (s string, k int) {
	const letters = "abcdefghijklmnopqrstuvwxyz"

	var str = make([]byte, rand.Intn(1000)+1)
	for i, _ := range str {
		str[i] = letters[rand.Intn(len(letters))]
	}
	return string(str), rand.Intn(len(str)) + 1
}

func bruteForce(s string, k int) (ans int64) {
	var n = len(s)

	for i := 0; i < n; i++ {
		vowels, consonants := 0, 0
		for j := i; j < n; j++ {
			switch s[j] {
			case 'a', 'e', 'i', 'o', 'u':
				vowels++
			default:
				consonants++
			}
			if (vowels == consonants) && (vowels*vowels)%k == 0 {
				ans++
			}
		}
	}
	return
}
