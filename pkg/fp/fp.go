package util_fp

type Pair[T, U any] struct {
	First  T
	Second U
}

func ForEach[T any](values []T, consumeFunc func(int, T)) {
	for i, value := range values {
		consumeFunc(i, value)
	}
}

func Map[T, R any](values []T, mapFunc func(T) R) []R {
	mapped := make([]R, 0, len(values))

	for _, value := range values {
		mapped = append(mapped, mapFunc(value))
	}

	return mapped
}

func Filter[T any](values []T, filterFunc func(T) bool) []T {
	filtered := make([]T, 0)

	for _, value := range values {
		if filterFunc(value) {
			filtered = append(filtered, value)
		}
	}

	return filtered
}

func Reduce[T, R any](values []T, reduceFunc func(R, T) R, initialVal R) R {
	reduced := initialVal

	for _, value := range values {
		reduced = reduceFunc(reduced, value)
	}

	return reduced
}

func Zip[T, U any](values1 []T, values2 []U) []*Pair[T, U] {
	minLen := min(len(values1), len(values2))
	zipped := make([]*Pair[T, U], minLen)

	for i := 0; i < minLen; i++ {
		zipped[i] = &Pair[T, U]{
			First:  values1[i],
			Second: values2[i],
		}
	}

	return zipped
}
