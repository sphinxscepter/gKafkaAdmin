package config

type AppConfig struct {
	App App `mapstructure:"app" yaml: "app"`
}

type App struct {
	AppName string     `mapstructure:"appName" yaml: "appNmae"`
	Server  ServerInfo `mapstructure:"server" yaml: "server"`
}

type ServerInfo struct {
	Ip   string `mapstructure:"ip" yaml: "ip"`
	Port int    `mapstructure:"port" yaml: "port"`
}

var AppConf = new(AppConfig)
