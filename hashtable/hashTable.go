package hashtable

import (
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
)

type hashTable struct {
	table map[int][]keyValue
	size  int
}

type keyValue struct {
	key   string
	value int
}

func NewHashTable(size int) *hashTable {
	return &hashTable{
		table: make(map[int][]keyValue, size),
		size:  size,
	}
}

// Hash https://pkg.go.dev/hash
// 同じkeyに対して、同じintを返す
func (ht *hashTable) Hash(key string) int {
	// keyをエンコード
	encodedStr := hex.EncodeToString([]byte(key))
	// md5でハッシュ値作成
	first := md5.New()
	first.Write([]byte(encodedStr))

	// []byte -> int化
	data := binary.BigEndian.Uint32(first.Sum(nil))

	// ハッシュテーブルのサイズで割った余りを返す
	return int(data) % ht.size
}

func (ht *hashTable) Add(insertKey string, insertValue int) {
	index := ht.Hash(insertKey)

	for i, pair := range ht.table[index] {
		// 既に同じkeyで値入っていたら、削除
		if pair.key == insertKey {
			ht.table[index] = append(ht.table[index][:i], ht.table[index][i+1:]...)
		}
	}

	ht.table[index] = append(ht.table[index], keyValue{key: insertKey, value: insertValue})

}

func (ht *hashTable) Delete(key string) {
	index := ht.Hash(key)

	for i, pair := range ht.table[index] {
		if pair.key == key {
			ht.table[index] = append(ht.table[index][:i], ht.table[index][i+1:]...)
		}
	}
}

func (ht *hashTable) Get(key string) int {
	index := ht.Hash(key)

	for _, pair := range ht.table[index] {
		if pair.key == key {
			return pair.value
		}
	}
	return -1
}
