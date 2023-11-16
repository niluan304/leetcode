package main

// 拓扑排序 + bfs
// 时间复杂度：O(m+n)
// 空间复杂度：O(m+n)
func findOrder(numCourses int, prerequisites [][]int) []int {
	var (
		ingress = make([]int, numCourses)   //入度表
		graph   = make([][]int, numCourses) //邻接表
	)

	for _, p := range prerequisites {
		graph[p[1]] = append(graph[p[1]], p[0])
		ingress[p[0]]++
	}

	var queen = make([]int, 0, numCourses)
	for i, v := range ingress {
		if v != 0 {
			continue
		}
		queen = append(queen, i)
	}

	var ans = make([]int, 0, numCourses)
	for len(queen) > 0 {
		// 处理队列的头结点
		head := queen[0]
		queen = queen[1:]
		ans = append(ans, head)

		for _, w := range graph[head] {
			ingress[w]--
			// 前驱结点为空, 更新队列
			if ingress[w] == 0 {
				queen = append(queen, w)
			}
		}
	}

	if len(ans) != numCourses {
		return nil
	}
	return ans
}

// 拓扑排序 + dfs
// 时间复杂度：O(m+n)
// 空间复杂度：O(m+n)
func findOrder2(numCourses int, prerequisites [][]int) []int {
	const (
		WaitVisit = iota
		Visit
		HasVisited
	)
	var (
		edges   = make([][]int, numCourses)  //邻接表
		stack   = make([]int, 0, numCourses) //栈
		visited = make([]int, numCourses)

		dfs func(v int) (hasCycle bool)
	)
	for _, p := range prerequisites {
		edges[p[1]] = append(edges[p[1]], p[0])
	}

	dfs = func(u int) (hasCycle bool) {
		visited[u] = Visit

		for _, v := range edges[u] {
			switch visited[v] {
			case WaitVisit:
				if dfs(v) {
					return true
				}
			case Visit:
				return true
			}
		}
		visited[u] = HasVisited
		stack = append(stack, u)
		return false
	}

	for edge, _ := range edges {
		if visited[edge] == WaitVisit {
			if dfs(edge) {
				return nil
			}
		}
	}

	// 栈是逆序的, 原地翻转栈就为结果
	for i := 0; i < numCourses/2; i++ {
		j := numCourses - i - 1
		stack[i], stack[j] = stack[j], stack[i]
	}

	return stack
}
