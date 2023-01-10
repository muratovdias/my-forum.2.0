package config

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
)

type ConfigDB struct {
	Host     string
	User     string
	Password string
	DbName   string
	Port     int
	SslMode  string
}

func NewConfDB() *ConfigDB {
	configPath := flag.String("config", "", "Path the config file")
	flag.Parse()
	configs := &ConfigDB{}
	_, err := toml.DecodeFile(*configPath, configs)
	if err != nil {
		log.Printf("Ошибка декодирования файла конфигов %v", err)
	}
	return configs
}
