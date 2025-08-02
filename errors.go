package bitcaskgo

import "errors"

var (
	ErrKeyIsEmpty        = errors.New("the key is empty")
	ErrKeyNotFound       = errors.New("key not found in database")
	ErrIndexUpdateFailed = errors.New("failed to update index")
	ErrDataFileNotFound  = errors.New("datafile not found in database")
)
