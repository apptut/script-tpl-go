package config

import (
	"github.com/BurntSushi/toml"
)

var Mysql mysql
var Mongo mongo
var Cache cache

type mysql struct {
	Host     string
	User     string
	Pwd      string `toml:"password"`
	Port     string
	Database string
}

type mongo struct {
	Uri    string
	DbName string `toml:"dbname"`
}

type config struct {
	Mysql mysql
	Mongo mongo
	Cache cache
}

type cache struct {
	Driver string
	Path   string
}

// 初始化加载配置文件
func ConfigInit(rootPath string) {
	var conf config
	if _, err := toml.DecodeFile(rootPath+"/config/config.toml", &conf); err != nil {
		panic(err)
	}

	Mysql = conf.Mysql
	Mongo = conf.Mongo
	Cache = conf.Cache
}
