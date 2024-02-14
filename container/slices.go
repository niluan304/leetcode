package container

import (
	"cmp"
)

func Sum[S ~[]E, E cmp.Ordered](x S) E {
	var m E
	for i := 0; i < len(x); i++ {
		m += x[i]
	}
	return m
}

func Unique[T comparable](values []T) []T {
	exist := make(map[T]bool, len(values))
	ans := make([]T, 0, len(values))
	for _, v := range values {
		if exist[v] {
			continue
		}
		ans = append(ans, v)
		exist[v] = true
	}
	return ans
}
