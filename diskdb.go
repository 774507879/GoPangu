package diskdb

import (
	"io"
	"os"
	"sync"
)

type DiskDB struct {
	indexes map[string]int64 //内存中的索引信息
	dbFile  *DBFile          //数据文件
	dirPath string           //数据目录
	mu      sync.RWMutex
}

func Open(dirPath string) (*DiskDB, error) {
	if _, err := os.Stat(dirPath); os.IsExist(err) {
		if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
			return nil, err
		}
	}

	dbFile, err := NewDBFile(dirPath)
	if err != nil {
		return nil, err
	}
	db := &DiskDB{
		dbFile:  dbFile,
		indexes: make(map[string]int64),
		dirPath: dirPath,
	}

	// 加载索引
	db.loadIndexesFromFile()
	return db, nil
}

// 从文件当中加载索引
func (db *DiskDB) loadIndexesFromFile() {
	if db.dbFile == nil {
		return
	}

	var offset int64
	for {
		e, err := db.dbFile.Read(offset)
		if err != nil {
			// 读取完毕
			if err == io.EOF {
				break
			}
			return
		}

		// 设置索引状态
		db.indexes[string(e.Key)] = offset

		if e.Mark == DEL {
			// 删除内存中的 key
			delete(db.indexes, string(e.Key))
		}

		offset += e.GetSize()
	}
	return
}

func (db *DiskDB) Put(key []byte, value []byte) (err error) {
	if len(key) == 0 {
		return
	}
	db.mu.Lock()
	defer db.mu.Unlock()
	offset := db.dbFile.Offset

	entry := NewEntry(key, value, PUT)
	//追到文件数据当中
	err = db.dbFile.Write(entry)

	//写到内存
	db.indexes[string(key)] = offset
	return

}

func (db *DiskDB) Get(key []byte) (val []byte, err error) {
	if len(key) == 0 {
		return
	}

	db.mu.RLock()
	defer db.mu.RUnlock()
	// 从内存当中取出索引信息
	offset, ok := db.indexes[string(key)]
	// key 不存在
	if !ok {
		return
	}
	// 从磁盘中读取数据
	var e *Entry
	e, err = db.dbFile.Read(offset)
	if err != nil && err != io.EOF {
		return
	}
	if e != nil {
		val = e.Value
	}
	return

}

// Del 删除数据
func (db *DiskDB) Del(key []byte) (err error) {
	if len(key) == 0 {
		return
	}

	db.mu.Lock()
	defer db.mu.Unlock()
	// 从内存当中取出索引信息
	_, ok := db.indexes[string(key)]
	// key 不存在，忽略
	if !ok {
		return
	}

	// 封装成 Entry 并写入
	e := NewEntry(key, nil, DEL)
	err = db.dbFile.Write(e)
	if err != nil {
		return
	}

	// 删除内存中的 key
	delete(db.indexes, string(key))
	return
}
