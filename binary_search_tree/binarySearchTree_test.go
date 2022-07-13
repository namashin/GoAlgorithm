package binary_search_tree

import (
	"fmt"
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

func TestBinarySearchTree_Search(t *testing.T) {
	// Init
	bst := NewBinarySearchTree()

	bst.Insert(6)
	bst.Insert(4)
	bst.Insert(9)
	bst.Insert(7)
	bst.Insert(2)
	bst.Insert(10)

	// Execution and Test
	assert.True(t, bst.Search(6))
	assert.True(t, bst.Search(2))
	assert.False(t, bst.Search(3))

	// Watch values
	bst.Inorder()
}

// 先頭を削除する場合のテスト
func TestBinarySearchTree_Remove(t *testing.T) {
	// Init
	bst := NewBinarySearchTree()

	bst.Insert(6)
	bst.Insert(4)
	bst.Insert(9)
	bst.Insert(7)
	bst.Insert(2)
	bst.Insert(10)

	// Execution
	bst.Remove(6)

	// Test
	assert.Equal(t, 7, bst.root.value)
	assert.Equal(t, 9, bst.root.right.value)
	assert.Equal(t, 4, bst.root.left.value)
	assert.Equal(t, 2, bst.root.left.left.value)
	assert.Equal(t, 10, bst.root.right.right.value)

	// Watch values
	bst.Inorder()
}

// 子コードが左右あるノードの削除
func TestBinarySearchTree_Remove2(t *testing.T) {
	// Init
	bst := NewBinarySearchTree()

	bst.Insert(6)
	bst.Insert(4)
	bst.Insert(9)
	bst.Insert(7)
	bst.Insert(2)
	bst.Insert(10)

	// Execution
	bst.Remove(9)

	// Test
	assert.Equal(t, 6, bst.root.value)
	assert.Equal(t, 10, bst.root.right.value)
	assert.Equal(t, 4, bst.root.left.value)
	assert.Equal(t, 2, bst.root.left.left.value)
	assert.Equal(t, 7, bst.root.right.left.value)

	// Watch values
	bst.Inorder()
}

// 片方子ノードがあるノードを削除する場合
func TestBinarySearchTree_Remove3(t *testing.T) {
	// Init
	bst := NewBinarySearchTree()

	bst.Insert(6)
	bst.Insert(4)
	bst.Insert(9)
	bst.Insert(7)
	bst.Insert(2)
	bst.Insert(10)

	// Execution
	bst.Remove(4)

	// Test
	assert.Equal(t, 6, bst.root.value)
	assert.Equal(t, 9, bst.root.right.value)
	assert.Equal(t, 2, bst.root.left.value)
	assert.Equal(t, 7, bst.root.right.left.value)
	assert.Equal(t, 10, bst.root.right.right.value)

	// Watch values
	bst.Inorder()
}

// 子ノード無しのノード削除する場合
func TestBinarySearchTree_Remove4t(t *testing.T) {
	// Init
	bst := NewBinarySearchTree()

	bst.Insert(6)
	bst.Insert(4)
	bst.Insert(9)
	bst.Insert(7)
	bst.Insert(2)
	bst.Insert(10)

	// Execution
	bst.Remove(7)

	// Test
	assert.Equal(t, 6, bst.root.value)
	assert.Equal(t, 9, bst.root.right.value)
	assert.Equal(t, 4, bst.root.left.value)
	assert.Equal(t, 2, bst.root.left.left.value)
	assert.Equal(t, 10, bst.root.right.right.value)

	// Watch values
	bst.Inorder()
}

func TestBinarySearchTree_MaxDepth(t *testing.T) {
	// Init
	bst := NewBinarySearchTree()

	bst.Insert(6)
	bst.Insert(4)
	bst.Insert(9)
	bst.Insert(7)
	bst.Insert(2)
	bst.Insert(10)

	// Execution and Test
	assert.Equal(t, 3, bst.MaxDepth())

	// Watch values
	bst.Inorder()
}

func TestBinarySearchTree_PathSum(t *testing.T) {
	// Init
	bst := NewBinarySearchTree()

	bst.Insert(6)
	bst.Insert(4)
	bst.Insert(9)
	bst.Insert(7)
	bst.Insert(2)
	bst.Insert(10)

	// Execution and Test
	assert.True(t, bst.PathSum(12))
	assert.True(t, bst.PathSum(22))
	assert.True(t, bst.PathSum(25))

	assert.False(t, bst.PathSum(10))
	assert.False(t, bst.PathSum(38))

	// Watch values
	bst.Inorder()

}
func TestBinarySearchTree_InvertTree(t *testing.T) {
	// Init
	bst := NewBinarySearchTree()

	bst.Insert(6)
	bst.Insert(4)
	bst.Insert(9)
	bst.Insert(7)
	bst.Insert(2)
	bst.Insert(10)

	// Execution
	bst.InvertTree()

	// Test
	assert.Equal(t, 6, bst.root.value)
	assert.Equal(t, 4, bst.root.right.value)
	assert.Equal(t, 9, bst.root.left.value)
	assert.Equal(t, 10, bst.root.left.left.value)
	assert.Equal(t, 2, bst.root.right.right.value)
	assert.Equal(t, 7, bst.root.left.right.value)

	// Watch values
	bst.Inorder()
}

func TestBinarySearchTree_SumOfRightLeaves(t *testing.T) {
	// Init
	bst := NewBinarySearchTree()

	bst.Insert(6)
	bst.Insert(4)
	bst.Insert(9)
	bst.Insert(7)
	bst.Insert(2)
	bst.Insert(10)

	// Execution and Test
	assert.Equal(t, 10, bst.SumOfRightLeaves())

}

func TestBinarySearchTree_SumOfLeftLeaves(t *testing.T) {
	// Init
	bst := NewBinarySearchTree()

	bst.Insert(6)
	bst.Insert(4)
	bst.Insert(9)
	bst.Insert(7)
	bst.Insert(2)
	bst.Insert(10)

	// Execution and Test
	assert.Equal(t, 9, bst.SumOfLeftLeaves())
}

func TestBinarySearchTree_AllElements(t *testing.T) {
	// Init
	bst := NewBinarySearchTree()

	bst.Insert(6)
	bst.Insert(4)
	bst.Insert(9)
	bst.Insert(7)
	bst.Insert(2)
	bst.Insert(10)

	// Execution
	got := bst.AllElements()
	want := []int{6, 4, 2, 9, 7, 10}
	fmt.Println(got)

	// Test
	for i, value := range got {
		if value != want[i] {
			t.Fail()
		}
	}

	if len(got) != 6 {
		t.Fail()
	}
}

func TestBinarySearchTree_LevelOrder(t *testing.T) {
	// Init
	bst := NewBinarySearchTree()

	bst.Insert(6)
	bst.Insert(4)
	bst.Insert(9)
	bst.Insert(7)
	bst.Insert(2)
	bst.Insert(10)
	bst.Insert(1)
	bst.Insert(3)
	bst.Insert(8)
	bst.Insert(13)
	bst.Insert(15)

	// Execution
	got := bst.LevelOrder()
	want := [][]int{{6}, {4, 9}, {2, 7, 10}, {1, 3, 8, 13}, {15}}

	// Test
	assert.Equal(t, want, got)

}
