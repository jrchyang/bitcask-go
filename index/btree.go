package index

import (
	"sync"

	"github.com/google/btree"
	"github.com/jrchyang/bitcask-go/data"
)

// BTree 索引，主要封装了 google 的 btree 库
// https://github.com/google/btree
type BTree struct {
	lock *sync.RWMutex
	tree *btree.BTree
}

// NewBTree 初始化 BTree 索引结构题
func NewBTree() *BTree {
	return &BTree{
		lock: new(sync.RWMutex),
		tree: btree.New(32),
	}
}

func (bt *BTree) Put(key []byte, pos *data.LogRecordPos) bool {
	it := &Item{key: key, pos: pos}
	bt.lock.Lock()
	bt.tree.ReplaceOrInsert(it)
	bt.lock.Unlock()
	return true
}

func (bt *BTree) Get(key []byte) *data.LogRecordPos {
	it := &Item{key: key}
	btreeItem := bt.tree.Get(it)

	if btreeItem == nil {
		return nil
	}
	return btreeItem.(*Item).pos
}

func (bt *BTree) Delete(key []byte) bool {
	it := &Item{key: key}
	bt.lock.Lock()
	oldItem := bt.tree.Delete(it)
	bt.lock.Unlock()
	return oldItem != nil
}
