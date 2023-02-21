package config

type Zap struct {
	// 级别
	Level string `mapstructure:"level" json:"level" yaml:"level"`
	// 输出
	Format string `mapstructure:"format" json:"format" yaml:"format"`
	// 日志前缀
	Prefix string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
	// 日志文件夹
	Director string `mapstructure:"director" json:"director" yaml:"director"`
	// 显示行
	ShowLine bool `mapstructure:"show-line" json:"showLine" yaml:"show-line"`
	// 编码级
	EncodeLevel string `mapstructure:"encode-level" json:"encodeLevel" yaml:"encode-level"`
	// 栈名
	StacktraceKey string `mapstructure:"stacktrace-key" json:"stacktraceKey" yaml:"stacktrace-key"`
	// 输出控制台
	LogInConsole bool `mapstructure:"log-in-console" json:"logInConsole" yaml:"log-in-console"`
}
