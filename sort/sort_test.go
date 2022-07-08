package sort

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestBubbleSort(t *testing.T) {
	// Init
	randomArray := make([]int, 10)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		randomArray[i] = rand.Intn(100)
	}

	// Execution
	got := BubbleSort(randomArray)

	// Test
	previous := 0
	for _, value := range got {
		if previous > value {
			t.Fail()
		}
		previous = value
	}
}

func TestSelectionSort(t *testing.T) {
	// Init
	randomArray := make([]int, 10)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		randomArray[i] = rand.Intn(100)
	}

	// Execution
	got := SelectionSort(randomArray)

	// Test
	previous := -1 * math.MaxInt
	for _, value := range got {
		if previous > value {
			t.Fail()
		}
		previous = value
	}
}

func TestInsertionSort(t *testing.T) {
	// Init
	randomArray := make([]int, 10)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		randomArray[i] = rand.Intn(100)
	}

	// Execution
	got := InsertionSort(randomArray)

	// Test
	previous := -1 * math.MaxInt
	for _, value := range got {
		if previous > value {
			t.Fail()
		}
		previous = value
	}
}

func TestQuickSort(t *testing.T) {
	// Init
	randomArray := make([]int, 10)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		randomArray[i] = rand.Intn(100)
	}

	// Execution
	got := QuickSort(randomArray)

	// Test
	previous := -1 * math.MaxInt
	for _, value := range got {
		if previous > value {
			t.Fail()
		}
		previous = value
	}
}

func TestMergeSort(t *testing.T) {
	// Init
	randomArray := make([]int, 10)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		randomArray[i] = rand.Intn(100)
	}

	// Execution
	got := MergeSort(randomArray)

	// Test
	previous := -1 * math.MaxInt
	for _, value := range got {
		if previous > value {
			t.Fail()
		}
		previous = value
	}
}
