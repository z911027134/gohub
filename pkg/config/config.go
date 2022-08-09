package config

import (
	"github.com/spf13/cast"
	viperlib "github.com/spf13/viper"
	"gohub.com/pkg/helpers"
	"os"
)

// viper实例
var viper *viperlib.Viper

// 定义一个方法类型，返回map[string]interface{}
type ConfigFunc func() map[string]interface{}

// 定义一个map, 存储需要加载的配置
var ConfigFuncs map[string]ConfigFunc

func init() {
	// 初始化viper实例
	viper = viperlib.New()
	// 设置配置文件类型env
	viper.SetConfigType("env")
	// 设置配置文件路径， 相对于引用的位置（main.go）
	viper.AddConfigPath(".")
	// 环境变量前缀
	viper.SetEnvPrefix("appenv")
	// 读取环境变量（支持 flags）
	viper.AutomaticEnv()
	// ConfigFuncs 开辟空间
	ConfigFuncs = make(map[string]ConfigFunc)
}

// 初始化配置信息，完成对环境变量以及 config 信息的加载
func InitConfig(env string) {
	// config加载
	loadEnv(env)
	// 设置配置信息
	loadConfig()
}

func loadEnv(envSuffix string) {

	// 默认加载 .env 文件，如果有传参 --env=name 的话，加载 .env.name 文件
	envPath := ".env"
	if len(envSuffix) > 0 {
		filepath := ".env." + envSuffix
		if _, err := os.Stat(filepath); err == nil {
			// 如 .env.testing 或 .env.stage
			envPath = filepath
		}
	}
	// 加载 env
	viper.SetConfigName(envPath)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	// 监控 .env 文件，变更时重新加载
	viper.WatchConfig()
}

func loadConfig() {
	for name, fn := range ConfigFuncs {
		viper.Set(name, fn())
	}
}

// Env 读取环境变量，支持默认值
func Env(envName string, defaultVal ...interface{}) interface{} {
	if len(defaultVal) > 0 {
		return internalGet(envName, defaultVal[0])
	}
	return internalGet(envName)
}

func internalGet(path string, defaultVal ...interface{}) interface{} {
	// 如果viper没有读取到配置信息或读取到空值，并且有默认值，则返回默认值
	if !viper.IsSet(path) || helpers.Empty(viper.Get(path)) {
		if len(defaultVal) > 0 {
			return defaultVal[0]
		}
	}
	return viper.Get(path)
}

// add新增配置项
func Add(name string, fn ConfigFunc) {
	ConfigFuncs[name] = fn
}

// Get函数获取字符串类型的配置
func Get(path string, defaultVal ...interface{}) string {
	return GetString(path, defaultVal...)
}

func GetString(path string, defaultVal ...interface{}) string {
	return cast.ToString(internalGet(path, defaultVal...))
}

// GetInt 获取 Int 类型的配置信息
func GetInt(path string, defaultValue ...interface{}) int {
	return cast.ToInt(internalGet(path, defaultValue...))
}

// GetFloat64 获取 float64 类型的配置信息
func GetFloat64(path string, defaultValue ...interface{}) float64 {
	return cast.ToFloat64(internalGet(path, defaultValue...))
}

// GetInt64 获取 Int64 类型的配置信息
func GetInt64(path string, defaultValue ...interface{}) int64 {
	return cast.ToInt64(internalGet(path, defaultValue...))
}

// GetUint 获取 Uint 类型的配置信息
func GetUint(path string, defaultValue ...interface{}) uint {
	return cast.ToUint(internalGet(path, defaultValue...))
}

// GetBool 获取 Bool 类型的配置信息
func GetBool(path string, defaultValue ...interface{}) bool {
	return cast.ToBool(internalGet(path, defaultValue...))
}

// GetStringMapString 获取结构数据
func GetStringMapString(path string) map[string]string {
	return viper.GetStringMapString(path)
}
