package main

func partition(s string) [][]string {
	var stack []string

	nodes = [][]string{}
	dfs(s, 0, stack)
	return nodes
}

func dfs(s string, idx int, stack []string) {
	if idx == len(s) {
		tmp := make([]string, len(stack))
		copy(tmp, stack)
		nodes = append(nodes, tmp)
		return
	}
	i := idx
	for i < len(s) {
		i++
		top := s[idx:i]
		if !isPartition(top) {
			continue
		}

		stack = append(stack, top)
		dfs(s, i, stack)
		stack = stack[:len(stack)-1]
	}

	return
}

func isPartition(s string) bool {
	j := len(s)
	for i := 0; i < j/2; i++ {
		if s[i] != s[j-1-i] {
			return false
		}
	}
	return true
}

var nodes [][]string

func partition2(s string) [][]string {
	nodes = [][]string{}
	dfs2(s, []string{})
	return nodes
}

func dfs2(s string, path []string) {
	if s == "" {
		var node = make([]string, len(path))
		copy(node, path)
		nodes = append(nodes, node)
		return
	}

	for i := 1; i <= len(s); i++ {
		if !isPartition(s[:i]) {
			continue
		}
		path = append(path, s[:i])

		dfs2(s[i:], path)
		path = path[:len(path)-1]
	}
}
