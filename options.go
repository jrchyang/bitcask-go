package bitcaskgo

import "os"

type Options struct {
	DirPath      string      // 数据库目录
	DataFileSize int64       // 数据文件的大小
	SyncWrite    bool        // 每次写数据是否持久化
	IndexType    IndexerType // 内存索引类型
}

// 迭代器配置项
type IteratorOptions struct {
	Prefix  []byte // 前缀
	Reverse bool   // 是否反向遍历，默认 false 是整箱
}

type IndexerType = int8

const (
	BTree IndexerType = iota + 1
	ART
)

var DefaultOptions = Options{
	DirPath:      os.TempDir(),
	DataFileSize: 256 * 1024 * 1024,
	SyncWrite:    false,
	IndexType:    BTree,
}

var DefaultIteratorOptions = IteratorOptions{
	Prefix:  nil,
	Reverse: false,
}
