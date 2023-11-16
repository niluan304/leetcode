package main

func generate(numRows int) [][]int {
	var rows = [][]int{{1}}

	for i := 1; i < numRows; i++ {
		var row = []int{1}
		for j := 1; j < i; j++ {
			row = append(row, rows[i-1][j-1]+rows[i-1][j])
		}
		row = append(row, 1)
		rows = append(rows, row)
	}

	return rows
}

func leetcode(numRows int) [][]int {
	ans := make([][]int, numRows)
	for i := range ans {
		ans[i] = make([]int, i+1)
		ans[i][0] = 1
		ans[i][i] = 1
		for j := 1; j < i; j++ {
			ans[i][j] = ans[i-1][j] + ans[i-1][j-1]
		}
	}
	return ans
}
