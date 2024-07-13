package config

type AppConfig struct {
	App Configuartion `mapstructure:"app" yaml: "appNmae"`
}

type Configuartion struct {
	AppName string     `mapstructure:"appName" yaml: "appNmae"`
	Server  ServerInfo `mapstructure:"server" yaml: "server"`
}

type ServerInfo struct {
	Ip   string `mapstructure:"ip" yaml: "ip"`
	Port int    `mapstructure:"port" yaml: "port"`
}

var AppConf = new(AppConfig)
