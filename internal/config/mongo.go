package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type mongoCfg hostPortConfig

func (c *mongoCfg) ConnectionString() string {
    return fmt.Sprintf(
		"mongodb://%v:%v",
        c.Host,
        c.Port,
    )
}

func Mongo() *mongoCfg {
    return &cfg.Mongo
}

func loadMongoCfg() mongoCfg {
	return mongoCfg{
		Host: viper.GetString("mongo_host"),
		Port: viper.GetInt("mongo_port"),
	}
}
