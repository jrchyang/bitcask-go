package fio

const DataFilePerm = 0644

// 抽象 IO 管理器接口，可以接入不同的 IO 类型，目前支持标准文件 IO
type IOManager interface {
	// 从文件的给定位置读取对应的数据
	Read([]byte, int64) (int, error)
	// 写入字节数组到文件中
	Write([]byte) (int, error)
	// 持久化数据
	Sync() error
	// 关闭文件
	Close() error
}
