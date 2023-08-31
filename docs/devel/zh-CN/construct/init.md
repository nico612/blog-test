## 初始化代码仓库
开始 Go 项目开发的第一步便是：初始化一个项目仓库，对于 Go 项目来说主要包括以下 3 内容：

1. 创建项目目录；

2. 初始化目录为 Go 模块；

3. 初始化目录为 Git 仓库。


### 创建项目目录
- 项目名称： 项目名要具有一定语义，说明该项目的功能等，建议的格式为纯小写的精短名字，如果项目名字过长，可以按单词用 - 分割，但最好不用到 -。这些是一些合格的名字：api、controllermanager、controller-manager。不建议这样命名：controller_manager；


- 确认大小写格式： 还要确认项目名在代码中的大小写格式，统一大小写格式，可以使整个代码看起来更加一体。例如：controller-manager / controllermanager 项目的小写格式为 controllermanager，大写格式为 ControllerManager；


- 确认项目名在代码中的简写格式： 有些项目名字出于易读目的，会比较长。在编写代码时，如果引用了项目名，可能会使代码行过长，为了使代码行简短易读，通常会使用简写模式。带 - 分割的项目名的简短模式，一般为每个单词的首字母，例如：controller-manager 为 cm。不带 - 分割的简短模式，需要你根据具体名字确定，并没有统一命名规则，例如：controller 可以为 ctrl。

### 初始化目录为Go模块
因为我们这是一个 Go 项目，根据 Go 语法要求，还需要将该项目初始化为一个 Go 模块。初始化命令如下：
```shell
$ go mod init # 初始化当前项目为一个 Go 模块
$ go work use . # 添加当前模块到 Go 工作区
```

