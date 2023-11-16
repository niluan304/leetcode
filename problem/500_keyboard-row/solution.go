package main

func findWords(words []string) []string {
	var ans []string

Next:
	for _, word := range words {
		// word may be empty string ""
		if len(word) == 0 {
			continue
		}

		var line1 = Line(word[0])
		if line1 == 0 {
			continue
		}

		// find line
		for i := 1; i < len(word); i++ {
			if Line(word[i]) != line1 {
				continue Next
			}
		}

		ans = append(ans, word)
	}

	return ans
}

func Line(b byte) int {
	// to upper
	if b > 'Z' {
		b = b - ('z' - 'Z')
	}

	switch b {
	case 'Q', 'W', 'E', 'R', 'T', 'Y', 'U', 'I', 'O', 'P':
		return 1
	case 'A', 'S', 'D', 'F', 'G', 'H', 'J', 'K', 'L':
		return 2
	case 'Z', 'X', 'C', 'V', 'B', 'N', 'M':
		return 3
	default:
		return 0
	}
}

func findWords2(words []string) []string {
	const rowIdx = "12210111011122000010020202"
	var ans []string

Next:
	for _, word := range words {
		idx := rowIdx[Index(word[0])]
		// find line
		for i := 1; i < len(word); i++ {
			if rowIdx[Index(word[i])] != idx {
				continue Next
			}
		}

		ans = append(ans, word)
	}
	return ans
}

func Index(b byte) byte {
	if b > 'Z' {
		b = b - ('z' - 'Z')
	}
	return b - 'A'
}
