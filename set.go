package main

type Set[V comparable] map[V]struct{}

func NewSet[V comparable](size int) Set[V] {
	return make(Set[V], size)
}

func (s Set[V]) Add(values ...V) {
	for _, val := range values {
		s[val] = struct{}{}
	}
}

func (s Set[V]) Has(value V) bool {
	_, exists := s[value]
	return exists
}

func (s Set[V]) Remove(value V) {
	delete(s, value)
}

func (s Set[V]) Values() []V {
	a := make([]V, 0, len(s))
	for key := range s {
		a = append(a, key)
	}
	return a
}

func (s Set[V]) Filter(f func(value V) bool) []V {
	a := make([]V, 0)
	for key := range s {
		if f(key) {
			a = append(a, key)
		}
	}
	return a
}

func (s Set[V]) Apply(f func(value V) V) []V {
	a := make([]V, 0, len(s))
	for key := range s {
		a = append(a, f(key))
	}
	return a
}
