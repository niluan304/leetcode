package main

import "sort"

func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)
	var pairs = make([]int, len(spells))
	for i, spell := range spells {
		pairs[i] = len(potions) - sort.SearchInts(potions, int(success-1)/spell+1)
	}
	return pairs
}
