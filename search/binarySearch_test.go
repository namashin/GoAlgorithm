package search

import "testing"

func TestBinarySearch(t *testing.T) {
	// Init
	sortedNumbers := []int{1, 2, 6, 8, 9, 10, 22}

	// Execution
	want := 6
	targetIndex := BinarySearch(sortedNumbers, want)

	// Test
	if targetIndex != 2 {
		t.Fail()
	}

	// Execution
	want = 100
	targetIndex = BinarySearch(sortedNumbers, want)

	// Test
	if targetIndex != -1 {
		t.Fail()
	}
}

func TestBinarySearchRecursive(t *testing.T) {
	// Init
	sortedNumbers := []int{1, 2, 6, 8, 9, 10, 22}

	// Execution
	want := 6
	targetIndex := BinarySearchRecursive(sortedNumbers, want, 0, len(sortedNumbers)-1)

	// Test
	if targetIndex != 2 {
		t.Fail()
	}

	// Execution
	want = 100
	targetIndex = BinarySearchRecursive(sortedNumbers, want, 0, len(sortedNumbers)-1)

	// Test
	if targetIndex != -1 {
		t.Fail()
	}
}
