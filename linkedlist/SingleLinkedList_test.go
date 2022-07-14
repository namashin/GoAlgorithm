package linkedlist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedList_Append(t *testing.T) {
	// Init
	l := NewSingleLinkedList()

	// Execution
	l.Append(5)
	l.Append(9)
	l.Append(2)
	l.Append(4)

	// Test
	assert.Equal(t, 5, l.head.value)
	assert.Equal(t, 9, l.head.next.value)
	assert.Equal(t, 2, l.head.next.next.value)
	assert.Equal(t, 4, l.head.next.next.next.value)

	// watch values
	l.Print()
}

func TestSingleLinkedList_Insert(t *testing.T) {
	// Init
	l := NewSingleLinkedList()
	l.Append(0)
	l.Append(1)
	l.Append(2)
	l.Append(3)

	// Execution
	l.Insert(4)

	// Test
	assert.Equal(t, 4, l.head.value)
	assert.Equal(t, 0, l.head.next.value)
	assert.Equal(t, 1, l.head.next.next.value)
	assert.Equal(t, 2, l.head.next.next.next.value)
	assert.Equal(t, 3, l.head.next.next.next.next.value)

	// watch values
	l.Print()
}

// 先頭を削除
func TestSingleLinkedList_RemoveElement(t *testing.T) {
	// Init
	l := NewSingleLinkedList()
	l.Append(0)
	l.Append(1)
	l.Append(2)
	l.Append(3)

	// Execution
	l.RemoveElement(0)

	// Test
	assert.Equal(t, 1, l.head.value)
	assert.Equal(t, 2, l.head.next.value)
	assert.Equal(t, 3, l.head.next.next.value)

	// watch values
	l.Print()
}

// 真ん中を削除
func TestSingleLinkedList_RemoveElement2(t *testing.T) {
	// Init
	l := NewSingleLinkedList()
	l.Append(0)
	l.Append(1)
	l.Append(2)
	l.Append(1)
	l.Append(3)

	// Execution
	l.RemoveElement(1)

	// Test
	assert.Equal(t, 0, l.head.value)
	assert.Equal(t, 2, l.head.next.value)
	assert.Equal(t, 1, l.head.next.next.value)
	assert.Equal(t, 3, l.head.next.next.next.value)

	// watch values
	l.Print()
}

// 最後尾を削除
func TestSingleLinkedList_RemoveElement3(t *testing.T) {
	// Init
	l := NewSingleLinkedList()
	l.Append(0)
	l.Append(1)
	l.Append(2)
	l.Append(3)

	// Execution
	l.RemoveElement(3)

	// Test
	assert.Equal(t, 0, l.head.value)
	assert.Equal(t, 1, l.head.next.value)
	assert.Equal(t, 2, l.head.next.next.value)

	// watch values
	l.Print()
}

func TestSingleLinkedList_RemoveElements(t *testing.T) {
	// Init
	l := NewSingleLinkedList()
	l.Append(1)
	l.Append(1)
	l.Append(0)
	l.Append(1)
	l.Append(2)
	l.Append(1)
	l.Append(1)
	l.Append(3)
	l.Append(1)
	l.Append(1)

	// Execution
	l.RemoveElements(1)

	// Test
	assert.Equal(t, 0, l.head.value)
	assert.Equal(t, 2, l.head.next.value)
	assert.Equal(t, 3, l.head.next.next.value)

	// watch values
	l.Print()
}

func TestSingleLinkedList_Reverse(t *testing.T) {
	// Init
	l := NewSingleLinkedList()
	l.Append(0)
	l.Append(1)
	l.Append(2)
	l.Append(3)

	// Execution
	l.ReverseIterative()

	// Test
	assert.Equal(t, 3, l.head.value)
	assert.Equal(t, 2, l.head.next.value)
	assert.Equal(t, 1, l.head.next.next.value)
	assert.Equal(t, 0, l.head.next.next.next.value)

	// watch values
	l.Print()
}

