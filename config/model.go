package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"runtime"
)

type ConfigurationEnv string

const (
	DevEnv ConfigurationEnv = "dev"
	PrdEnv ConfigurationEnv = "prd"
)

type Configuration map[ConfigurationEnv]*EachConfig

type EachConfig struct {
	AppName          string `yaml:"app_name"`
	Host             string `yaml:"host"`
	Port             int    `yaml:"port"`
	MongodbAddr      string `yaml:"mongodb_addr"`
	MysqlAddr        string `yaml:"mysql_addr"`
	Password         string `yaml:"password"`
	Username         string `yaml:"username"`
	Database         string `yaml:"database"`
	ConsulAddr       string `yaml:"consul_addr"`
	NetworkInterface string `yaml:"network_interface"`
}

const (
	//DefaultConfigFilepath = "./config.yml"
	DefaultConfigFilepath = "/Users/zhujianxing/GoLandProjects/code/MS_RemoteCode/config.yml"
)

func parse(configFilepath string) Configuration {
	println()
	if configFilepath == "" {
		configFilepath = DefaultConfigFilepath
	}
	bs, err := ioutil.ReadFile(configFilepath)
	if err != nil {
		log.Printf("ConfigForEnv parse failed, read file failed, err=[%v]", err)
	}
	conf := make(Configuration)
	err = yaml.Unmarshal(bs, &conf)
	if err != nil {
		log.Printf("ConfigForEnv parse failed, unmarshal config failed, err=[%v]", err)
	}
	return conf
}

func (c *EachConfig) GetEnv() ConfigurationEnv {
	log.Printf("RUNNING ON %s\n", runtime.GOOS)
	if runtime.GOOS == "linux" {
		return PrdEnv
	} else {
		return DevEnv
	}
}
