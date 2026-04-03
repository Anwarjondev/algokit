package queue


type Queue[T any] struct {
	queue []T
	head int
}

func New[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Size() int {
	return len(q.queue) - q.head
}

func (q *Queue[T]) Enqueue(val T) {
	q.queue = append(q.queue, val)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	var zeroVal T
	if q.head >= len(q.queue) {
		return zeroVal, false
	}
	val := q.queue[q.head]
	q.queue[q.head] = zeroVal
	q.head ++
	if q.head > 100 && q.head * 2 > q.Size() {
		q.queue = q.queue[q.head:]
		q.head = 0
	}
	return val, true
}

func (q *Queue[T]) Front() (T, bool) {
	if q.Size() == 0 {
		var zeroVal T
		return zeroVal, false
	}
	return q.queue[q.head], true
}

func (q *Queue[T]) Back() (T, bool) {
	if q.Size() == 0 {
		var zeroVal T
		return zeroVal, false
	}
	return q.queue[len(q.queue)-1], true
}

func (q *Queue[T]) IsEmpty() bool {
	return q.Size() == 0
}
