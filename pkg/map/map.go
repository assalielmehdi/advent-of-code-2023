package util_map

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
