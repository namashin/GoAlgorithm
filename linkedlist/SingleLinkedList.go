package linkedlist

import (
	"fmt"
	"reflect"
)

type Node struct {
	value int
	next  *Node
}

type SingleLinkedList struct {
	head *Node
}

func NewSingleLinkedList() *SingleLinkedList {
	return &SingleLinkedList{
		head: nil,
	}
}

func (sl *SingleLinkedList) Append(appendValue int) {
	var newNode = &Node{value: appendValue, next: nil}

	if sl.head == nil {
		sl.head = newNode
		return
	}

	current := sl.head
	for {
		if current.next == nil {
			break
		}
		current = current.next
	}
	current.next = newNode
}

func (sl SingleLinkedList) Print() {
	current := sl.head

	for {
		if current == nil {
			break
		}
		fmt.Println(current.value)
		current = current.next
	}
}

func (sl *SingleLinkedList) Size() int {
	size := 0
	current := sl.head

	for {
		if current == nil {
			break
		}

		current = current.next
		size += 1
	}

	return size
}

func (sl *SingleLinkedList) Insert(insertValue int) {
	var newNode = &Node{value: insertValue, next: nil}

	newNode.next = sl.head
	sl.head = newNode
}

func (sl *SingleLinkedList) RemoveElement(target int) {
	current := sl.head
	var previous *Node = nil

	for {
		if current == nil {
			break
		}

		if current.value == target {
			if previous != nil {
				previous.next = current.next
			} else {
				sl.head = current.next
			}

			return

		} else {
			previous = current
			current = current.next
		}
	}
}

func (sl *SingleLinkedList) RemoveElements(target int) {
	var current = sl.head
	var previous *Node = nil

	for {
		if current == nil {
			break
		}

		if current.value == target {
			if previous != nil {
				previous.next = current.next
			} else {
				sl.head = current.next
			}

			current = current.next
		} else {
			previous = current
			current = current.next
		}
	}
}

func (sl *SingleLinkedList) ReverseIterative() {
	var previous *Node = nil
	current := sl.head

	for {
		if current == nil {
			break
		}

		nextNode := current.next
		current.next = previous
		previous = current
		current = nextNode
	}
	sl.head = previous
}

func (sl *SingleLinkedList) ReverseRecursive() {
	sl.head = sl.reverseRecursive(sl.head, nil)
}

func (sl *SingleLinkedList) reverseRecursive(current *Node, previous *Node) *Node {
	if current == nil {
		return previous
	}

	nextNode := current.next
	current.next = previous
	previous = current
	current = nextNode

	return sl.reverseRecursive(current, previous)

}

func (sl *SingleLinkedList) ReverseEven() {
	sl.head = sl.reverseEven(sl.head, nil)
}

func (sl *SingleLinkedList) reverseEven(head *Node, previous *Node) *Node {
	if head == nil {
		return head
	}

	current := head
	for (current != nil) && (current.value%2 == 0) {
		nextNode := current.next
		current.next = previous
		previous = current
		current = nextNode
	}

	if current != head {
		head.next = current
		sl.reverseEven(current, nil)
		return previous
	} else {
		head.next = sl.reverseEven(head.next, head)
		return head
	}
}

func (sl *SingleLinkedList) SortMyself() {
	var current = sl.head

	for {
		if current == nil {
			break
		}

		nextNode := current.next
		for {
			if nextNode == nil {
				break
			}
			if current.value > nextNode.value {
				current.value, nextNode.value = nextNode.value, current.value
			}
			nextNode = nextNode.next
		}
		current = current.next
	}
}

func (sl *SingleLinkedList) MiddleNode() *Node {
	slow := sl.head
	fast := sl.head

	for (fast != nil) && (fast.next != nil) {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}

func ReverseSlice(input []int) []int {
	inputLen := len(input)
	output := make([]int, inputLen)

	for i, n := range input {
		j := inputLen - i - 1

		output[j] = n
	}

	return output
}

func (sl *SingleLinkedList) IsPalindrome() bool {
	var nodeValues []int
	current := sl.head
	for current != nil {
		nodeValues = append(nodeValues, current.value)
		current = current.next
	}
	reversedNodeValues := ReverseSlice(nodeValues)

	return reflect.DeepEqual(nodeValues, reversedNodeValues)
}

func (sl *SingleLinkedList) MergeTwoLinkedList(head *Node) {
	current1 := sl.head
	current2 := head

	result := NewSingleLinkedList()

	for (current1 != nil) && (current2 != nil) {
		if current1.value < current2.value {
			result.Append(current1.value)
			current1 = current1.next
		} else {
			result.Append(current2.value)
			current2 = current2.next
		}
	}

	for current1 != nil {
		result.Append(current1.value)
		current1 = current1.next
	}

	for current2 != nil {
		result.Append(current2.value)
		current2 = current2.next
	}

	sl.head = result.head
}

func (sl *SingleLinkedList) RemoveNthNodeFromEnd(n int) {
	size := sl.Size()
	if size == n {
		sl.head = sl.head.next
		return
	} else if n < 0 {
		return
	} else if size < n {
		return
	}

	nthFromHead := size - n
	var previous *Node = nil
	current := sl.head
	for i := 0; i < nthFromHead; i++ {
		previous = current
		current = current.next
	}

	previous.next = current.next
}

func (sl *SingleLinkedList) RemoveNthNodeFromHead(n int) {
	size := sl.Size()

	if n == 1 {
		sl.head = sl.head.next
		return
	} else if (size < n) || (sl.head == nil) || (n <= 0) {
		return
	}

	var previous *Node = nil
	current := sl.head
	for i := 1; i < n; i++ {
		previous = current
		current = current.next
	}

	previous.next = current.next
}
