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
