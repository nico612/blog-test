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