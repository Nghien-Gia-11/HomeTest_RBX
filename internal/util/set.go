package util

type Set[T comparable] map[T]struct{}

func NewSet[T comparable]() Set[T] {
	return make(map[T]struct{})
}

func (s Set[T]) Add(value T) {
	s[value] = struct{}{}
}

func (s Set[T]) Remove(value T) {
	delete(s, value)
}

func (s Set[T]) Contains(value T) bool {
	_, ok := s[value]

	return ok
}

func (s Set[T]) Len() int {
	return len(s)
}

func (s Set[T]) Empty() bool {
	return len(s) == 0
}

func (s Set[T]) Values() []T {
	keys := make([]T, 0, s.Len()) // 0: số phần tử hiện tại, s.Len() capacity của mảng
	for k := range s {
		keys = append(keys, k)
	}
	return keys
}
