package main

import (
	"gopkg.in/guoliang1994/installer.v1/cmd"
	"gopkg.in/guoliang1994/installer.v1/example/program"
)

func main() {
	description := `xyz程序描述`
	installer := cmd.NewInstaller()
	client := cmd.NewProgram("client", "xyz客户端", "xyz客户端描述", "v1.0.0", &program.Client{})
	server := cmd.NewProgram("server", "xyz client", "xyz client description", "v1.0.0", &program.Server{})
	installer.AddProgram(client).
		AddProgram(server).
		SetRootCmd("np", "nat proxy", description).
		Install()
}
