package stack

import "testing"

// TestStackLinkedList for testing Stack with LinkedList
func TestStackLinkedList(t *testing.T) {
	s := New[int]()

	s.Push(1)
	s.Push(2)

	t.Run("Stack Push", func(t *testing.T) {
		result := s.Show()
		expected := []any{2, 1}
		for x := range result {
			if result[x] != expected[x] {
				t.Errorf("Stack Push is not work, got %v but expected %v", result, expected)
			}
		}
	})

	t.Run("Stack isEmpty", func(t *testing.T) {
		if s.IsEmpty() {
			t.Error("Stack isEmpty is returned true but expected false", s.IsEmpty())
		}
	})

	t.Run("Stack Length", func(t *testing.T) {
		if s.Size != 2 {
			t.Error("Stack Length should be 2 but got", s.Size)
		}
	})

	s.Pop()
	pop := s.Pop()

	t.Run("Stack Pop", func(t *testing.T) {
		if *pop != 1 {
			t.Error("Stack Pop should return 1 but is returned", pop)
		}

	})

	s.Push(52)
	s.Push(23)
	s.Push(99)
	t.Run("Stack Peek", func(t *testing.T) {
		if *s.Peek() != 99 {
			t.Error("Stack Peak should return 99 but got ", s.Peek())
		}
	})
}
