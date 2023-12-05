package util_fp

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
