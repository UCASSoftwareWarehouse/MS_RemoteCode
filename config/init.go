package config

import "fmt"

var Conf *EachConfig

func InitConfig(configFilepath string, env ConfigurationEnv) {
	c := parse(configFilepath)
	Conf = c[env]
}

func InitConfigDefault() {
	c := parse(DefaultConfigFilepath)
	Conf = c[DevEnv]
	fmt.Printf("%+v",Conf)

}