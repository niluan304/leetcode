package main

func findAnagrams(s string, p string) []int {
	var (
		ans []int
		n1  = len(s)
		n2  = len(p)
	)
	if n1 < n2 {
		return nil
	}

	var (
		count = [26]int{}
	)
	for i := 0; i < n2; i++ {
		count[p[i]-'a']++
	}

	for i := 0; i < n1-n2+1; i++ {
		var compare = [26]int{}
		for j := 0; j < n2; j++ {
			compare[s[i+j]-'a']++
		}

		if compare == count {
			ans = append(ans, i)
		}
	}

	return ans
}

func findAnagrams2(s string, p string) []int {
	if len(p) > len(s) {
		return nil
	}

	var (
		ans []int

		total = [26]int{}
		count = [26]int{}
		n     = len(p)
		left  = 0
	)

	// init
	for i := 0; i < n; i++ {
		total[p[i]-'a']++
		count[s[i]-'a']++
	}

	// begin with left = 0, 等同于 leetcode1的写法
	count[s[n-1]-'a']--
	for right := n - 1; right < len(s); right++ {
		count[s[right]-'a']++

		if count == total {
			ans = append(ans, left)
		}
		count[s[left]-'a']--
		left++
	}

	return ans
}
