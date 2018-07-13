package cache

import (
	"demo/config"
)

var Manager Cache

type Cache interface {
	Get(key string) (string, bool)
	Put(key string, value string) error
}

func Load() {
	switch config.Cache.Driver {
	case "leveldb":
		Manager = NewLeveldb(config.Cache.Path)
	}
}
