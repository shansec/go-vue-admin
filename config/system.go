package config

type System struct {
	Env           string `mapstructure:"env" json:"env" yaml:"env"`                              // 环境值
	Addr          int    `mapstructure:"addr" json:"addr" yaml:"addr"`                           // 端口值
	DbType        string `mapstructure:"db-type" json:"dbType" yaml:"db-type"`                   // 数据库类型
	RouterPrefix  string `mapstructure:"router-prefix" json:"routerPrefix" yaml:"router-prefix"` // 路由前缀
	OssType       string `mapstructure:"oss-type" json:"ossType" yaml:"oss-type"`                // OSS 类型
	UseMultipoint bool   `mapstructure:"use-multipoint" json:"useMultipoint" yaml:"use-multipoint"`
	UseRedis      bool   `mapstructure:"use-redis" json:"useRedis" yaml:"use-redis"`
	LimitCountIP  int    `mapstructure:"limit-count-ip" json:"limitCountIP" yaml:"limit-count-ip"`
	LimitTimeIP   int    `mapstructure:"limit-time-ip" json:"limitTimeIP" yaml:"limit-time-ip"`
}
