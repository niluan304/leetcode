package main

func equalSubstring(s string, t string, maxCost int) int {
	var (
		n        = len(s)
		costList = make([]int, 0, n)

		ans  = 0
		left = 0
		cost = 0
	)
	for i := 0; i < n; i++ {
		costList = append(costList, _abs(int(s[i]), int(t[i])))
	}

	for right, b := range costList {
		cost += b

		// if和for 在这里有区别吗?
		for cost > maxCost {
			cost -= costList[left]
			left++
		}

		ans = max(ans, right-left+1)
	}

	return ans
}

func _abs(x, y int) int {
	c := x - y
	if c > 0 {
		return c
	}
	return -c
}

func equalSubstring2(s string, t string, maxCost int) int {
	var (
		n        = len(s)
		costList = make([]int, 0, n)

		left = 0
		cost = 0
	)
	for i := 0; i < n; i++ {
		costList = append(costList, _abs(int(s[i]), int(t[i])))
	}

	for _, b := range costList {
		cost += b

		// if和for 在这里有区别吗?
		if cost > maxCost {
			cost -= costList[left]
			left++
		}
	}

	return n - left
}
