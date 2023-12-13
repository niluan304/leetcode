package main

// 双指针 + 贪心
// 时间复杂度: O(n)
// 空间复杂度: O(1)
func makeSmallestPalindrome(s string) string {
	b := []byte(s)
	n := len(b)
	for i := 0; i < n/2; i++ {
		if b[i] > b[n-1-i] {
			b[i] = b[n-1-i]
		} else if b[i] < b[n-1-i] {
			b[n-1-i] = b[i]
		}
	}
	return string(b)
}
