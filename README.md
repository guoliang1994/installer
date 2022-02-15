# Introduction
installer 将帮助你在各种操作系统上提供程序安装卸载，以及程序开机自启功能，开发者更少的去关注各种平台之间的差异。
## 功能
+ [x] 安装应用
+ [x] 卸载应用
+ [x] 启动应用
+ [x] 停止应用
+ [x] 注册应用到开机启动
+ [x] 查看应用版本
+ [ ] 设置应用语言（多语言支持）
+ [x] 可扩展命令

### 如何使用

```go
package main

import (
	"fmt"
	"github.com/guoliang1994/installer/cmd"
	"github.com/kardianos/service"
	"time"
)


type Client struct {
}

func (t *Client) Start(s service.Service) error {
	status, err := s.Status()
	fmt.Println(err)
	fmt.Println(status)
	go t.run()
	return nil
}
func (t *Client) run() {
	for {
		fmt.Println("nice to meet you")
		time.Sleep(time.Second * 1)
	}
}

func (t *Client) Stop(s service.Service) error {
	return nil
}

func main() {
	description := `xyz程序描述`
	installer := cmd.NewInstaller()
	client := cmd.NewProgram("xyz", "xyz客户端", "xyz客户端描述", "v1.0.0", &Client{})
	installer.AddProgram(client).
		SetRootCmd("xyz", "xyz程序", description).
		Install()
}

```
效果

```shell
go run main.go

xyz客户端描述

Usage:
  xyz [command]

Available Commands:
  help        Help about any command
  install     安装 xyz客户端
  lang        设置 xyz客户端 语言
  restart     重启 xyz客户端
  run         前台运行 xyz客户端
  start       后台启动 xyz客户端
  status      查看 xyz客户端 状态
  stop        停止 xyz客户端
  uninstall   卸载 xyz客户端
  version     查看 xyz客户端 版本

Flags:
  -h, --help   help for xyz

Use "nice [command] --help" for more information about a command.
```

### Centos Ubuntu systemctl
```shell
chmod +x xyz

./xyz install

之后就可以支持

systemctl start xyz
systemctl stop xyz
systemctl ...  xyz

```

## 平台支持
+ [x] windows
+ [x] linux
+ [x] openWrt
+ [x] solaris
+ [x] open-rc
+ [x] freebsd

## 语言支持
+ [x] 中文
+ [x] English
+ [ ] 日文
+ [ ] 韩文
