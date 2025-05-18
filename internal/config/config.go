package config

import (
	"fmt"
	"log"
	"sync"

	"github.com/spf13/viper"
)

type mongoCfg struct {
    Host string
    Port int
}

func (c *mongoCfg) ConnectionString() string {
    return fmt.Sprintf(
		"mongodb://%v:%v",
        c.Host,
        c.Port,
    )
}

type config struct {
    Mongo mongoCfg
}

var (
    once = sync.Once{}
    cfg *config
)

func Load() {
    once.Do(loadConfig)
}

func Mongo() *mongoCfg {
    return &cfg.Mongo
}

func loadConfig() {
    viper.SetConfigName(".env")
    viper.SetConfigType("yml")
    viper.AddConfigPath(".")

    err := viper.ReadInConfig()
    if err != nil {
        log.Fatal(err)
    }

    cfg = &config{
        Mongo: mongoCfg{
            Host: viper.GetString("mongo.host"),
            Port: viper.GetInt("mongo.port"),
        },
    }
}
