package main

func firstUniqChar(s string) int {
	var m = map[byte]int{}
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}

	for i := 0; i < len(s); i++ {
		if m[s[i]] == 1 {
			return i
		}
	}

	return -1
}

func firstUniqChar2(s string) int {
	var arr = [26]int{}
	for i := 0; i < len(s); i++ {
		arr[s[i]-'a']++
	}

	for i := 0; i < len(s); i++ {
		if arr[s[i]-'a'] == 1 {
			return i
		}
	}

	return -1
}

func firstUniqChar4(s string) int {
	var arr = [26]int{}

	for i := range s {
		arr[s[i]-'a']++
	}

	for i := range s {
		if arr[s[i]-'a'] == 1 {
			return i
		}
	}

	return -1
}

func firstUniqChar3(s string) int {
	var arr = [26]int{}

	for _, ch := range s {
		arr[ch-'a']++
	}

	for i, ch := range s {
		if arr[ch-'a'] == 1 {
			return i
		}
	}

	return -1
}

type pair struct {
	ch  byte
	pos int
}

func leetcode3(s string) int {
	n := len(s)
	pos := [26]int{}
	for i := range pos[:] {
		pos[i] = n
	}
	var q []pair
	for i := range s {
		ch := s[i] - 'a'
		if pos[ch] == n {
			pos[ch] = i
			q = append(q, pair{ch, i})
		} else {
			pos[ch] = n + 1
			for len(q) > 0 && pos[q[0].ch] == n+1 {
				q = q[1:]
			}
		}
	}
	if len(q) > 0 {
		return q[0].pos
	}
	return -1
}
