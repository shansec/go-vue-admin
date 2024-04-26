package core

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github/shansec/go-vue-admin/global"
	"github/shansec/go-vue-admin/utils"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Viper(path ...string) *viper.Viper {
	var config string
	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose config file")
		flag.Parse()
		if config == "" {
			if configEnv := os.Getenv(utils.CONFIG_ENV); configEnv == "" {
				config = utils.CONFIG_FILE
				fmt.Printf("你正在使用 config 的默认值,config 的路径为%v\n", utils.CONFIG_FILE)
			} else {
				config = configEnv
				fmt.Printf("你正在使用 CONFIG 环境变量,config 的路径为%v\n", config)
			}
		} else {
			fmt.Printf("你正在使用命令行-c参数传递的值,config 的路径为%v\n", config)
		}
	} else {
		config = path[0]
		fmt.Printf("你正在使用 Viper 传递的值,config 的路径为%v\n", config)
	}
	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		fmt.Printf("Fatal error config file: %s \n", err)
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("config file changed: %s\n", e.Name)
		if err := v.Unmarshal(&global.MAY_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&global.MAY_CONFIG); err != nil {
		fmt.Println(err)
	}

	global.MAY_CONFIG.AutoCode.Root, _ = filepath.Abs(".")
	global.MAY_CONFIG.AutoCode.WRoot, _ = filepath.Abs("../go-vue")
	return v
}
