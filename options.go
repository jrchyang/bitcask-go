package bitcaskgo

type Options struct {
	DirPath      string      // 数据库目录
	DataFileSize int64       // 数据文件的大小
	SyncWrite    bool        // 每次写数据是否持久化
	IndexType    IndexerType // 内存索引类型
}

type IndexerType = int8

const (
	BTree IndexerType = iota + 1
	ART
)
