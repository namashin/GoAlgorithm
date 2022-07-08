package stack_queue

import (
	"testing"
)

func TestStack_Push(t *testing.T) {
	// Init
	s := NewStack()

	// Execution
	s.Push(6)
	s.Push(2)
	s.Push(1)
	s.Push(7)
	s.Push(4)
	s.Push(8)

	// Test
	want := []int{6, 2, 1, 7, 4, 8}

	for i, value := range s.stack {
		if want[i] != value {
			t.Fail()
		}
	}
}

func TestStack_Pop(t *testing.T) {
	// Init
	s := NewStack()

	s.Push(6)
	s.Push(2)
	s.Push(1)
	s.Push(7)
	s.Push(4)
	s.Push(8)

	// Execution
	got := s.Pop()

	// Test
	want := 8
	if got != want {
		t.Fail()
	}

	want2 := []int{6, 2, 1, 7, 4}
	for i, value := range s.stack {
		if want2[i] != value {
			t.Errorf("want %v, but got %v", want2[i], value)
		}
	}

	if len(s.stack) != 5 {
		t.Fail()
	}
}

func TestQueue_Push(t *testing.T) {
	// Init
	q := NewQueue()

	// Execution
	q.Push(5)
	q.Push(8)
	q.Push(1)
	q.Push(3)
	q.Push(0)
	q.Push(6)

	// Test
	want := []int{5, 8, 1, 3, 0, 6}

	for i, value := range q.queue {
		if want[i] != value {
			t.Fail()
		}
	}

	if len(q.queue) != 6 {
		t.Fail()
	}
}

func TestQueue_PopLeft(t *testing.T) {
	// Init
	q := NewQueue()

	q.Push(5)
	q.Push(8)
	q.Push(1)
	q.Push(3)
	q.Push(0)
	q.Push(6)

	// Execution
	got := q.PopLeft()

	// Test
	want := []int{8, 1, 3, 0, 6}
	for i, value := range q.queue {
		if want[i] != value {
			t.Errorf("want %v, but got %v", want[i], value)
		}
	}

	if got != 5 {
		t.Errorf("You got %v", got)
	}

	if len(q.queue) != 5 {
		t.Fail()
	}
}
