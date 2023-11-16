package main

func longestCommonPrefix(strs []string) string {
	var ans []byte

	for i := 0; i < 200; i++ {
		for j := 0; j < len(strs); j++ {
			elem := strs[j]
			if len(elem) <= i || (j > 0 && elem[i] != strs[j-1][i]) {
				return string(ans)
			}
		}

		ans = append(ans, strs[0][i])
	}

	return string(ans)
}

func longestCommonPrefix2(strs []string) string {
	var (
		n = len(strs)

		str0 = strs[0]
	)

	for j := 0; j < len(str0); j++ {
		for i := 1; i < n; i++ {
			elem := strs[i]
			if (len(elem) <= j) || (elem[j] != str0[j]) {
				return str0[:j]
			}
		}
	}

	return str0
}
