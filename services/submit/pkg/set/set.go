package set

type Set[T int | int64] map[T]struct{}

func NewSet[T int | int64]() *Set[T] {
	return new(Set[T])
}

func (s Set[T]) Insert(x T) {
	if s == nil {
		s = make(Set[T])
	}
	s[x] = struct{}{}
}

func (s Set[T]) Find(x T) bool {
	_, ok := s[x]
	return ok
}

func (s Set[T]) Remove(x T) {
	delete(s, x)
}

func (s Set[T]) Size() int {
	return len(s)
}

func (s Set[T]) Range(f func(x T))  {
	for x := range s {
		f(x)
	}
}