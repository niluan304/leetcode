package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 72 ms 的代码示例
func leetcodeMinTime(fruits []int) int {
	FF := -1
	Fidx := 0
	SF := -1
	Sidx := 0
	j := 0
	i := 0
	res := 0
	for ; i < len(fruits); i++ {
		if FF == -1 || FF == fruits[i] {
			FF = fruits[i]
			Fidx = i
			res = max(res, i-j+1)
		} else if SF == -1 || SF == fruits[i] {
			SF = fruits[i]
			Sidx = i
			res = max(res, i-j+1)
		} else {
			if Fidx < Sidx {
				j = Fidx + 1
				Fidx, Sidx = Sidx, i
				FF, SF = SF, fruits[i]
			} else {
				j = Sidx + 1
				Sidx = i
				SF = fruits[i]
			}
		}

	}

	return res
}
