package cache

import (
	"github.com/syndtr/goleveldb/leveldb"
	"strings"
)

// leveldb 缓存接口实现
type LeveldbCache struct {
	db *leveldb.DB
}

func NewLeveldb(path string) *LeveldbCache {
	db, err := leveldb.OpenFile(strings.TrimRight(path, "/") + "/leveldb", nil)
	if err != nil {
		panic(err)
	}
	return &LeveldbCache{
		db: db,
	}
}

func (this LeveldbCache) Get(key string) (string, bool) {
	val, err := this.db.Get([]byte(key), nil)
	if err == leveldb.ErrNotFound {
		return "", true
	} else if err != nil {
		return "", false
	}
	return string(val), true
}

func (this LeveldbCache) Put(key string, val string) error {
	return this.db.Put([]byte(key), []byte(val), nil)
}
