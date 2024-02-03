package queue

import "testing"

func TestQueueFIFO(t *testing.T) {
	q := New[int]()
	for i := 0; i < 10000; i++ {
		q.Enqueue(i)
	}

	// Retrieve 10k elements from the queue.
	// they SHOULD be in order
	current := q.Head
	for i := 0; i < 10000; i++ {
		// Value should be i, because it is fifo
		v := current.Value
		if v != i {
			t.Fatalf("The popped element should be %d, got %d", i, v)
		}
		current = current.Next
	}
}

func TestDequeue(t *testing.T) {
	q := New[int]()
	for i := 0; i < 10000; i++ {
		q.Enqueue(i)
	}
	for i := 0; i < 10000; i++ {
		v := q.Dequeue()
		if *v != i {
			t.Fatalf("value should be %v got %v", i, *v)
		}
	}
	if q.Size != 0 {
		t.Fatal("size should be 0")
	}
	if q.Head != nil {
		t.Fatal("Head should be nil")
	}
	if q.Tail != nil {
		t.Fatal("Tail should be nil")
	}
}
