package main

func totalFruit(fruits []int) int {
	var (
		left = 0
		ans  = 0
		set  = map[int]int{}
	)

	for right, fruit := range fruits {
		set[fruit]++

		for len(set) > 2 {
			v := fruits[left]
			set[v]--
			if set[v] == 0 {
				delete(set, v)
			}
			left++
		}

		ans = _max(ans, right-left+1)
	}

	return ans
}

func _max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
