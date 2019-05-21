# quick-docker-api

封装自官方docker api sdk。

为ctf平台搭建中的比赛容器自动下发而使用。

可以使用该库快速开发“容器下发”功能。

不需要额外引入其他依赖，所有内部依赖已打包在vendor目录。

## User Guide

### Prerequisites

- Golang Version >= 1.11.2

### Installation

```sh
$ go get github.com/MashiroC/quick-docker-api
```


## Release History

* 1.0.0

    - docker容器创建、开启、关闭等基础功能