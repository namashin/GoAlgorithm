package linkedlist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDoublyLinkedList_Append(t *testing.T) {
	// Init
	dll := NewDoublyLInkedList()

	// Execution
	dll.Append(5)
	dll.Append(2)
	dll.Append(8)
	dll.Append(6)

	// Test
	assert.Equal(t, 5, dll.head.value)
	assert.Equal(t, 5, dll.head.next.prev.value)

	assert.Equal(t, 2, dll.head.next.value)
	assert.Equal(t, 2, dll.head.next.next.prev.value)

	assert.Equal(t, 8, dll.head.next.next.value)
	assert.Equal(t, 8, dll.head.next.next.next.prev.value)

	assert.Equal(t, 6, dll.head.next.next.next.value)

	// watch values
	dll.Print()
}

func TestDoublyLinkedList_Insert(t *testing.T) {
	// Init
	dll := NewDoublyLInkedList()
	dll.Append(5)
	dll.Append(2)
	dll.Append(8)
	dll.Append(6)

	// Execution
	dll.Insert(0)

	// Test
	assert.Equal(t, 0, dll.head.value)
	assert.Equal(t, 0, dll.head.next.prev.value)

	assert.Equal(t, 5, dll.head.next.value)
	assert.Equal(t, 5, dll.head.next.next.prev.value)

	assert.Equal(t, 2, dll.head.next.next.value)
	assert.Equal(t, 2, dll.head.next.next.next.prev.value)

	assert.Equal(t, 8, dll.head.next.next.next.value)
	assert.Equal(t, 8, dll.head.next.next.next.next.prev.value)

	assert.Equal(t, 6, dll.head.next.next.next.next.value)

	// watch values
	dll.Print()
}

func TestDoublyLinkedList_Size(t *testing.T) {
	// Init
	dll := NewDoublyLInkedList()
	dll.Append(5)
	dll.Append(2)
	dll.Append(8)
	dll.Append(6)

	// Execution
	got := dll.Size()

	// Test
	assert.Equal(t, 4, got)

	// watch values
	dll.Print()
}

func TestDoublyLinkedList_ReverseIterative(t *testing.T) {
	// Init
	dll := NewDoublyLInkedList()
	dll.Append(5)
	dll.Append(2)
	dll.Append(8)
	dll.Append(6)
	dll.Append(1)

	// Execution
	dll.ReverseIterative()

	// Test
	assert.Equal(t, 1, dll.head.value)
	assert.Equal(t, 6, dll.head.next.value)
	assert.Equal(t, 8, dll.head.next.next.value)
	assert.Equal(t, 2, dll.head.next.next.next.value)
	assert.Equal(t, 5, dll.head.next.next.next.next.value)

	// watch values
	dll.Print()
}

func TestDoublyLinkedList_ReverseRecursive(t *testing.T) {
	// Init
	dll := NewDoublyLInkedList()
	dll.Append(5)
	dll.Append(2)
	dll.Append(8)
	dll.Append(6)
	dll.Append(1)

	// Execution
	dll.ReverseRecursive()

	// Test
	assert.Equal(t, 1, dll.head.value)
	assert.Equal(t, 6, dll.head.next.value)
	assert.Equal(t, 8, dll.head.next.next.value)
	assert.Equal(t, 2, dll.head.next.next.next.value)
	assert.Equal(t, 5, dll.head.next.next.next.next.value)

	// watch values
	dll.Print()
}
