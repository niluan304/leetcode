package main

import "sort"

func findMinDifference(timePoints []string) int {
	sort.Strings(timePoints)
	n := len(timePoints)
	t0 := Minutes(timePoints[0])
	tn := Minutes(timePoints[n-1])

	mn := t0 + 24*60 - tn
	for i := 1; i < n; i++ {
		ti := Minutes(timePoints[i])
		sub := ti - t0
		t0 = ti
		if sub < mn {
			mn = sub
		}
	}

	return mn
}

func Minutes(point string) int {
	const zero = '0'
	return int(point[4]-zero) +
		int(point[3]-zero)*10 +
		int(point[1]-zero)*60 +
		int(point[0]-zero)*600
}
