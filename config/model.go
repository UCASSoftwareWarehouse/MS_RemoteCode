package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type ConfigurationEnv string

const (
	DevEnv ConfigurationEnv = "dev"
	PrdEnv ConfigurationEnv = "prd"
)

type Configuration map[ConfigurationEnv]*EachConfig

type EachConfig struct {
	AppName                   string `yaml:"app_name"`
	Host                      string `yaml:"host"`
	Port                      int    `yaml:"port"`
	TransformCodeSplitter     string `yaml:"transform_code_splitter"`
}

const (
	DefaultConfigFilepath = "./config.yml"
)


func parse(configFilepath string) Configuration {
	println()
	if configFilepath == "" {
		configFilepath = DefaultConfigFilepath
	}
	bs, err := ioutil.ReadFile(configFilepath)
	if err != nil {
		log.Fatalf("ConfigForEnv parse failed, read file failed, err=[%v]", err)
	}
	conf := make(Configuration)
	err = yaml.Unmarshal(bs, &conf)
	if err != nil {
		log.Fatalf("ConfigForEnv parse failed, unmarshal config failed, err=[%v]", err)
	}
	return conf
}
