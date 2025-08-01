package fio

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func deleteFile(name string) {
	if err := os.RemoveAll(name); err != nil {
		panic(err)
	}
}

func TestNewFileIOManager(t *testing.T) {
	filename := filepath.Join("/tmp/bitcask", "a.data")
	fio, err := NewFileIOManager(filename)
	defer deleteFile(filename)

	assert.Nil(t, err)
	assert.NotNil(t, fio)
}

func TestFileIO_Write(t *testing.T) {
	filename := filepath.Join("/tmp/bitcask", "a.data")
	fio, err := NewFileIOManager(filename)
	defer deleteFile(filename)
	assert.Nil(t, err)
	assert.NotNil(t, fio)

	n, err := fio.Write([]byte(""))
	assert.Equal(t, 0, n)
	assert.Nil(t, err)
	n, err = fio.Write([]byte("bitcask kv"))
	assert.Equal(t, 10, n)
	assert.Nil(t, err)
	n, err = fio.Write([]byte("stroage"))
	assert.Equal(t, 7, n)
	assert.Nil(t, err)
}

func TestFileIO_Read(t *testing.T) {
	filename := filepath.Join("/tmp/bitcask", "a.data")
	fio, err := NewFileIOManager(filename)
	defer deleteFile(filename)
	assert.Nil(t, err)
	assert.NotNil(t, fio)

	_, err = fio.Write([]byte("key-a"))
	assert.Nil(t, err)
	_, err = fio.Write([]byte("key-b"))
	assert.Nil(t, err)

	b1 := make([]byte, 5)
	n, _ := fio.Read(b1, 0)
	assert.Equal(t, 5, n)
	assert.Equal(t, []byte("key-a"), b1)

	b2 := make([]byte, 5)
	n, _ = fio.Read(b2, 5)
	assert.Equal(t, 5, n)
	assert.Equal(t, []byte("key-b"), b2)
}
