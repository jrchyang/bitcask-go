package index

import (
	"testing"

	"github.com/jrchyang/bitcask-go/data"
	"github.com/stretchr/testify/assert"
)

func TestBtree_Put(t *testing.T) {
	bt := NewBTree()

	res := bt.Put(nil, &data.LogRecordPos{Fid: 1, Offset: 100})
	assert.True(t, res)

	res = bt.Put([]byte("a"), &data.LogRecordPos{Fid: 1, Offset: 2})
	assert.True(t, res)
}

func TestBtree_Get(t *testing.T) {
	bt := NewBTree()

	res := bt.Put(nil, &data.LogRecordPos{Fid: 1, Offset: 100})
	assert.True(t, res)
	pos := bt.Get(nil)
	assert.Equal(t, uint32(1), pos.Fid)
	assert.Equal(t, int64(100), pos.Offset)

	res = bt.Put([]byte("a"), &data.LogRecordPos{Fid: 1, Offset: 2})
	assert.True(t, res)
	res = bt.Put([]byte("a"), &data.LogRecordPos{Fid: 1, Offset: 3})
	assert.True(t, res)
	pos = bt.Get([]byte("a"))
	assert.Equal(t, uint32(1), pos.Fid)
	assert.Equal(t, int64(3), pos.Offset)
}

func TestBtree_Delete(t *testing.T) {
	bt := NewBTree()

	res := bt.Put(nil, &data.LogRecordPos{Fid: 1, Offset: 100})
	assert.True(t, res)
	res = bt.Delete(nil)
	assert.True(t, res)

	res = bt.Put([]byte("aaa"), &data.LogRecordPos{Fid: 1, Offset: 2})
	assert.True(t, res)
	res = bt.Delete([]byte("aaa"))
	assert.True(t, res)
}

func TestBtree_Iterator(t *testing.T) {
	bt1 := NewBTree()

	// 1. BTree 为空的场景
	it1 := bt1.Iterator(false)
	assert.Equal(t, false, it1.Valid())

	// 2. BTree 有数据的场景
	bt1.Put([]byte("ccde"), &data.LogRecordPos{Fid: 1, Offset: 10})
	it2 := bt1.Iterator(false)
	assert.Equal(t, true, it2.Valid())
	assert.NotNil(t, it2.Key())
	assert.NotNil(t, it2.Value())
	it2.Next()
	assert.Equal(t, false, it1.Valid())

	// 3. 有多条数据
	bt1.Put([]byte("acee"), &data.LogRecordPos{Fid: 1, Offset: 10})
	bt1.Put([]byte("eede"), &data.LogRecordPos{Fid: 1, Offset: 10})
	bt1.Put([]byte("bbcd"), &data.LogRecordPos{Fid: 1, Offset: 10})
	it3 := bt1.Iterator(false)
	for it3.Rewind(); it3.Valid(); it3.Next() {
		assert.NotNil(t, it3.Key())
	}
	it4 := bt1.Iterator(true)
	for it4.Rewind(); it4.Valid(); it4.Next() {
		assert.NotNil(t, it4.Key())
	}

	// 4. Seek
	it5 := bt1.Iterator(false)
	for it5.Seek([]byte("cc")); it5.Valid(); it5.Next() {
		assert.NotNil(t, it5.Key())
	}

	// 5. 反向 Seek
	it6 := bt1.Iterator(true)
	for it6.Seek([]byte("cc")); it6.Valid(); it6.Next() {
		assert.NotNil(t, it6.Key())
	}
}
