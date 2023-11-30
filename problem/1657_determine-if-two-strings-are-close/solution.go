package main

import (
	"slices"
	"sort"
)

func closeStrings(word1 string, word2 string) bool {
	n1, n2 := len(word1), len(word2)
	if n1 != n2 {
		return false
	}

	const N = 26
	var cnt1, cnt2 = make([]int, N), make([]int, N)
	for i := 0; i < n1; i++ {
		cnt1[word1[i]-'a']++
		cnt2[word2[i]-'a']++
	}
	for i := 0; i < N; i++ {
		if (cnt1[i] == 0) != (cnt2[i] == 0) {
			return false
		}
	}

	sort.Ints(cnt1)
	sort.Ints(cnt2)
	return slices.Equal(cnt1, cnt2)
}
