package main

func checkIfPangram(sentence string) bool {
	if len(sentence) < 26 {
		return false
	}

	var m = make(map[rune]struct{}, 26)
	for _, char := range sentence {
		m[char] = struct{}{}
	}

	return len(m) == 26
}

func leetcode(sentence string) bool {
	state := 0
	for _, c := range sentence {
		state |= 1 << (c - 'a')
	}
	return state == 1<<26-1
}
