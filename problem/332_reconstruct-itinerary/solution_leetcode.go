package main

import "sort"

// 方法一：Hierholzer 算法
// 思路及算法
// Hierholzer 算法用于在连通图中寻找欧拉路径，其流程如下：
// 从起点出发，进行深度优先搜索。
// 每次沿着某条边从某个顶点移动到另外一个顶点的时候，都需要删除这条边。
// 如果没有可移动的路径，则将所在节点加入到栈中，并返回。
func leetcode1(tickets [][]string) []string {
	var (
		m   = map[string][]string{}
		res []string
	)

	for _, ticket := range tickets {
		src, dst := ticket[0], ticket[1]
		m[src] = append(m[src], dst)
	}
	for key := range m {
		sort.Strings(m[key])
	}

	var dfs func(curr string)
	dfs = func(curr string) {
		for {
			if v, ok := m[curr]; !ok || len(v) == 0 {
				break
			}
			tmp := m[curr][0]
			m[curr] = m[curr][1:]
			dfs(tmp)
		}
		res = append(res, curr)
	}

	dfs("JFK")
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}
	return res
}

type Pair struct {
	target string
	visted bool
}

type Pairs []*Pair

func (p Pairs) Len() int {
	return len(p)
}

func (p Pairs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Pairs) Less(i, j int) bool {
	return p[i].target < p[j].target
}

// 0 ms 的代码示例
func leetcodeMinTime(tickets [][]string) []string {
	ans := []string{"JFK"}
	targets := make(map[string]Pairs)

	for _, ticket := range tickets {
		if targets[ticket[0]] == nil {
			targets[ticket[0]] = make(Pairs, 0)
		}
		targets[ticket[0]] = append(targets[ticket[0]], &Pair{target: ticket[1], visted: false})
	}

	for k, _ := range targets {
		sort.Sort(targets[k])
	}

	var track func() bool
	track = func() bool {
		if len(tickets)+1 == len(ans) {
			return true
		}

		for _, pair := range targets[ans[len(ans)-1]] {
			if pair.visted {
				continue
			}

			ans = append(ans, pair.target)
			pair.visted = true
			if track() {
				return true
			}
			pair.visted = false
			ans = ans[:len(ans)-1]
		}
		return false

	}
	track()
	return ans

}

// 4.84 MB 的代码示例
func leetcodeMinMemory(tickets [][]string) []string {
	n := len(tickets)
	visited := make([]bool, n)
	tmp := make([]string, n+1)
	sort.Slice(tickets, func(i, j int) bool {
		if tickets[i][0] == tickets[j][0] {
			return tickets[i][1] < tickets[j][1]
		}
		return tickets[i][0] < tickets[j][0]
	})
	tmp[0] = "JFK"
	var backtrack func(int, string) []string
	backtrack = func(step int, prev string) []string {
		if step == n {
			return tmp
		}
		for i, ticket := range tickets {
			if visited[i] {
				continue
			}

			if step == 0 {
				if ticket[0] == "JFK" {
					visited[i] = true
					tmp[1] = ticket[1]
					t := backtrack(1, ticket[1])
					if t != nil {
						return t
					}
					visited[i] = false
				}
			} else {
				if ticket[0] == prev {
					tmp[step+1] = ticket[1]
					visited[i] = true
					t := backtrack(step+1, ticket[1])
					if t != nil {
						return t
					}
					visited[i] = false
				}
			}
		}
		return nil
	}
	return backtrack(0, "")
}
