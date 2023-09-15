package structs

func Sum(list []int) int {
	var sum = 0
	for _, n := range list {
		sum += n
	}

	return sum
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
