package main

import (
	// 解决 GOMAXPROCS 可能设置过大，导致生成线程过多，从而导致严重的上下文切换，浪费 CPU，降低程序性能的潜在问题。
	_ "go.uber.org/automaxprocs"

	"github.com/nico612/blog-test/internal/miniblog"
	"os"
)

// Go 程序的默认入口函数（主函数）

func main() {

	command := miniblog.NewMiniBlogCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
