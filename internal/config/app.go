package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type appCfg hostPortConfig

func loadAppCfg() appCfg {
	return appCfg{
		Host: viper.GetString("app_host"),
		Port: viper.GetInt("app_port"),
	}
}

func (c *appCfg) ConnectionString() string {
	return fmt.Sprintf(
		"%v:%v",
		c.Host,
		c.Port,
	)
}

func App() *appCfg {
	return &cfg.App
}
