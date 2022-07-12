package binary_search_tree

import (
	"fmt"
	"math"
)

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
	bst.root = insert(bst.root, InsertValue)
}

func insert(root *Node, insertValue int) *Node {
	newNode := &Node{value: insertValue, left: nil, right: nil}
	if root == nil {
		return newNode
	}

	if root.value < insertValue {
		root.right = insert(root.right, insertValue)
	} else if root.value > insertValue {
		root.left = insert(root.left, insertValue)
	}
	return root
}

func (bst *binarySearchTree) Inorder() {
	inorder(bst.root)
}

func inorder(node *Node) {
	if node != nil {
		inorder(node.left)
		fmt.Println(node.value)
		inorder(node.right)
	}
}

func (bst *binarySearchTree) Search(TargetValue int) bool {
	return search(bst.root, TargetValue)
}

func search(node *Node, targetValue int) bool {
	if node == nil {
		return false
	}

	if node.value == targetValue {
		return true
	} else if node.value < targetValue {
		return search(node.right, targetValue)
	} else {
		return search(node.left, targetValue)
	}
}

func getMinimumNode(node *Node) *Node {
	for {
		if node.left == nil {
			break
		}

		node = node.left
	}
	return node
}

func (bst *binarySearchTree) Remove(RemoveValue int) {
	bst.root = remove(bst.root, RemoveValue)
}

func remove(node *Node, removeValue int) *Node {
	if node == nil {
		return node
	}

	if node.value < removeValue {
		node.right = remove(node.right, removeValue)
	} else if node.value > removeValue {
		node.left = remove(node.left, removeValue)
	} else {
		if node.right == nil {
			return node.left
		} else if node.left == nil {
			return node.right
		}

		temp := getMinimumNode(node.right)
		node.value = temp.value
		node.right = remove(node.right, temp.value)
	}
	return node
}

func (bst *binarySearchTree) MaxDepth() int {
	return int(maxDepth(bst.root))
}

func maxDepth(node *Node) float64 {
	if node == nil {
		return 0
	}

	left := maxDepth(node.left)
	right := maxDepth(node.right)

	return math.Max(left, right) + 1

}

func (bst *binarySearchTree) PathSum(TargetSum int) bool {
	return pathSum(bst.root, TargetSum)
}

func pathSum(node *Node, targetSum int) bool {
	if node == nil {
		return false
	}

	if (node.left == nil) && (node.right == nil) && (node.value == targetSum) {
		return true
	}

	targetSum -= node.value

	return pathSum(node.left, targetSum) || pathSum(node.right, targetSum)

}

func (bst *binarySearchTree) InvertTree() {
	invertTree(bst.root)
}

func invertTree(node *Node) *Node {
	if node != nil {
		node.left, node.right = invertTree(node.right), invertTree(node.left)
		return node
	}
	return node
}

func (bst *binarySearchTree) SumOfRightLeaves() int {
	return sumOfRightLeaves(bst.root)
}

func sumOfRightLeaves(node *Node) int {
	if node == nil {
		return 0
	}
	value := 0
	if (node.right != nil) && (node.right.right == nil) && (node.right.left == nil) {
		value += node.right.value
	}

	return value + sumOfRightLeaves(node.left) + sumOfRightLeaves(node.right)
}

func (bst *binarySearchTree) SumOfLeftLeaves() int {
	return sumOfLeftLeaves(bst.root)
}

func sumOfLeftLeaves(node *Node) int {
	if node == nil {
		return 0
	}

	value := 0
	if (node.left != nil) && (node.left.left == nil) && (node.left.right == nil) {
		value += node.left.value
	}

	return value + sumOfLeftLeaves(node.left) + sumOfLeftLeaves(node.right)

}

func (bst *binarySearchTree) AllElements() []int {
	elements := new([]int)
	allElements(bst.root, elements)

	return *elements
}

func allElements(node *Node, elements *[]int) {
	if node == nil {
		return
	}

	*elements = append(*elements, node.value)

	allElements(node.left, elements)
	allElements(node.right, elements)
}
