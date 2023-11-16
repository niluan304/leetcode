package main

func balancedString(s string) int {
	var (
		ans = 0
	)

	var set [4]int
	for _, r := range s {
		switch r {
		case 'Q':
			set[0]++
		case 'W':
			set[1]++
		case 'E':
			set[2]++
		case 'R':
			set[3]++
		}
	}
	avg := len(s) / 4
	for _, v := range set {
		if v > avg {
			ans += v - avg
		}
	}

	return ans
}

func Index(r rune) int {
	switch r {
	case 'Q':
		return 0
	case 'W':
		return 1
	case 'E':
		return 2
	case 'R':
		return 3
	}
	return -1
}

func _min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func balancedString2(s string) int {
	var (
		left  = 0
		count = len(s)
	)

	var all = map[rune]int{}
	for _, r := range s {
		all[r]++
	}

	avg := len(s) / 4
	for k := range all {
		all[k] -= avg
	}
	if all['Q'] == 0 &&
		all['W'] == 0 &&
		all['E'] == 0 &&
		all['R'] == 0 {
		return 0
	}
	var set = map[byte]int{}
	for right := range s {
		set[s[right]]++

		for set['Q'] >= all['Q'] &&
			set['W'] >= all['W'] &&
			set['E'] >= all['E'] &&
			set['R'] >= all['R'] &&
			left <= right {
			count = _min(count, right-left+1)
			set[s[left]]--
			left++
		}
	}

	return count
}

func balancedString3(s string) int {
	var (
		left  = 0
		count = len(s)
	)

	avg := len(s) / 4
	var all = [4]int{avg, avg, avg, avg}
	for _, r := range s {
		all[Index(r)]--
	}

	if all == [4]int{0, 0, 0, 0} {
		return 0
	}

	var set = [4]int{}
	for right, r := range s {
		set[Index(r)]--

		for set[0] <= all[0] && set[1] <= all[1] && set[2] <= all[2] && set[3] <= all[3] && left <= right {
			count = _min(count, right-left+1)
			set[Index(rune(s[left]))]++
			left++
		}
	}
	return count
}

func balancedString4(s string) int {
	var (
		left  = 0
		count = len(s)
	)

	var set = [4]int{}
	for _, r := range s {
		set[Index(r)]++
	}

	avg := len(s) / 4
	if set == [4]int{avg, avg, avg, avg} {
		return 0
	}

Next:
	for right, r := range s {
		set[Index(r)]--
		for {
			for _, v := range set {
				if v > avg {
					continue Next
				}
			}
			count = _min(count, right-left+1)
			set[Index(rune(s[left]))]++
			left++
		}
	}
	return count
}
