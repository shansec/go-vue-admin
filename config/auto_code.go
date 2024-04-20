package config

type AutoCode struct {
	SModel          string `mapstructure:"s-model" json:"s-model" yaml:"s-model"`
	SRouter         string `mapstructure:"s-router" json:"s-router" yaml:"s-router"`
	SServer         string `mapstructure:"s-server" json:"s-server" yaml:"s-server"`
	SApi            string `mapstructure:"s-api" json:"s-api" yaml:"s-api"`
	SPlug           string `mapstructure:"s-plug" json:"s-plug" yaml:"s-plug"`
	SInitialize     string `mapstructure:"s-initialize" json:"s-initialize" yaml:"s-initialize"`
	Root            string `mapstructure:"root" json:"root" yaml:"root"`
	WRoot           string `mapstructure:"w-root" json:"w-root" yaml:"w-root"`
	WTable          string `mapstructure:"w-table" json:"w-table" yaml:"w-table"`
	WWeb            string `mapstructure:"w-web" json:"w-web" yaml:"w-web"`
	SService        string `mapstructure:"s-service" json:"s-service" yaml:"s-service"`
	SRequest        string `mapstructure:"s-request" json:"s-request" yaml:"s-request"`
	WApi            string `mapstructure:"w-api" json:"w-api" yaml:"w-api"`
	WForm           string `mapstructure:"w-form" json:"w-form" yaml:"w-form"`
	TransferRestart bool   `mapstructure:"transfer-restart" json:"transfer-restart" yaml:"transfer-restart"`
}
