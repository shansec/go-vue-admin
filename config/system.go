package config

type System struct {
	// 环境值
	Env string `mapstructure:"env" json:"env" yaml:"env"`
	// 端口值
	Addr int `mapstructure:"addr" json:"addr" yaml:"addr"`
	// 数据库类型
	DbType string `mapstructure:"db-type" json:"dbType" yaml:"db-type"`
	// OSS 类型
	OssType       string `mapstructure:"oss-type" json:"ossType" yaml:"oss-type"`
	UseMultipoint bool   `mapstructure:"use-multipoint" json:"useMultipoint" yaml:"use-multipoint"`
	UseRedis      bool   `mapstructure:"use-redis" json:"useRedis" yaml:"use-redis"`
	LimitCountIP  int    `mapstructure:"limit-count-ip" json:"limitCountIP" yaml:"limit-count-ip"`
	LimitTimeIP   int    `mapstructure:"limit-time-ip" json:"limitTimeIP" yaml:"limit-time-ip"`
}
