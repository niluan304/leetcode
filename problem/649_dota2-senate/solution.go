package main

import (
	"strings"
)

func predictPartyVictory(senate string) string {
	const Dire, Radiant = "D", "R"

	for {
		if !strings.Contains(senate, Dire) {
			return "Radiant"
		}
		if !strings.Contains(senate, Radiant) {
			return "Dire"
		}

		if senate[0] == 'R' {
			senate = strings.Replace(senate[1:]+senate[:1], Dire, "", 1)
		} else {
			senate = strings.Replace(senate[1:]+senate[:1], Radiant, "", 1)
		}
	}
}

// *复杂度分析**
//
// - 时间复杂度：$O(n)$，其中 $n$ 是字符串 $\textit{senate}$ 的长度。在模拟整个投票过程的每一步，我们进行的操作的时间复杂度均为 $O(1)$，并且会弹出一名天辉方或夜魇方的议员。由于议员的数量为 $n$，因此模拟的步数不会超过 $n$，时间复杂度即为 $O(n)$。
//
// - 空间复杂度：$O(n)$，即为两个队列需要使用的空间。
func predictPartyVictory2(senate string) string {
	var radiant, dire []int
	for i, s := range senate {
		if s == 'R' {
			radiant = append(radiant, i)
		} else {
			dire = append(dire, i)
		}
	}
	for len(radiant) > 0 && len(dire) > 0 {
		if radiant[0] < dire[0] {
			radiant = append(radiant, radiant[0]+len(senate))
		} else {
			dire = append(dire, dire[0]+len(senate))
		}
		radiant = radiant[1:]
		dire = dire[1:]
	}
	if len(radiant) > 0 {
		return "Radiant"
	}
	return "Dire"
}
