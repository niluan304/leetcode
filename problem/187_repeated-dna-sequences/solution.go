package main

func findRepeatedDnaSequences(s string) []string {
	var m = map[string]int{}
	for i := 0; i <= len(s)-10; i++ {
		m[s[i:i+10]]++
	}

	var ans []string
	for seq, v := range m {
		if v > 1 {
			ans = append(ans, seq)
		}
	}
	return ans
}

func findRepeatedDnaSequences2(s string) []string {
	var m = map[string]int{}
	var ans []string
	for i := 0; i <= len(s)-10; i++ {
		seq := s[i : i+10]
		m[seq]++
		if m[seq] == 2 {
			ans = append(ans, seq)
		}
	}
	return ans
}
