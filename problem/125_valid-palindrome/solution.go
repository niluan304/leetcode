package main

func isPalindrome(s string) bool {

	left, right := 0, len(s)-1
	for left < right {
		l := lower(s[left])
		if l == 0 {
			left++
			continue
		}

		r := lower(s[right])
		if r == 0 {
			right--
			continue
		}

		if l != r {
			return false
		}

		left++
		right--
	}

	return true
}

func lower(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b + 32
	}

	if (b >= 'a' && b <= 'z') || (b >= '0' && b <= '9') {
		return b
	}

	return 0
}

func isPalindrome2(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		for left < right && lower(s[left]) == 0 {
			left++
		}

		for left < right && lower(s[right]) == 0 {
			right--
		}

		if lower(s[left]) != lower(s[right]) {
			return false
		}
		left++
		right--
	}

	return true
}

func isPalindrome3(s string) bool {
	left, right := 0, len(s)-1
	l := func() byte {
		return lower(s[left])
	}

	r := func() byte {
		return lower(s[right])
	}

	for left < right {
		for left < right && l() == 0 {
			left++
		}

		for left < right && r() == 0 {
			right--
		}

		if l() != r() {
			return false
		}
		left++
		right--
	}

	return true
}
