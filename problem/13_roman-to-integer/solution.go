package main

func romanToInt(s string) int {
	var sum = values[s[0]]
	for i := 1; i < len(s); i++ {
		var (
			s1 = values[s[i]]
			s0 = values[s[i-1]]
		)
		if s0 < s1 {
			sum += s1 - 2*s0
		} else {
			sum += s1
		}
	}

	return sum
}

var values = map[uint8]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt2(s string) int {
	var sum = value(s[0])
	for i := 1; i < len(s); i++ {
		var (
			s0 = value(s[i-1])
			s1 = value(s[i])
		)
		if s0 < s1 {
			sum += s1 - 2*s0
		} else {
			sum += s1
		}
	}

	return sum
}

func romanToInt3(s string) int {
	var sum = value(s[0])
	for i := 1; i < len(s); i++ {
		var (
			s0 = value(s[i-1])
			s1 = value(s[i])
		)
		if s0 < s1 {
			s1 -= 2 * s0
		}
		sum += s1
	}

	return sum
}

func value(r byte) int {
	switch r {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	}
	return 0
}
