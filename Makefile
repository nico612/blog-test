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


# =============================================================================
# 定义版本相关变量

## 指定应用使用的version包， 会通过`-ldflags -X` 向该包中指定的变量注入值
VERSION_PACKAGE=github.com/nico612/blog-test/pkg/version

## 定义 VERSION 语义化版本号
ifeq ($(origin VERSION), undefined) # 判断变量VERSION 是否没有被定义
VERSION := $(shell git describe --tags --always --match='v*')  # 命令获取版本号；
endif

## 检查代码仓库是否是 dirty（默认dirty）
GIT_TREE_STATE:="dirty"

ifeq (, $(shell git status --porcelain 2>/dev/null))
	GIT_TREE_STATE="clean"
endif

GIT_COMMIT:=$(shell git rev-parse HEAD) # 获取构建时的 Commit ID；


## 向pkg/version包中的变量注入值
GO_LDFLAGS += \
	-X $(VERSION_PACKAGE).GitVersion=$(VERSION) \
	-X $(VERSION_PACKAGE).GitCommit=$(GIT_COMMIT) \
	-X $(VERSION_PACKAGE).GitTreeState=$(GIT_TREE_STATE) \
	-X $(VERSION_PACKAGE).BuildDate=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')



# ==============================================================================
# 定义 Makefile all 伪目标，执行 `make` 时，会默认会执行 all 伪目标
.PHONY: all
all: format build

# ==============================================================================
# 定义其他需要的伪目标

.PHONY: build
build: tidy # 编译源码，依赖 tidy 目标自动添加/移除依赖包.
#	@go build -v -o $(OUTPUT_DIR)/miniblog $(ROOT_DIR)/cmd/miniblog/main.go
#   编译时通过 -ldflags 注入版本信息
	@go build -v -ldflags "$(GO_LDFLAGS)" -o $(OUTPUT_DIR)/miniblog $(ROOT_DIR)/cmd/miniblog/main.go

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



