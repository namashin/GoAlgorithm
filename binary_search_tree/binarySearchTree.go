package binary_search_tree

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

type binarySearchTree struct {
	root *Node
}

func NewBinarySearchTree() *binarySearchTree {
	return &binarySearchTree{
		root: nil,
	}
}

func (bst *binarySearchTree) Insert(InsertValue int) {
	bst.root = bst.insert(bst.root, InsertValue)
}

func (bst *binarySearchTree) insert(root *Node, insertValue int) *Node {
	newNode := &Node{value: insertValue, left: nil, right: nil}
	if root == nil {
		return newNode
	}

	if root.value < insertValue {
		root.right = bst.insert(root.right, insertValue)
	} else if root.value > insertValue {
		root.left = bst.insert(root.left, insertValue)
	}
	return root
}

func (bst *binarySearchTree) Inorder() {
	bst.inorder(bst.root)
}

func (bst *binarySearchTree) inorder(node *Node) {
	if node != nil {
		bst.inorder(node.left)
		fmt.Println(node.value)
		bst.inorder(node.right)
	}
}
