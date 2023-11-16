package main

// 贪心：原地换座
// 时间复杂度：O(n^2)
// 空间复杂度：O(1)
func minSwapsCouples(row []int) int {
	var n = len(row)
	var ans = 0

	for i := 0; i < n; i += 2 {
		l, r := row[i]/2, row[i+1]/2
		if l == r {
			continue
		}

		for j := i + 2; j < n; j++ {
			// 贪心：原地交换
			// 如何证明贪心的正确性？
			if row[j]/2 == r {
				row[j], row[i] = row[i], row[j]
				ans++
				break
			}
		}
	}
	return ans
}

func minSwapsCouples2(row []int) int {
	n := len(row)
	graph := make([][]int, n/2)
	for i := 0; i < n; i += 2 {
		l, r := row[i]/2, row[i+1]/2
		if l != r {
			graph[l] = append(graph[l], r)
			graph[r] = append(graph[r], l)
		}
	}

	var ans = 0
	visit := make([]bool, n/2)
	for i, _ := range visit {
		if !visit[i] {
			visit[i] = true
			cnt := 0
			q := []int{i}
			for len(q) > 0 {
				cnt++
				v := q[0]
				q = q[1:]
				for _, w := range graph[v] {
					if !visit[w] {
						visit[w] = true
						q = append(q, w)
					}
				}
			}
			ans += cnt - 1
		}
	}
	return ans
}
