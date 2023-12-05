package util_ds

type Set[T comparable] struct {
	data map[T]bool
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		data: make(map[T]bool, 0),
	}
}

func (set *Set[T]) Size() int {
	return len(set.data)
}

func (set *Set[T]) Add(value T) *Set[T] {
	set.data[value] = true
	return set
}

func (set *Set[T]) Contains(value T) bool {
	_, exists := set.data[value]
	return exists
}

func (set *Set[T]) Retain(values *Set[T]) {
	toDelete := make([]T, 0)

	for key := range set.data {
		if !values.Contains(key) {
			toDelete = append(toDelete, key)
		}
	}

	for _, key := range toDelete {
		delete(set.data, key)
	}
}
