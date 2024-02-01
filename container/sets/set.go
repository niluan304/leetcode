package sets

import (
	"maps"

	"github.com/niluan304/leetcode/container"
)

func NewSet[T comparable](data ...T) *Set[T] {
	set := Set[T]{data: make(map[T]struct{}, len(data))}
	set.Add(data...)

	return &set
}

type Set[T comparable] struct {
	data map[T]struct{}
}

func (s *Set[T]) Add(values ...T) {
	for _, v := range values {
		s.data[v] = struct{}{}
	}
}

func (s *Set[T]) Remove(values ...T) {
	for _, v := range values {
		delete(s.data, v)
	}
}

func (s *Set[T]) Contain(x T) bool { _, ok := s.data[x]; return ok }
func (s *Set[T]) Clear()           { clear(s.data) }

func (s Set[T]) Size() int   { return len(s.data) }
func (s Set[T]) Empty() bool { return s.Size() == 0 }

func (s Set[T]) Values() []T {
	return container.MapKeys(s.data)
}

func (s Set[T]) Clone() *Set[T] {
	return &Set[T]{
		data: maps.Clone(s.data),
	}
}
