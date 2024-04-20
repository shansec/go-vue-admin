package config

type Server struct {
	JWT    JWT    `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap    Zap    `mapstructure:"zap" json:"zap" yaml:"zap"`
	System System `mapstructure:"system" json:"system" yaml:"system"`
	// autocode
	AutoCode AutoCode `mapstructure:"autocode" json:"autocode" yaml:"autocode"`
	// gorm
	Mysql   Mysql   `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Captcha Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
}
