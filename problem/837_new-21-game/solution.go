package main

func new21Game(n int, k int, maxPts int) float64 {
	if k == 0 {
		return 1.00
	}

	var (
		list = make([]float64, 0, k+maxPts)
	)

	pr := 1.0 / float64(maxPts)
	for i := 0; i < maxPts; i++ {
		list = append(list, pr)
	}

	var (
		left = 0
	)

	for left+1 < k {
		pr2 := list[left] * pr
		for i := left + 1; i < len(list); i++ {
			list[i] += pr2
		}

		list = append(list, pr2)
		left++
	}

	var ans = 0.00
	for i, item := range list[left:] {
		if i+left+1 <= n {
			ans += item
		}
	}

	return ans
}
