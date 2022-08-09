package config

import "gohub.com/pkg/config"

// 设置app的配置信息，默认从.env配置中读取
func init() {
	// 这里触发了pkg/config包的init函数
	config.Add("app", func() map[string]interface{} {
		return map[string]interface{}{
			// 应用名称
			"name": config.Env("APP_NAME", "gohub"),
			// 当前环境
			"env": config.Env("APP_ENV", "local"),
			// debug模式
			"debug": config.Env("APP_DEBUG", false),
			// 端口
			"port": config.Env("APP_PORT", "3000"),
			// 加密key
			"key": config.Env("APP_KEY", "33446a9dcf9ea060a0a6532b166da32f304af0de"),
			// 设置时区，JWT 里会使用，日志记录里也会使用到
			"timezone": config.Env("TIMEZONE", "Asia/Shanghai"),
		}
	})
}
