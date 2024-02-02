package dp

func MemorySearch[X comparable, Y any](dfs *func(X) Y) {
	var cache = make(map[X]Y, 1<<10)
	f := *dfs
	*dfs = func(x X) (res Y) {
		if v, ok := cache[x]; ok {
			return v
		}
		defer func() { cache[x] = res }()

		return f(x)
	}
}

type key2[X, Y any] struct {
	x X
	y Y
}

func MemorySearch2[X, Y comparable, Z any](dfs *func(X, Y) Z) {
	var cache = make(map[key2[X, Y]]Z, 1<<10)
	f := *dfs
	*dfs = func(x X, y Y) (res Z) {
		key := key2[X, Y]{x: x, y: y}
		if v, ok := cache[key]; ok {
			return v
		}
		defer func() { cache[key] = res }()

		return f(x, y)
	}
}
