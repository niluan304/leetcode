package main

// 换句话说，s1 的排列之一是 s2 的 子串 。
func checkInclusion(s1 string, s2 string) bool {
	var (
		n1 = len(s1)
		n2 = len(s2)
	)

	if n1 > n2 {
		return false
	}

	var (
		total = [26]int{}
		count = [26]int{}
		left  = 0
	)

	// init
	for i := 0; i < n1; i++ {
		total[s1[i]-'a']++
		count[s2[i]-'a']++
	}

	count[s2[n1-1]-'a']--
	for right := n1 - 1; right < n2; right++ {
		count[s2[right]-'a']++

		// 比较的花费: O(26/2)
		if count == total {
			return true
		}
		count[s2[left]-'a']--
		left++
	}

	return false
}

func checkInclusion2(s1 string, s2 string) bool {
	var (
		n1 = len(s1)
		n2 = len(s2)
	)

	if n1 > n2 {
		return false
	}

	var (
		count = [26]int{}
		left  = 0
	)

	// init
	for _, b := range s1 {
		count[b-'a']++
	}

	for right, b := range s2 {
		count[b-'a']--
		for count[b-'a'] < 0 {
			count[s2[left]-'a']++
			left++
		}

		if right-left+1 == n1 {
			return true
		}
	}

	return false
}
