package linkedlist

import "fmt"

// DoublyNode 双方向リンクリスト用ノード
type DoublyNode struct {
	value int
	next  *DoublyNode
	prev  *DoublyNode
}

type DoublyLinkedList struct {
	head *DoublyNode
}

func NewDoublyLInkedList() *DoublyLinkedList {
	return &DoublyLinkedList{
		head: nil,
	}
}

func (dl *DoublyLinkedList) Append(appendValue int) {
	var newNode = &DoublyNode{value: appendValue, next: nil, prev: nil}
	if dl.head == nil {
		dl.head = newNode
		return
	}

	current := dl.head
	for {
		if current.next == nil {
			break
		}

		current = current.next
	}
	current.next = newNode
	newNode.prev = current
}

func (dl *DoublyLinkedList) Print() {
	current := dl.head
	for {
		if current == nil {
			break
		}

		fmt.Println(current.value)
		current = current.next
	}
}

func (dl *DoublyLinkedList) Insert(insertValue int) {
	var newNode = &DoublyNode{value: insertValue, prev: nil, next: nil}
	if dl.head == nil {
		dl.head = newNode
		return
	}

	dl.head.prev = newNode
	newNode.next = dl.head
	dl.head = newNode
}

func (dl DoublyLinkedList) Size() int {
	current := dl.head
	size := 0
	for {
		if current == nil {
			break
		}
		current = current.next
		size += 1
	}
	return size
}

func (dl *DoublyLinkedList) ReverseIterative() {
	current := dl.head
	var previous *DoublyNode = nil

	for {
		if current == nil {
			break
		}

		nextNode := current.next
		current.next = previous
		previous = current
		current = nextNode
	}

	dl.head = previous
}

func (dl *DoublyLinkedList) ReverseRecursive() {
	dl.head = dl.reverseRecursive(dl.head, nil)
}

func (dl *DoublyLinkedList) reverseRecursive(current *DoublyNode, previous *DoublyNode) *DoublyNode {
	if current == nil {
		return previous
	}

	nextNode := current.next
	current.next = previous
	previous = current
	current = nextNode

	return dl.reverseRecursive(current, previous)

}
