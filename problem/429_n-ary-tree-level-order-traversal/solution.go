package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	var ans [][]int
	queue := []*Node{root}

	for {
		tmp := queue
		queue = nil

		var path = make([]int, 0, len(tmp))
		for _, node := range tmp {
			if node == nil {
				continue
			}
			path = append(path, node.Val)
			for _, child := range node.Children {
				queue = append(queue, child)
			}
		}

		if len(path) == 0 {
			break
		}
		ans = append(ans, path)
	}
	return ans
}
