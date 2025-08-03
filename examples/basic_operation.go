package main

import (
	"fmt"

	bitcaskgo "github.com/jrchyang/bitcask-go"
)

func main() {
	opts := bitcaskgo.DefaultOptions
	opts.DirPath = "/tmp/bitcask"

	db, err := bitcaskgo.Open(opts)
	if err != nil {
		panic(err)
	}

	err = db.Put([]byte("name"), []byte("bitcask"))
	if err != nil {
		panic(err)
	}

	val, err := db.Get([]byte("name"))
	if err != nil {
		panic(err)
	}
	fmt.Println("val = ", string(val))

	err = db.Delete([]byte("name"))
	if err != nil {
		panic(err)
	}
}
