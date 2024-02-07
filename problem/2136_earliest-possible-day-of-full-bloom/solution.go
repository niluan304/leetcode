package main

import (
	"sort"
)

func earliestFullBloom(plantTime []int, growTime []int) int {
	sort.Sort(&Sort{
		plantTime: plantTime,
		growTime:  growTime,
	})

	var bloom, day = 0, 0
	for i := range plantTime {
		day += plantTime[i]                 // 累加播种天数
		bloom = max(bloom, day+growTime[i]) // 再加上生长天数，就是这个种子的开花时间
	}

	return bloom
}

type Sort struct {
	plantTime []int
	growTime  []int
}

func (s *Sort) Len() int {
	return len(s.growTime)
}

func (s *Sort) Less(i, j int) bool {
	return s.growTime[i] > s.growTime[j]
}

func (s *Sort) Swap(i, j int) {
	s.growTime[i], s.growTime[j] = s.growTime[j], s.growTime[i]
	s.plantTime[i], s.plantTime[j] = s.plantTime[j], s.plantTime[i]
}
