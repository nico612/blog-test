## 版本规范

本项目遵循 [语义化版本 2.0.0](https://semver.org/lang/zh-CN/) 规范。
在实际开发中，当开发完一个 miniblog 特性后，会编译 miniblog 二进制文件并发布到生产环境，很多时候为了定位问题和出于安全目的（确认发的是正确的版本），我们需要知道当前 miniblog 的版本，以及一些编译信息，例如编译时 Go 的版本、Git 目录是否干净，以及基于哪个 git commmit 来编译的。在一个编译好的可执行程序中，我们通常可以用类似 ./appname -v 的方式来获取版本信息。

查看版本信息
```shell
$ _output/miniblog --version
   gitCommit: 93864ffc831f5565b85b274639eb6e816a3f1632
gitTreeState: dirty                                   
   buildDate: 2022-11-25T05:54:24Z                    
   goVersion: go1.19                                  
    compiler: gc                                      
    platform: linux/amd64     
```

都是用 go build -ldflags='-X main.Version=v1.0.0' 这种方式来实现的。

`go build -ldflags="-X main.Version=v1.0.0" `所执行操作的含义：go build 时，通过指定 -ldflags 选项，将 v1.0.0 赋值给 main 包中的 Version 变量中。之后，程序内可以通过打印 Version 变量的值，来输出版本号（v1.0.0）。