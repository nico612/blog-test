## 应用程序构建

合格的应用构建方式具有以下几个特点：

- 应用代码结构清晰： 清晰的结构不仅利于阅读，还能减少后期改动带来的 Bug；

- 应用代码按功能隔离： 好的隔离性可以确保后期代码变更不会相互影响。例如功能类的代码变更，不会影响应用框架；

- 应用代码具有扩展性： 能够清晰地知道在哪里添加新的功能代码，并且添加之后仍然能够保持整个应用框架逻辑一致性。

### 应用程序组成部分及构建方法

一个 Go 应用一般由以下 3 部分组成：

- 应用配置；

- 应用业务逻辑；

- 应用启动框架。

#### 应用配置
- 命令行选项、命令行参数： 选择 pflag；
- 配置文件： 建议选择支持多种配置文件格式的包，也即从：viper、configor、koanf、config 中选择其一，毫无疑问 viper 胜出：
- 环境变量： 如果环境变量不多，可以使用 os.Getenv，如果环境变量很多，可以使用 envconfig 直接将环境变量读取到 Go 结构体变量中。
- 配置中心： 可根据需要选择 viper、apollo、etcd、consul 等。其实一般的项目不需要引入配置中心，因为使用配置中心，会带来一些部署、维护的复杂度。

#### 应用业务逻辑处理

一般而言，一个 Go 应用中会执行以下类别的业务逻辑处理（可能会用到其中一个或多个）：

- 初始化缓存；

- 初始化并创建各类数据库客户端，例如：Redis、MySQL、Kafka、MongoDB、Etcd 等；

- 初始化并创建其他服务的客户端等；

- 初始化并启动Web服务，例如：HTTP、HTTPS、GRPC；

- 启动异步任务，这些异步任务可以执行任何业务需要的操作，例如：watch kube-apiserver、定期从第三方服务拉取数据，并缓存、注册 /metrics 并监听指定的端口、启动 kafka 消费队列等等；

- 执行特定的业务处理，并退出程序；

- 还有很多其他业务逻辑。



#### 应用启动框架
使用cobra 可以和 pflag 组合来读取命令行，实现更强大、更易用的命令行参数处理能力。

启动框架你可以理解为一个 main 函数，只不过这里的 main 函数是有代码结构的，并可能分散在多个 Go 源码文件中，在这个大函数中，你可以读取配置文件、初始化业务逻辑、启动 Web 服务等，例如：


``` go
package main

import (
    "fmt"
    "net/http"

    "github.com/spf13/pflag"
)

const helpText = `Usage: main [flags] arg [arg...]

This is a very simple app framework (does nothing).

Flags:`

var (
    addr = pflag.String("addr", ":8777", "The address to listen to.")
    help = pflag.BoolP("help", "h", false, "Show this help message.")

    usage = func() {
        fmt.Println(helpText)
        pflag.PrintDefaults()
    }
)

func main() {
    // 1. 命令行参数处理：解析，并读取命令行参数
    pflag.Usage = usage
    pflag.Parse()
    if *help {
        pflag.Usage()
        return
    }

    // 2. 业务处理：初始化路由
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, "Hello world")
    })
    server := http.Server{Addr: *addr}
    fmt.Printf("Starting http server at %s\n", *addr)

    // 3. 业务处理：启动 HTTP Web 服务
    if err := server.ListenAndServe(); err != nil {
        panic(err)
    }
}
```

### 最佳构建方法
使用 pflag、viper、cobra 来构建一个强大的应用程序（这也是当前大部分团队选择的构建方法）。因为 pflag、viper、cobra 功能强大，


pflag：[如何使用Pflag给应用添加命令行标识](https://github.com/marmotedu/geekbang-go/blob/master/%E5%A6%82%E4%BD%95%E4%BD%BF%E7%94%A8Pflag%E7%BB%99%E5%BA%94%E7%94%A8%E6%B7%BB%E5%8A%A0%E5%91%BD%E4%BB%A4%E8%A1%8C%E6%A0%87%E8%AF%86.md)；

viper：配置解析神器-Viper全解；

cobra：现代化的命令行框架-Cobra全解。
```shell
go get -u github.com/spf13/cobra 
```

### 应用构建

