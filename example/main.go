package main

import (
	"gopkg.in/guoliang1994/go-i18n.v2"
	"gopkg.in/guoliang1994/installer.v1/cmd"
	"gopkg.in/guoliang1994/installer.v1/example/program"
)

func main() {
	description := `xyz程序描述`
	installer := cmd.NewInstaller()
	client := cmd.NewProgram(i18n.Chinese, "client", "xyz客户端", "xyz客户端描述", "v1.0.0", &program.Client{})
	server := cmd.NewProgram(i18n.English, "server", "xyz client", "xyz client description", "v1.0.0", &program.Server{})
	installer.AddProgram(client).
		AddProgram(server).
		SetRootCmd("np", "nat proxy", description).
		Install()
}
