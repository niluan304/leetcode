package main

func numberOfBoomerangs(points [][]int) int {
	var count = map[int]int{}
	var ans = 0
	for _, point := range points {
		clear(count)
		p := Point{X: point[0], Y: point[1]}
		for _, q := range points {
			d := p.Distance(q[0], q[1])
			count[d]++
		}
		for _, cnt := range count {
			ans += cnt * (cnt - 1)
		}
	}
	return ans
}

type Point struct{ X, Y int }

func (p Point) Distance(x, y int) int {
	x, y = p.X-x, p.Y-y
	return x*x + y*y
}

func numberOfBoomerangs2(points [][]int) int {
	var count = map[int]int{}
	var ans = 0
	for _, p := range points {
		clear(count)
		for _, q := range points {
			d := (p[0]-q[0])*(p[0]-q[0]) + (p[1]-q[1])*(p[1]-q[1])
			// ### 思考角度
			// 1. $C^{x}_{x+1} = C^{x-1}_{x} * 2$。
			// 2. 从怎么组成点对的角度思考，每个 j 都可以和左边的 cnt 个 k 组成 cnt*2 个点对。
			ans += count[d] * 2
			count[d]++
		}
	}
	return ans
}
