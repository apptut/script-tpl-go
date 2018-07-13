# script-tpl-go

脚本开发go版项目模板

之所以使用`go`语言开发脚本，主要是因为`go`的部署方便，只需要拷贝编译后的二进制文件，在某些场景下变得极为方便：

1. 运行环境不允许安装脚本运行的依赖库。
2. 不想对运行环境因安装扩展库造成的侵入性。


## 目录结构

```
├── config
│   ├── config-example.toml
│   └── config.toml -> config-example.toml
└── src
    ├── demo
    │   ├── cache
    │   ├── config
    │   ├── dep.sh
    │   ├── Gopkg.lock
    │   ├── Gopkg.toml
    │   ├── main
    │   ├── task
    │   └── vendor
    └── go.sh
```

## 使用方式

```sh
# 安装依赖库，是自己情况自行安装配置Gopkg.toml文件，例如需要用到的mysql mongo redis leaveldb等
./src/demo/dep.sh ensure

# config配置文件
cp config/config-example.toml config/config.toml

# 如果需要用到leveldb 请修改配置中的`cache.path`字段。

```
