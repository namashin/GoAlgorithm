package binary_search_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinarySearchTree_Insert(t *testing.T) {
	// Init
	bst := NewBinarySearchTree()

	// Execution
	bst.Insert(6)
	bst.Insert(4)
	bst.Insert(9)
	bst.Insert(7)
	bst.Insert(2)
	bst.Insert(10)

	// Test
	assert.Equal(t, 6, bst.root.value)
	assert.Equal(t, 4, bst.root.left.value)
	assert.Equal(t, 9, bst.root.right.value)
	assert.Equal(t, 7, bst.root.right.left.value)
	assert.Equal(t, 2, bst.root.left.left.value)
	assert.Equal(t, 10, bst.root.right.right.value)

	// Watch values
	bst.Inorder()
}
