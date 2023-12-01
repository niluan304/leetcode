package main

// 哈希表
// 时间复杂度: O(m*n)
// 空间复杂度: O(m*n)
func firstCompleteIndex(arr []int, mat [][]int) int {
	var m, n = len(mat), len(mat[0])
	var row, col = make([]int, m), make([]int, n)

	type Value struct{ i, j int }
	//var idx = map[int]Value{}
	var idx = make([]Value, m*n+1) // arr 和 mat 都包含范围 [1，m * n] 内的 所有 整数。
	for i, _ := range mat {
		for j, v := range mat[i] {
			idx[v] = Value{i: i, j: j}
		}
	}

	for i, v := range arr {
		x := idx[v]
		row[x.i]++
		col[x.j]++
		if row[x.i] == n || col[x.j] == m {
			return i
		}
	}
	return -1
}
