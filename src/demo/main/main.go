package main

import (
	"demo/config"
	"flag"
	"strings"
	"fmt"
	"os"
	"demo/task"
	"demo/cache"
)

func main() {
	// 项目根目录
	var rootPath string

	flag.StringVar(&rootPath, "rootPath", "", "project absolute root path")
	flag.Parse()

	rootPath = strings.TrimRight(rootPath, "/")
	if rootPath == ""{
		fmt.Println("missing flag path !")
		flag.PrintDefaults()
		os.Exit(0)
	}

	// 加载配置项
	config.ConfigInit(rootPath)

	// 加载缓存配置
	cache.Load()

	// 开始执行业务脚本
	task.SyncDataRun()

	fmt.Println("done !!!")
}
