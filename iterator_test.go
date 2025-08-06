package bitcaskgo

import (
	"os"
	"testing"

	"github.com/jrchyang/bitcask-go/utils"
	"github.com/stretchr/testify/assert"
)

func TestDB_NewIterator(t *testing.T) {
	opts := DefaultOptions
	dir, _ := os.MkdirTemp("", "bitcask-go")
	opts.DirPath = dir
	db, err := Open(opts)
	defer destroyDB(db)
	assert.Nil(t, err)
	assert.NotNil(t, db)

	it := db.NewIterator(DefaultIteratorOptions)
	assert.Equal(t, false, it.Valid())
}

func TestDB_Iterator_One_Value(t *testing.T) {
	opts := DefaultOptions
	dir, _ := os.MkdirTemp("", "bitcask-go")
	opts.DirPath = dir
	db, err := Open(opts)
	defer destroyDB(db)
	assert.Nil(t, err)
	assert.NotNil(t, db)

	err = db.Put(utils.GetTestKey(10), utils.RandomValue(10))
	assert.Nil(t, err)

	it := db.NewIterator(DefaultIteratorOptions)
	assert.NotNil(t, it)
	assert.Equal(t, true, it.Valid())
	assert.Equal(t, utils.GetTestKey(10), it.Key())
	_, err = it.Value()
	assert.Nil(t, err)
}

func TestDB_Iterator_Multi_Value(t *testing.T) {
	opts := DefaultOptions
	dir, _ := os.MkdirTemp("", "bitcask-go")
	opts.DirPath = dir
	db, err := Open(opts)
	defer destroyDB(db)
	assert.Nil(t, err)
	assert.NotNil(t, db)

	err = db.Put([]byte("aabc"), utils.RandomValue(10))
	assert.Nil(t, err)
	err = db.Put([]byte("bbcd"), utils.RandomValue(10))
	assert.Nil(t, err)
	err = db.Put([]byte("ccde"), utils.RandomValue(10))
	assert.Nil(t, err)
	err = db.Put([]byte("defg"), utils.RandomValue(10))
	assert.Nil(t, err)
	err = db.Put([]byte("abcd"), utils.RandomValue(10))
	assert.Nil(t, err)
	err = db.Put([]byte("nmlk"), utils.RandomValue(10))
	assert.Nil(t, err)

	// 正向迭代
	it1 := db.NewIterator(DefaultIteratorOptions)
	for it1.Rewind(); it1.Valid(); it1.Next() {
		assert.NotNil(t, it1.Key())
	}
	it1.Rewind()
	for it1.Seek([]byte("c")); it1.Valid(); it1.Next() {
		assert.NotNil(t, it1.Key())
	}

	// 反向迭代
	iterOptions := DefaultIteratorOptions
	iterOptions.Reverse = true
	it2 := db.NewIterator(iterOptions)
	for it2.Rewind(); it2.Valid(); it2.Next() {
		assert.NotNil(t, it2.Key())
	}
	it2.Rewind()
	for it2.Seek([]byte("z")); it2.Valid(); it2.Next() {
		// t.Log(string(it2.Key()))
		assert.NotNil(t, it2.Key())
	}

	// 指定 prefix
	iterOptions = DefaultIteratorOptions
	iterOptions.Prefix = []byte("aa")
	it3 := db.NewIterator(iterOptions)
	for it3.Rewind(); it3.Valid(); it3.Next() {
		t.Log(string(it3.Key()))
	}
}
