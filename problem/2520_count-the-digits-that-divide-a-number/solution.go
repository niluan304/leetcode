package main

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

func countDigits(num int) (ans int) {
	for x := num; x > 0; x /= 10 {
		if num%(x%10) == 0 {
			ans++
		}
	}
	return
}

func countDigits2(num int) int {
	var ans, x = 0, num
	for x != 0 {
		y := x % 10
		if num%y == 0 {
			ans++
		}
		x /= 10
	}
	return ans
}
