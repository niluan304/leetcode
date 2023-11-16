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
