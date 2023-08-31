package miniblog

import (
	"github.com/nico612/blog-test/internal/pkg/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
)

const (
	// recommendedHomeDir 定义放置 miniblog 服务配置的默认目录.
	recommendedHomeDir = ".miniblog"

	// 指定miniblog服务的默认配置文件名
	defaultConfigName = "miniblog.yaml"
)

// initConfig 设置需要读取的配置文件名、环境变量，并读取配置文件内容到 viper 中.
func initConfig() {
	if cfgFile != "" {
		// 从命令行选项指定的配置文件中读取
		viper.SetConfigFile(cfgFile)
	} else {
		// 使用默认配置文件

		// 查找用户主目录
		home, err := os.UserHomeDir()

		// 如果获取用户主目录失败，打印`Error.xxx` 错误，并退出程序（退出码为1）
		cobra.CheckErr(err)

		// 将用 `$HOME/<recommendedHomeDir>` 目录加入到配置文件的搜索路径中 `$HOME/.miniblog` 目录
		viper.AddConfigPath(filepath.Join(home, recommendedHomeDir))

		// 把当前目录加入到配置文件的搜索路径中国
		viper.AddConfigPath(".")

		// 设置配置文件格式为 Yaml
		viper.SetConfigType("yaml")

		// 配置文件名称 （没有文件扩展名）
		viper.SetConfigName(defaultConfigName)
	}

	// 读取匹配的环境变量
	viper.AutomaticEnv()

	//  读取环境变量的前缀为MINIBLOG，如果是miniblog, 将自动装变为大写， 一个独有的环境变量前缀，可以有效避免环境变量命名冲突；
	viper.SetEnvPrefix("MINIBlOG")

	// 以下2行，将viper.Get(key) key 字符串中 '.' 和 '-' '_'
	replacer := strings.NewReplacer(".", "_", "-", "_")
	viper.SetEnvKeyReplacer(replacer)

	// 读取配置文件，如果指定了配置文件名，则使用指定的配置文件，否则在注册的搜索路径中搜索
	if err := viper.ReadInConfig(); err != nil {
		log.Errorw("Failed to read viper configuration file", "err", err)
	}

	// 打印viper 当前使用的配置文件，方便Debug
	log.Infow("Using config file", "file", viper.ConfigFileUsed())

}

// logOptions 从 viper 中读取日志配置，构建 `*log.Options` 并返回.
// 注意：`viper.Get<Type>()` 中 key 的名字需要使用 `.` 分割，以跟 YAML 中保持相同的缩进.
func logOptions() *log.Options {
	return &log.Options{
		DisableCaller:     viper.GetBool("log.disable-caller"),
		DisableStacktrace: viper.GetBool("log.disable-stacktrace"),
		Level:             viper.GetString("log.level"),
		Format:            viper.GetString("log.format"),
		OutputPaths:       viper.GetStringSlice("log.output-paths"),
	}
}
