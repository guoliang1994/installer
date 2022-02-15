package cmd

import (
	"fmt"
	"github.com/kardianos/service"
	"github.com/spf13/cobra"
	"os/exec"
)

type program struct {
	binName     string
	appName     string
	description string
	version     string
	program     service.Interface
	service     service.Service
	RootCmd     *cobra.Command
}

type installer struct {
	programs     []*program
	rootCmd      *cobra.Command
	programLen   int
	isSetRootCmd bool
}

func NewInstaller() *installer {
	return &installer{}
}

func (this *installer) SetRootCmd(binName, appName, description string) *installer {
	this.rootCmd = &cobra.Command{
		Use:               binName,
		Short:             appName,
		Long:              description,
		CompletionOptions: cobra.CompletionOptions{HiddenDefaultCmd: true},
	}
	this.isSetRootCmd = true
	return this
}
func (this *installer) AddProgram(p *program) *installer {
	this.programs = append(this.programs, p)
	this.programLen++
	return this
}
func (this *installer) execute() {
	for _, p := range this.programs {
		p.start()
		p.stop()
		p.restart()
		p.install()
		p.uninstall()
		p.status()
		p.ver()
		p.run()
		p.Lang()
	}
	cobra.CheckErr(this.rootCmd.Execute())
}
func (this *installer) isMultiInstall() bool {
	return this.programLen > 1
}
func (this *installer) Install() {
	for _, p := range this.programs {
		// 初始化 service
		options := make(service.KeyValue)
		svcConfig := &service.Config{
			Name:        p.appName,
			DisplayName: p.appName,
			Description: p.description,
			Option:      options,
			UserName:    "root",
		}
		svcConfig.Dependencies = []string{
			"Requires=network.target",
			"After=network-online.target syslog.target"}
		if service.Platform() == "unix-systemv" {
			svcConfig.UserName = "root"
		}
		var err error
		// 增加 service 运行时的参数，最后得到 nps run
		if this.isMultiInstall() {
			if this.isSetRootCmd {
				// 如果有多个程序，则需要增加程序名称再 run
				svcConfig.Arguments = append(svcConfig.Arguments, p.binName)
				this.rootCmd.AddCommand(p.RootCmd)
			} else {
				this.rootCmd = p.RootCmd
				fmt.Println("您安装了多个程序，但是没有设置根命令，请使用 SetRootCmd 设置根命令，否则只会安装最后一个程序")
			}
		} else {
			// 如果只有一个程序，那根程序就等于子程序
			this.rootCmd = p.RootCmd
		}
		svcConfig.Arguments = append(svcConfig.Arguments, "run")
		p.service, err = service.New(p.program, svcConfig)
		if err != nil {
			fmt.Println(err)
		}
	}
	this.execute()
}

func NewProgram(binName, appName, description, version string, p service.Interface) *program {
	app := &program{
		binName:     binName,
		appName:     appName,
		description: description,
		version:     version,
		program:     p,
	}
	// 程序的名称就是根命令
	app.RootCmd = &cobra.Command{
		Use:   binName,
		Short: appName,
		Long:  description,
	}
	// 隐藏默认的命令
	app.RootCmd.CompletionOptions.HiddenDefaultCmd = true
	return app
}

func (i *program) install() {
	c := "install"
	var installCmd = &cobra.Command{
		Use:   c,
		Short: "安装 " + i.appName,
		Long:  "安装 " + i.appName + ",开机自启",
		RunE: func(cmd *cobra.Command, args []string) error {
			_ = i.service.Stop()
			_ = i.service.Uninstall()
			return i.service.Install()
		},
	}
	i.RootCmd.AddCommand(installCmd)
}

func (i *program) uninstall() {
	c := "uninstall"
	var installCmd = &cobra.Command{
		Use:   c,
		Short: "卸载 " + i.appName,
		Long:  "卸载 " + i.appName,
		RunE:  i.control(c),
	}
	i.RootCmd.AddCommand(installCmd)
}

func (i *program) run() error {
	c := "run"
	var installCmd = &cobra.Command{
		Use:   c,
		Short: "前台运行 " + i.appName,
		Long:  "前台运行 " + i.appName,
		RunE: func(cmd *cobra.Command, args []string) error {
			return i.service.Run()
		},
	}
	i.RootCmd.AddCommand(installCmd)
	return nil
}
func (i *program) start() {
	c := "start"
	var installCmd = &cobra.Command{
		Use:   c,
		Short: "后台启动 " + i.appName,
		Long:  "后台启动 " + i.appName,
		RunE:  i.control(c),
	}
	i.RootCmd.AddCommand(installCmd)
}

func (i *program) stop() {
	c := "stop"
	var installCmd = &cobra.Command{
		Use:   c,
		Short: "停止 " + i.appName,
		Long:  "停止 " + i.appName,
		RunE:  i.control(c),
	}
	i.RootCmd.AddCommand(installCmd)
}
func (i *program) restart() {
	c := "restart"
	var installCmd = &cobra.Command{
		Use:   c,
		Short: "重启 " + i.appName,
		Long:  "重启 " + i.appName,
		RunE:  i.control(c),
	}
	i.RootCmd.AddCommand(installCmd)
}
func (i *program) status() {
	c := "status"
	var installCmd = &cobra.Command{
		Use:   c,
		Short: "查看 " + i.appName + " 状态",
		Long:  "查看 " + i.appName + " 状态",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("")
		},
	}
	i.RootCmd.AddCommand(installCmd)
}
func (i *program) ver() {
	c := "version"
	var installCmd = &cobra.Command{
		Use:   c,
		Short: "查看 " + i.appName + " 版本",
		Long:  "查看 " + i.appName + " 版本",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(i.version)
		},
	}
	i.RootCmd.AddCommand(installCmd)
}
func (i *program) Lang() {
	c := "lang"
	var installCmd = &cobra.Command{
		Use:   c,
		Short: "设置 " + i.appName + " 语言",
		Long:  "设置 " + i.appName + " 语言",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(i.version)
		},
	}
	i.RootCmd.AddCommand(installCmd)
}

// start stop restart
func (i *program) control(command string) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		if service.Platform() == "unix-systemv" {
			terminal := exec.Command("/etc/init.d/"+i.appName, command)
			return terminal.Run()
		}
		return service.Control(i.service, command)
	}
}