func TestSingleLinkedList_ReverseRecursive(t *testing.T) {
	// Init
	l := NewSingleLinkedList()
	l.Append(0)
	l.Append(1)
	l.Append(2)
	l.Append(3)

	// Execution
	l.ReverseRecursive()

	// Test
	assert.Equal(t, 3, l.head.value)
	assert.Equal(t, 2, l.head.next.value)
	assert.Equal(t, 1, l.head.next.next.value)
	assert.Equal(t, 0, l.head.next.next.next.value)

	// watch values
	l.Print()
}

func TestSingleLinkedList_SortMyself(t *testing.T) {
	// Init
	l := NewSingleLinkedList()
	l.Append(6)
	l.Append(1)
	l.Append(9)
	l.Append(4)
	l.Append(7)
	l.Append(3)

	// Execution
	l.SortMyself()

	// Test
	assert.Equal(t, 1, l.head.value)
	assert.Equal(t, 3, l.head.next.value)
	assert.Equal(t, 4, l.head.next.next.value)
	assert.Equal(t, 6, l.head.next.next.next.value)
	assert.Equal(t, 7, l.head.next.next.next.next.value)
	assert.Equal(t, 9, l.head.next.next.next.next.next.value)

	// watch values
	l.Print()
}

func TestSingleLinkedList_Size(t *testing.T) {
	// Init
	l := NewSingleLinkedList()
	l.Append(6)
	l.Append(1)
	l.Append(9)
	l.Append(4)
	l.Append(7)
	l.Append(3)

	// Execution
	got := l.Size()

	// Test
	assert.Equal(t, 6, got)

	// watch values
	l.Print()
}

func TestSingleLinkedList_ReverseEven(t *testing.T) {
	// Init
	l := NewSingleLinkedList()
	l.Append(6)
	l.Append(2)
	l.Append(4)
	l.Append(5)
	l.Append(8)
	l.Append(10)

	// Execution
	l.ReverseEven()

	// Test
	assert.Equal(t, 4, l.head.value)
	assert.Equal(t, 2, l.head.next.value)
	assert.Equal(t, 6, l.head.next.next.value)
	assert.Equal(t, 5, l.head.next.next.next.value)
	assert.Equal(t, 10, l.head.next.next.next.next.value)
	assert.Equal(t, 8, l.head.next.next.next.next.next.value)

	// watch values
	l.Print()
}

func TestSingleLinkedList_MiddleNode(t *testing.T) {
	// Init
	l := NewSingleLinkedList()
	l.Append(6)
	l.Append(2)
	l.Append(4)

	l.Append(5)

	l.Append(4)
	l.Append(2)
	l.Append(6)

	// Execution
	got := l.MiddleNode()

	// Test
	assert.Equal(t, got.value, 5)
}

func TestSingleLinkedList_IsPalindrome(t *testing.T) {
	// Init
	l := NewSingleLinkedList()
	l.Append(6)
	l.Append(2)
	l.Append(4)
	l.Append(4)
	l.Append(2)
	l.Append(6)

	// Execution and Test
	assert.True(t, l.IsPalindrome())

}

func TestSingleLinkedList_MergeTwoLinkedList(t *testing.T) {
	// Init
	l1 := NewSingleLinkedList()
	l1.Append(3)
	l1.Append(5)
	l1.Append(8)
	l1.Append(9)

	l2 := NewSingleLinkedList()
	l2.Append(0)
	l2.Append(2)
	l2.Append(5)
	l2.Append(7)

	// Execution
	l1.MergeTwoLinkedList(l2.head)

	// Test
	want := []int{0, 2, 3, 5, 5, 7, 8, 9}
	head := l1.head
	for _, value := range want {
		if head.value != value {
			t.Error("正しくマージソートされてません")
		}
		head = head.next
	}

	// Watch Values
	l1.Print()

}
