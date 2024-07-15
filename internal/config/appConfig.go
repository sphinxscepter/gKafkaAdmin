package config

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func (appConfigInfo *AppConfig) InitConfiguration() *AppConfig {
	configFile := "conf/app.yaml"

	path, _ := os.Getwd()

	if configEnv := os.Getenv("VIPER_CONFIG"); configEnv != "" {
		configFile = configEnv
	}

	v := *viper.New()
	fmt.Println(path)
	v.AddConfigPath(path)
	v.SetConfigFile(configFile)
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("read config file error: %s\n", err))
	}

	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed: %s\n", in.Name)

		if err := v.Unmarshal(&appConfigInfo); err != nil {
			fmt.Println(err)
			panic(err)
		}
	})

	if err := v.Unmarshal(&appConfigInfo); err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(appConfigInfo.App.Server.Port)

	return appConfigInfo
}
