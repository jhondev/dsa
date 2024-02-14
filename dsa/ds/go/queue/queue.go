package queue

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

type Queue[T any] struct {
	Head *Node[T]
	Tail *Node[T]
	Size int
}

func New[T any]() *Queue[T] {
	return &Queue[T]{
		Size: 0,
	}
}

func (q *Queue[T]) Enqueue(value T) {
	q.Size++
	if q.Head == nil {
		q.Head = &Node[T]{Value: value}
		q.Tail = q.Head
		return
	}
	q.Tail.Next = &Node[T]{Value: value}
	q.Tail = q.Tail.Next
}

func (q *Queue[T]) Dequeue() *T {
	if q.Head == nil {
		return nil
	}
	q.Size--
	v := q.Head.Value
	q.Head = q.Head.Next
	if q.Size == 0 {
		q.Tail = nil
	}
	return &v
	// TODO: free memory
}

func (q *Queue[T]) Peak() T {
	return q.Head.Value
}
