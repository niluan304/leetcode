package main

// dfs
// 时间复杂度: O(n)
// 空间复杂度: O(n)
func minimumFuelCost(roads [][]int, seats int) int64 {
	var n = len(roads) + 1
	var graph = make([][]int, n)
	for _, road := range roads {
		a, b := road[0], road[1]
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	var dfs func(i int, fa int) (people int, cars int, cost int)
	dfs = func(i int, fa int) (people int, cars int, cost int) {
		for _, j := range graph[i] {
			if j == fa {
				continue
			}
			p, car, c := dfs(j, i)
			people += p
			cars += car
			cost += c
		}
		people++
		cost += cars
		cars = (people-1)/seats + 1
		return
	}

	_, _, ans := dfs(0, -1)
	return int64(ans)
}

// dfs
// 时间复杂度: O(n)
// 空间复杂度: O(n)
func minimumFuelCost2(roads [][]int, seats int) int64 {
	var graph = make([][]int, len(roads)+1)
	for _, road := range roads {
		a, b := road[0], road[1]
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}
	var ans = 0
	var dfs func(i int, fa int) (people int, cars int)
	dfs = func(i int, fa int) (people int, cars int) {
		for _, j := range graph[i] {
			if j == fa {
				continue
			}
			p, car := dfs(j, i)
			people += p
			cars += car
		}
		ans += cars

		people++
		cars = (people-1)/seats + 1
		return
	}

	dfs(0, -1)
	return int64(ans)
}
