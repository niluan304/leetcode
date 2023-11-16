package main

import (
	"strings"
)

func simplifyPath(path string) string {
	list := strings.Split(path, "/")

	var stack = make([]string, 0, len(list))
	for _, s := range list {
		switch s {
		// 点".", 处于当前目录
		// 空格"": 有多个连续"/", 无效路径
		case ".", "":
			continue
		case "..":
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		default:
			stack = append(stack, s)
		}
	}

	return "/" + strings.Join(stack, "/")
}
