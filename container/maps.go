package container

func MapKeys[M ~map[K]V, K comparable, V any](m M) (keys []K) {
	for k, _ := range m {
		keys = append(keys, k)
	}
	return keys
}
