package main

func countServers(grid [][]int) int {
	row := make([]int, len(grid))
	col := make([]int, len(grid[0]))

	for i := range grid {
		for j, v := range grid[i] {
			if v == 1 {
				row[i]++
				col[j]++
			}
		}
	}

	count := 0
	for i := range grid {
		for j, v := range grid[i] {
			if v == 1 && (row[i] > 1 || col[j] > 1) {
				count++
			}
		}
	}
	return count
}
