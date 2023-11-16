package main

func ReverseString(s []byte) {
	reverseString(s)
}

func reverseString(s []byte) {
	n := len(s)
	for i := 0; i < n/2; i++ {
		s[i], s[n-1-i] = s[n-1-i], s[i]
	}
}
