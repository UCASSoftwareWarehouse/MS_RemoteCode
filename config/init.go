package config

import (
	"log"
)

var Conf *EachConfig

func InitConfig(configFilepath string, env ConfigurationEnv) {
	c := parse(configFilepath)
	Conf = c[env]
}

func InitConfigDefault() {
	c := parse(DefaultConfigFilepath)
	Conf = c[DevEnv]
	log.Printf("InitConfigDefault %+v", Conf)
}
