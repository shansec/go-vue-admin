package config

type JWT struct {
	// jwt 签名
	SigningKey string `mapstructure:"signing-key" json:"signingKey" yaml:"signing-key"`
	// 过期时间
	ExpiresTime int64 `mapstructure:"expires-time" json:"expiresTime" yaml:"expires-time"`
	// 缓冲时间
	BufferTime int64 `mapstructure:"buffer-time" json:"bufferTime" yaml:"buffer-time"`
	// 签发者
	Issuer string `mapstructure:"issuer" json:"issuer" yaml:"issuer"`
}
