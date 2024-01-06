package main

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}

	var count = [26]int{}
	for _, b := range magazine {
		count[b-'a']++
	}

	for _, b := range ransomNote {
		count[b-'a']--
		if count[b-'a'] < 0 {
			return false
		}
	}
	return true
}
