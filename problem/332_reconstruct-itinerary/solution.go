package main

import (
	"sort"
)

const Start = "JFK"

func findItinerary(tickets [][]string) []string {
	type To struct {
		Value string
		Used  bool
	}

	var m = map[string][]To{}
	for _, ticket := range tickets {
		var (
			from = ticket[0]
			to   = ticket[1]
		)

		m[from] = append(m[from], To{
			Value: to,
			Used:  false,
		})
	}

	for form, tos := range m {
		sort.Slice(tos, func(i, j int) bool {
			return tos[i].Value < tos[j].Value
		})

		m[form] = tos
	}

	var ans []string
	var tmp = make([]string, 0, len(tickets))
	var found = false

	var dfs func(from string)
	dfs = func(from string) {
		if len(tmp) == len(tickets) {
			ans = append([]string{Start}, tmp...)
			found = true
			return
		}

		for i := range m[from] {
			if found {
				return
			}

			v := &(m[from][i])
			if v.Used {
				continue
			}

			tmp = append(tmp, v.Value)
			v.Used = true

			dfs(v.Value)

			tmp = tmp[:len(tmp)-1]
			v.Used = false
		}
	}

	dfs(Start)
	return ans
}

func findItinerary2(tickets [][]string) []string {
	type To struct {
		Value string
		Used  bool
	}

	var m = map[string][]*To{}
	for _, ticket := range tickets {
		var (
			from = ticket[0]
			to   = ticket[1]
		)

		m[from] = append(m[from], &To{
			Value: to,
			Used:  false,
		})
	}

	for from := range m {
		sort.Slice(m[from], func(i, j int) bool {
			return m[from][i].Value < m[from][j].Value
		})
	}

	var ans = append(make([]string, 0, len(tickets)), Start)
	var found = false

	var dfs func()
	dfs = func() {
		if len(ans) == len(tickets)+1 {
			found = true
			return
		}

		from := ans[len(ans)-1]
		for _, v := range m[from] {
			if v.Used {
				continue
			}

			ans = append(ans, v.Value)
			v.Used = true

			dfs()
			if found {
				return
			}

			ans = ans[:len(ans)-1]
			v.Used = false
		}
	}

	dfs()
	return ans
}
