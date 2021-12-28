package config

import (
	"github.com/spf13/viper"
)

type config struct {
	IndexServer server   `mapstructure:"index_server"`
	AdminServer server   `mapstructure:"admin_server"`
	Database    database `mapstructure:"database"`
	Cache       cache    `mapstructure:"cache"`
	Es          es       `mapstructure:"es"`
	Kafka       kafka    `mapstructure:"kafka"`
}

type database struct {
	Address string `mapstructure:"address"`
	Sync    bool   `mapstructure:"sync"` // 是否同步数据表结构
	Debug   bool   `mapstructure:"debug"`
}

type server struct {
	Enable  bool   `mapstructure:"enable"`
	Address string `mapstructure:"address"`
	Port    int    `mapstructure:"port"`
	JwtKey  string `mapstructure:"jwt_key"`
}

type cache struct {
	Address string `mapstructure:"address"`
}

type es struct {
	Address  string `mapstructure:"address"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Prefix   string `mapstructure:"prefix"`
}

type kafka struct {
	Address  string `mapstructure:"address"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Prefix   string `mapstructure:"prefix"`
}

var Params config

func init() {
	viper.SetConfigName("dev")
	viper.AddConfigPath("../../config")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(&Params); err != nil {
		panic(err)
	}
}
