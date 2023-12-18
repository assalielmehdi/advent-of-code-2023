package util

type Comparable[T any] interface {
	Less(T) bool
}

type PriorityQueue[T Comparable[T]] []T

func NewPriorityQueue[T Comparable[T]]() *PriorityQueue[T] {
	pq := PriorityQueue[T](make([]T, 0))
	return &pq
}

func (pq PriorityQueue[T]) Len() int {
	return len(pq)
}

func (pq PriorityQueue[T]) Less(i, j int) bool {
	return pq[i].Less(pq[j])
}

func (pq PriorityQueue[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue[T]) Push(t any) {
	*pq = append(*pq, t.(T))
}

func (pq *PriorityQueue[T]) Pop() any {
	n := len(*pq)
	top := (*pq)[n-1]
	*pq = (*pq)[0 : n-1]
	return top
}
