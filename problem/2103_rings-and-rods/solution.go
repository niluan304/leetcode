package main

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func countPoints(rings string) int {
	var n = len(rings)
	var list [10]int
	for i := 0; i < n; i += 2 {
		idx := rings[i+1] - '0'
		list[idx] |= 1 << (rings[i] - 'A')
	}

	const RBG = 1<<('R'-'A') | 1<<('G'-'A') | 1<<('B'-'A')
	var ans = 0
	for _, v := range list {
		if v == RBG {
			ans++
		}
	}
	return ans
}
