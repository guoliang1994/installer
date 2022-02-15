package main

import (
	"embed"
	"github.com/guoliang1994/installer/cmd"
	program2 "github.com/guoliang1994/installer/example/program"
	"github.com/guoliang1994/installer/global"
)

//go:embed lang
var LangFs embed.FS

func main() {
	global.LangFs = LangFs
	description := `xyz程序描述`
	installer := cmd.NewInstaller()
	client := cmd.NewProgram("client", "xyz客户端", "xyz客户端描述", "v1.0.0", &program2.Client{})
	server := cmd.NewProgram("server", "xyz服务端", "xyz服务端描述", "v1.0.0", &program2.Server{})
	installer.AddProgram(client).
		AddProgram(server).
		SetRootCmd("xyz", "xyz程序", description).
		Install()
}