### 初始化目录为 Git 仓库
初始化为 Git 仓库的第一步，就是在当前目录添加一个 .gitignore 文件，里面包含不期望 Git 跟踪的文件，例如：临时文件等。你可以使用生成工具 [gitignore.io](https://www.toptal.com/developers/gitignore) 来生成 .gitignore：

**注意**：Git 不追踪空目录，为了让 Git 追踪空目录，我们可以在空目录下创建一个空文件 .keep，并在适当的时候执行以下命令删除这些临时的 .keep 文件：`find . -name .keep | xargs -i rm {}`


### 创建Hello World 程序

```shell
# 格式化
$ gofmt -s -w ./
# 编译
$ go build -o _output/miniblog -v cmd/miniblog/main.go

# 编译成功后输出
command-line-arguments

$ ls _output/
miniblog

# 执行可执行文件
$ ./_output/miniblog
Hello MiniBlog
```

### 程序实时加载、构建、启动
在开发过程中，我们经常需要修改代码、编译代码、重新启动程序，然后测试程序。这个过程如果每次都手工操作，无疑效率是比较低的，那么有没有什么手段或者工具能够解放双手，提高开发效率呢？答案是：可以使用一些程序热加载工具。


业界当前有很多好用的程序热加载工具，在 Go 项目开发中，比较受欢迎的是 [air](https://github.com/cosmtrek/air) 工具。关于如何使用 air 工具，你可以直接参考官方文档 [Air 官方文档](https://github.com/cosmtrek/air/blob/master/README-zh_cn.md)。

1. 安装air工具
   ```shell
        go install github.com/cosmtrek/air@latest
    ```
2. 配置 air 工具。
   使用 air 官方仓库中给出的示例配置：[air_example.toml](https://github.com/cosmtrek/air/blob/master/air_example.toml)。air_example.toml 里面的示例配置基本能满足绝大部分的项目需求，一般只需要再配置 cmd、bin、args_bin、 full_bin 4 个参数即可， 因为官方会更新，具体看文档。

   在 miniblog 项目根目录下创建 .air.toml 文件，
   
   .air.toml 基于 air_example.toml 文件修改了以下参数配置：
   ```shell
   # 只需要写你平常编译使用的 shell 命令。你也可以使用 `make`.
   cmd = "make build"
   # 由 `cmd` 命令得到的二进制文件名.
   bin = "_output/miniblog"
   ```
   参数介绍：

   - cmd：指定了监听文件有变化时，air 需要执行的命令，这里指定了 make build 重新构建 miniblog 二进制文件；

   - bin：指定了执行完 cmd 命令后，执行的二进制文件，这里指定了编译构建后的二进制文件 _output/miniblog。
3. 启动air工具
   配置好后，在项目根目录下运行 air 命令：

```shell
$ air # 默认使用当前目录下的 .air.toml 配置，你可以通过 `-c` 选项指定配置，例如：`air -c .air.toml`

...

mkdir /home/colin/workspace/golang/src/github.com/marmotedu/miniblog/tmp
watching .
watching _output
...
watching scripts
!exclude tmp
building...
running...
Hello MiniBlog
```

### 编写 API 文档
采用 REST API 风格 和 RPC API 风格，

推荐编写方法：

- 如果 API 接口文档可以暴露在 SwaggerHub 平台，建议基于 SwaggerHub 平台来编写，因为 SwaggerHub 平台具备 API 编辑、展示、Mock、保存等功能，并能很方便地进行 API 共享。

- 如果 API 接口文档比较敏感，则可以基于本地编辑器编辑，并将 API 接口文档粘贴在 [Swagger Editor](https://editor-next.swagger.io/) 进行正确性校验。

- 不建议基于代码标注生成，因为基于代码标注生成，往往说明接口实现已经开发好，再基于接口生成 API 文档，违反接口定义先行的原则，不利于并行开发以提高开发效率；

- 文档编写工具也在不断发展，如果你有更好编辑方式，也可以采用。

miniblog OpenAPI 文档为 api/openapi/openapi.yaml, 相当于是本地编辑的文档，然后通过 swagger 工具，渲染并在线展示，具体步骤如下：

1. 安装 swagger 工具，安装命令如下：
```shell
$ go install github.com/go-swagger/go-swagger/cmd/swagger@latest
```

2. 运行 swagger 命令：
```shell
$ swagger serve -F=swagger --no-open --port 65534 ./api/openapi/openapi.yaml
2022/11/22 21:19:49 serving docs at http://localhost:65534/docs
```

**这里需要注意**：使用 swagger serve 渲染 OpenAPI 文档需要确保 OpenAPI 文档版本为：swagger: "2.0"，例如：

```shell
swagger: "2.0"
servers:
  - url: http://127.0.0.1:8080/v1
    description: development server
info:
  version: "1.0.0"
  title: miniblog api definition
```

否则 swagger serve 命令会渲染失败。

编写后的 OpenAPI 文档需要根据目录规范存放在：api/openapi 目录下。

### 使用构建工具提高开发效率

- 如果没有特殊需求，建议 make 和 cmake 中，选择 make（原因：make 更普适）；

- 对于一般的项目（应该是绝大多数的 Go 项目），可以使用更加通用的 make 工具；

- 对于超大型的项目（例如：公司级别的 Git 大仓）可以考虑使用 bazel。

miniblog 项目选择 make 作为构建工具。业界优秀的项目基本都是采用 make 来管理的，例如：Kubernetes、Docker、Istio 等等。


#### 编写简单的 Makefile
建议你通过以下方式来学习 Makefile 编程：
1. 学习 Makefile 基本语法：可参考文档 [Makefile基础知识.md](https://github.com/marmotedu/geekbang-go/blob/master/makefile/Makefile%E5%9F%BA%E7%A1%80%E7%9F%A5%E8%AF%86.md)；
2. 学习 Makefile 高级语法（如果有时间/感兴趣）：陈皓老师编写的 跟我一起写 Makefile (PDF 重制版) 。

编写后的 Makefile 文件位于项目根目录下，内容为：

```makefile
# ==============================================================================
# 定义全局 Makefile 变量方便后面引用



# MAKEFILE_LIST)：是Makefile的内置变量，表示：make所需处理的makefile文件列表，当前makefile的文件名总是位于列表的最后，文件名之前以空格进行分割；

# 函数 $(lastword <text>) 取字符串<text>中的最后一个单词，并返回字符串最后一个单词
# 函数 $(dir <names...>) 从文件名序列 <names> 中取出目录部分。目录部分是指最后一个反斜杠（/）之前的部分。如果没有反斜杠，那么返回 ./；
# ./
COMMON_SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))

# 项目根目录 绝对路径
#  $(abspath <text>) 将text中的各路径转换成绝对路径，并将转换后的结果返回
ROOT_DIR := $(abspath $(shell cd $(COMMON_SELF_DIR)/ && pwd -P))

# 构建产物、临时文件存放目录
OUTPUT_DIR := $(ROOT_DIR)/_output

# ==============================================================================
# 定义 Makefile all 伪目标，执行 `make` 时，会默认会执行 all 伪目标
.PHONY: all
all: format build

# ==============================================================================
# 定义其他需要的伪目标

.PHONY: build
build: tidy # 编译源码，依赖 tidy 目标自动添加/移除依赖包.
	@go build -v -o $(OUTPUT_DIR)/miniblog $(ROOT_DIR)/cmd/miniblog/main.go

.PHONY: format
format: # 格式化 Go 源码.
	@gofmt -s -w ./

.PHONY: add-copyright
add-copyright: # 添加版权头信息.
	@addlicense -v -f $(ROOT_DIR)/scripts/boilerplate.txt $(ROOT_DIR) --skip-dirs=third_party,vendor,$(OUTPUT_DIR)

.PHONY: swagger
swagger: # 启动 swagger 在线文档.
	@swagger serve -F=swagger --no-open --port 65534 $(ROOT_DIR)/api/openapi/openapi.yaml

.PHONY: tidy
tidy: # 自动添加/移除依赖包.
	@go mod tidy

.PHONY: clean
clean: # 清理构建产物、临时文件等. 要忽略命令的出错，需要在各个command之前加上减号-
	@-rm -vrf $(OUTPUT_DIR)

```

