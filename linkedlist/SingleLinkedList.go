package linkedlist

import (
	"fmt"
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
	var newNode *Node = &Node{value: appendValue, next: nil}

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

func (sl *SingleLinkedList) Insert(insertValue int) {
	var newNode *Node = &Node{value: insertValue, next: nil}

	newNode.next = sl.head
	sl.head = newNode
}

func (sl *SingleLinkedList) RemoveElement(target int) {
	var current *Node = sl.head
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
	var current *Node = sl.head
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
	var current *Node = sl.head

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

func (sl *SingleLinkedList) SortMyself() {
	var current *Node = sl.head

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
