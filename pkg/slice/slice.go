package util_slice

type Summable interface {
	int | int64 | float64
}

func Sum[T Summable](values []T) T {
	sum := T(0)

	for _, value := range values {
		sum += value
	}

	return sum
}
