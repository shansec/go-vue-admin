package config

type Mysql struct {
	// 服务器地址
	Path string `mapstructure:"path" json:"path" yaml:"path"`
	// 端口
	Port string `mapstructure:"port" json:"port" yaml:"port"`
	// 高级配置
	Config string `mapstructure:"config" json:"config" yaml:"config"`
	// 数据库名称
	Dbname string `mapstructure:"db-name" json:"dbname" yaml:"db-name"`
	// 数据库用户名
	Username string `mapstructure:"username" json:"username" yaml:"username"`
	// 数据库密码
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	// 最大连接数
	MaxIdleConns int `mapstructure:"max-idle-conns" json:"maxIdleConns" yaml:"max-idle-conns"`
	// 打开数据库的最大数量
	MaxOpenConns int `mapstructure:"max-open-conns" json:"maxOpenConns" yaml:"max-open-conns"`
	// 是否开启Gorm全局日志
	LogMode string `mapstructure:"log-mode" json:"logMode" yaml:"log-mode"`
	// 是否通过zap写入日志
	LogZap bool `mapstructure:"log-zap" json:"logZap" yaml:"log-zap"`
}

func (m *Mysql) Dns() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ":" + m.Port + ")/" + m.Dbname + "?" + m.Config
}

func (m *Mysql) GetLogMode() string {
	return m.LogMode
}
