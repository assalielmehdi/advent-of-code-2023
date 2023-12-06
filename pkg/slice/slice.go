package util_slice

import (
	util_fp "assalielmehdi/adventofcode2023/pkg/fp"
)

type Summable interface {
	int | int64 | float64
}

type Ordered interface {
	int | int64 | float64
}

func Sum[T Summable](values []T) T {
	sum := T(0)

	for _, value := range values {
		sum += value
	}

	return sum
}

func Min[T Ordered](values []T) T {
	if len(values) == 1 {
		return values[0]
	}

	return util_fp.Reduce(values, func(min, cur T) T {
		if min <= cur {
			return min
		}

		return cur
	}, values[0])
}
