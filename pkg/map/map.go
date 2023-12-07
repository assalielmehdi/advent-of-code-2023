package util_map

import (
	"cmp"
)

type Summable interface {
	int | int64 | float64
}

func Sum[K comparable, V Summable](values map[K]V) V {
	sum := V(0)

	for _, value := range values {
		sum += value
	}

	return sum
}

func MaxVal[K comparable, V cmp.Ordered](values map[K]V) V {
	var max *V

	for _, val := range values {
		if max == nil {
			max = &val
		} else if val > *max {
			max = &val
		}
	}

	return *max
}

func MinVal[K comparable, V cmp.Ordered](values map[K]V) V {
	var min *V

	for _, val := range values {
		if min == nil {
			min = &val
		} else if val < *min {
			min = &val
		}
	}

	return *min
}

func CountVal[K, V comparable](values map[K]V, v V) int {
	c := 0

	for _, val := range values {
		if val == v {
			c++
		}
	}

	return c
}
