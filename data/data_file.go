package data

import "github.com/jrchyang/bitcask-go/fio"

const DataFileNameSuffix = ".data"

// 数据文件
type DataFile struct {
	FileId      uint32        // 文件 ID
	WriteOffset int64         // 文件写到了哪个位置
	IOManager   fio.IOManager // IO 读写管理
}

// 打开新的数据文件
func OpenDataFile(dirPath string, fileId uint32) (*DataFile, error) {
	return nil, nil
}

func (df *DataFile) Sync() error {
	return nil
}

func (df *DataFile) Write(buf []byte) error {
	return nil
}

func (df *DataFile) ReadLogRecord(offset int64) (*LogRecord, int64, error) {
	return nil, 0, nil
}
