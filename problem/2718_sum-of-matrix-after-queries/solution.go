package main

func matrixSumQueries(n int, queries [][]int) int64 {
	type Item struct {
		Visit bool
		Val   int
	}

	var matrix = make([][]Item, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]Item, n)
	}

	for i := len(queries) - 1; i >= 0; i-- {
		t, idx, val := queries[i][0], queries[i][1], queries[i][2]
		switch t {
		case 0: // 将 idx行 的值更新为 val
			for j, _ := range matrix[idx] {
				if matrix[idx][j].Visit {
					continue
				}
				matrix[idx][j] = Item{Visit: true, Val: val}
			}
		case 1: // 将 idx列 的值更新为 val
			for j := 0; j < n; j++ {
				if matrix[j][idx].Visit {
					continue
				}
				matrix[j][idx] = Item{Visit: true, Val: val}
			}
		}
	}

	var ans = 0
	for _, row := range matrix {
		for _, item := range row {
			ans += item.Val
		}
	}
	return int64(ans)
}

func matrixSumQueries2(n int, queries [][]int) int64 {
	var visitRow, visitCol = make([]bool, n), make([]bool, n)
	var ans, count = 0, 0
	for i := len(queries) - 1; i >= 0; i-- {
		if count >= n*n {
			break
		}

		t, idx, val := queries[i][0], queries[i][1], queries[i][2]

		var curVisit, checkVisit = &visitRow, &visitCol // 将 idx行 的值更新为 val
		if t == 1 {
			curVisit, checkVisit = &visitCol, &visitRow // 将 idx列 的值更新为 val
		}

		if (*curVisit)[idx] {
			continue
		}
		(*curVisit)[idx] = true

		num := n
		for j := 0; j < n; j++ {
			if (*checkVisit)[j] {
				num--
			}
		}
		count += num
		ans += val * num
	}
	return int64(ans)
}

func matrixSumQueries3(n int, queries [][]int) int64 {
	var visitRow, visitCol = make(map[int]bool, n), make(map[int]bool, n)
	var ans = 0
	for i := len(queries) - 1; i >= 0; i-- {
		t, idx, val := queries[i][0], queries[i][1], queries[i][2]

		if t == 0 { // 将 idx行 的值更新为 val
			if visitRow[idx] {
				continue
			}
			ans += (n - len(visitCol)) * val
			visitRow[idx] = true
		} else { // 将 idx列 的值更新为 val
			if visitCol[idx] {
				continue
			}
			ans += (n - len(visitRow)) * val
			visitCol[idx] = true
		}
	}
	return int64(ans)
}

// 整合 matrixSumQueries3 的 if else
func matrixSumQueries4(n int, queries [][]int) int64 {
	var visit = [2]map[int]bool{make(map[int]bool), make(map[int]bool)}
	var ans = 0
	for i := len(queries) - 1; i >= 0; i-- {
		tp, idx, val := queries[i][0], queries[i][1], queries[i][2]
		if visit[tp][idx] { // tp == 0 表示计算行，tp == 1 表示计算列
			continue
		}
		visit[tp][idx] = true
		ans += (n - len(visit[tp^1])) * val
	}
	return int64(ans)
}
