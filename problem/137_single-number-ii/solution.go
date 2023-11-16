package main

func singleNumber(nums []int) int {
	a, b := 0, 0
	for _, x := range nums {
		_a := a
		a = (a ^ x) & (a | b)
		b = (b ^ x) &^ _a
	}

	return b
}

func _max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func _min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func _abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
