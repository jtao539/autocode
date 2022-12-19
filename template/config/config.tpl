package config

import (
	"fmt"
	"github.com/jinzhu/configor"
	"go.uber.org/zap"
)

var Conf *ServerConfig

type ServerConfig struct {
	DB     DBConfig     `yaml:"db"`
	App    AppConfig    `yaml:"app"`
	Redis  Redis        `yaml:"redis"`
}

type DBConfig struct {
	DBHost string `default:"127.0.0.1"`
	DBUser string `default:"root"`
	DBPass string `default:"123456"`
	DBPort string `default:"3306"`
	DBName string `default:"sale"`
}

type AppConfig struct {
	UploadPath string `yaml:"uploadPath"`
	Port       string `yaml:"port"`
	LogPath    string `yaml:"logPath"`
	ZapLog     bool   `yaml:"zapLog"`
}

type Redis struct {
	DB       int    ` yaml:"db"`      // redis的哪个数据库
	Addr     string ` yaml:"addr"`    // 服务器地址:端口
	Password string `yaml:"password"` // 密码
}

func Init() {
	conf := new(ServerConfig)
	err := configor.Load(conf, "./config/config.yml")
	if err != nil {
		zap.L().Error(err.Error())
		panic(fmt.Sprintf("Error in config: %s", err.Error()))
	}
	conf.App.Port = ":" + conf.App.Port
	Conf = conf
}
