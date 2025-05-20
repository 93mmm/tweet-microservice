package config

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

type config struct {
	Mongo mongoCfg
	App   appCfg
}

type hostPortConfig struct {
	Host string
	Port int
}

var (
	once = sync.Once{}
	cfg  *config
)

func Load() {
	once.Do(loadConfig)
}

func loadConfig() {
	viper.SetConfigFile(".env")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Can't read config file:", err)
	}

	cfg = &config{
		Mongo: loadMongoCfg(),
		App: loadAppCfg(),
	}
	if err := cfg.validate(); err != nil {
		log.Fatal(err)
	}
}
