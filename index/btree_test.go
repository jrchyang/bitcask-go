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
