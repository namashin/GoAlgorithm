package hashtable

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewHashTable(t *testing.T) {
	// Init and Execution
	h := NewHashTable(10)

	// Execution
	key := h.Hash("password x")
	key2 := h.Hash("password x")

	key3 := h.Hash("password y")
	key4 := h.Hash("password y")

	// Test
	assert.Equal(t, key, key2)
	assert.Equal(t, key3, key4)
}

func TestHashTable_Add(t *testing.T) {
	// Init and Execution
	h := NewHashTable(10)

	// Execution
	h.Add("key1", 1)
	h.Add("key1", 2)
	h.Add("key1", 1)
	h.Add("key2", 2)
	h.Add("key3", 3)
	h.Add("key4", 4)
	h.Add("key1", 100)
	h.Add("key4", 200)

	fmt.Println(h.table)

}

func TestHashTable_Delete(t *testing.T) {
	// Init and Execution
	h := NewHashTable(10)

	h.Add("key2", 2)
	h.Add("key3", 3)
	h.Add("key1", 100)
	h.Add("key4", 200)
	fmt.Println("before", h.table)
	// Execution
	deleteKey := "key1"
	h.Delete(deleteKey)

	// Test
	index := h.Hash(deleteKey)
	for _, pair := range h.table[index] {
		if pair.key == deleteKey {
			t.Fail()
		}
	}

	// Watch values
	fmt.Println("after ", h.table)
}

func TestHashTable_Get(t *testing.T) {
	// Init and Execution
	h := NewHashTable(10)

	h.Add("key2", 2)
	h.Add("key3", 3)
	h.Add("key1", 100)
	h.Add("key4", 200)

	// Execution and Test
	assert.Equal(t, 2, h.Get("key2"))
	assert.Equal(t, 3, h.Get("key3"))
	assert.Equal(t, 100, h.Get("key1"))
	assert.Equal(t, 200, h.Get("key4"))
	assert.Equal(t, -1, h.Get("there is no key"))

}
