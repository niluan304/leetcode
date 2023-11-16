package main

import (
	"strings"
)

func restoreIpAddresses(s string) []string {
	n := len(s)
	if n < 4 {
		return nil
	}

	var ans []string

	for i := 1; i <= 4 && i <= n-3; i++ {
		if notIpNode(s[:i]) {
			continue
		}
		for j := i + 1; j <= i+4 && j <= n-2; j++ {
			if notIpNode(s[i:j]) {
				continue
			}
			for k := j + 1; k <= j+4 && k <= n-1; k++ {
				if notIpNode(s[j:k]) || notIpNode(s[k:]) {
					continue
				}
				ans = append(ans, s[:i]+"."+s[i:j]+"."+s[j:k]+"."+s[k:])
			}
		}
	}

	return ans
}

func notIpNode(node string) bool {
	switch len(node) {
	case 1:
		return false
	case 2, 3:
		if node[0] == '0' {
			return true
		}

		var v int32
		for _, b := range node {
			v = v*10 + (b - '0')
		}
		return v > 255
	default:
		return true
	}
}

func restoreIpAddresses2(s string) (ans []string) {
	var (
		path []string
	)

	var dfs func(s string)
	dfs = func(s string) {
		var n = len(s)
		if n == 0 {
			if len(path) == 4 {
				ans = append(ans, strings.Join(path, "."))
			}
			return
		}

		for i := 1; i <= n && i <= 3; i++ {
			if (4-len(path))*3 < n {
				return
			}

			node := s[:i]
			if i > 1 {
				// '0' 只能做后缀
				if node[0] == '0' {
					return
				}
				if i == 3 && node > "255" {
					return
				}
			}

			path = append(path, node)
			dfs(s[i:])
			path = path[:len(path)-1]
		}
	}

	dfs(s)
	return ans
}
