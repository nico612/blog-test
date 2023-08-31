package main

import (
	"flag"
	"fmt"

	// 解决 GOMAXPROCS 可能设置过大，导致生成线程过多，从而导致严重的上下文切换，浪费 CPU，降低程序性能的潜在问题。
	_ "go.uber.org/automaxprocs"

	"github.com/nico612/blog-test/internal/miniblog"
	"os"
)

var (
	// GitVersion 时候语义化的版本号
	GitVersion = "v0.0.0-master+$Format:%h$"
	// BuildDate 是 ISO8601 格式的构建时间, $(date -u +'%Y-%m-%dT%H:%M:%SZ') 命令的输出.
	BuildDate = "1970-01-01T00:00:00Z"
)

// Go 程序的默认入口函数（主函数）

func main() {

	version := flag.Bool("version", false, "Print version info.")
	flag.Parse()

	if *version {
		fmt.Println("GitVersion", GitVersion)
		fmt.Println("BuildDate", BuildDate)
	}

	fmt.Println("ok")

	command := miniblog.NewMiniBlogCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
