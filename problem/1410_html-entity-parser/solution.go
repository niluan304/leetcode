package main

import (
	"strings"
)

func entityParser(text string) string {
	var parse = map[string]string{
		`&quot;`:  `"`,
		`&apos;`:  `'`,
		`&amp;`:   `&`,
		`&gt;`:    `>`,
		`&lt;`:    `<`,
		`&frasl;`: `/`,
	}

	var b strings.Builder
	b.Grow(len(text))

	start := 0
	for {
		eIdx := start
		eIdx += strings.Index(text[start:], ";") + 1
		sIdx := strings.LastIndex(text[:eIdx], "&")
		if eIdx <= start || sIdx < start {
			break
		}

		b.WriteString(text[start:sIdx])
		p := text[sIdx:eIdx]
		if v, ok := parse[p]; ok {
			p = v
		}
		b.WriteString(p)
		start = eIdx
	}
	b.WriteString(text[start:])
	return b.String()
}
