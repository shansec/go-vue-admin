package config

type JWT struct {
	SigningKey   string `mapstructure:"signing-key" json:"signingKey" yaml:"signing-key"`         // jwt 签名
	AExpiresTime int64  `mapstructure:"a-expires-time" json:"AExpiresTime" yaml:"a-expires-time"` // access_token 过期时间
	RExpiresTime int64  `mapstructure:"r-expires-time" json:"RExpiresTime" yaml:"r-expires-time"` // refresh_token 过期时间
	BufferTime   int64  `mapstructure:"buffer-time" json:"bufferTime" yaml:"buffer-time"`         // 缓冲时间
	Issuer       string `mapstructure:"issuer" json:"issuer" yaml:"issuer"`                       // 签发者
}
