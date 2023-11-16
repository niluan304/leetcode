package main

// [聊聊算法]拓扑排序原理及实战
// https://www.luozhiyun.com/archives/748
func canFinish(numCourses int, prerequisites [][]int) bool {
	var (
		ingress = make([]int, numCourses)         //入度表
		graph   = make(map[int][]int, numCourses) //邻接表
	)
	for _, v := range prerequisites {
		ingress[v[0]]++ //入度加1
		if _, ok := graph[v[1]]; !ok {
			graph[v[1]] = []int{}
		}
		// 建立关联关系，存储和v[1]课关联课程集
		graph[v[1]] = append(graph[v[1]], v[0])
	}

	var queen = make([]int, 0, numCourses)
	for i, v := range ingress {
		if v == 0 { //找出所有入度为0的课程
			queen = append(queen, i)
		}
	}

	count := 0
	for len(queen) > 0 {
		//依次将队首节点出队
		first := queen[0]
		queen = queen[1:]
		count++ // 记录入度为 0 的课程个数,表示这个课程已经可以修炼了

		for _, v := range graph[first] { //找出与这个课程关联的所有课程
			ingress[v]--         //因为它的前提课程可以修炼了，所以这个课程的入度减一
			if ingress[v] == 0 { //如果入度都为0了，那么代表这个课程可以修炼了，入队
				queen = append(queen, v)
			}
		}
	}
	return count == numCourses //如果可以修炼的课程和课程总数相等，那么表示可以完成所有课程的修炼
}

// 拓扑排序
// 时间复杂度：O(m*n)
// 空间复杂度：O(m*n)
func canFinish2(numCourses int, prerequisites [][]int) bool {
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

	var count = 0
	for len(queen) > 0 {
		// 处理队列的头结点
		head := queen[0]
		queen = queen[1:]
		count++

		for _, w := range graph[head] {
			ingress[w]--
			// 前驱结点为空, 更新队列
			if ingress[w] == 0 {
				queen = append(queen, w)
			}
		}
	}
	return count == numCourses
}
