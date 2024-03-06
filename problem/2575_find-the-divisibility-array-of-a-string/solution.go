package main

func divisibilityArray(word string, m int) []int {
	n := len(word)
	ans := make([]int, n)

	x := 0
	for i, b := range word {
		x = (x*10 + int(b-'0')) % m
		if x == 0 {
			ans[i] = 1
		}
	}

	return ans
}
