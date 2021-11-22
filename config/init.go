package config

import (
	"log"
	"os"
	"remote_code/utils"
	"runtime"
)

var Conf *EachConfig

func IsProd() bool {
	log.Println(runtime.GOOS)
	if runtime.GOOS == "linux" {
		return true
	}
	return false
}

func InitConfig(configFilepath string, env ConfigurationEnv) {
	c := parse(configFilepath)
	Conf = c[env]
}

func InitConfigDefault() {
	//相对路径
	pwd, _ := os.Getwd()
	pwd = utils.GetParentDirectory(pwd)
	pwd += "/config.yml"
	//c := parse(DefaultConfigFilepath)
	c := parse(pwd)
	//判断mac or linux
	if IsProd() {
		Conf = c[PrdEnv]
	} else {
		Conf = c[DevEnv]
	}
	log.Printf("InitConfigDefault %+v", Conf)
}
