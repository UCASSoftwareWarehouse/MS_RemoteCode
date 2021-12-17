package config

import (
	"errors"
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

func convert2Env(env string) (ConfigurationEnv, error) {
	switch env {
	case "dev":
		return DevEnv, nil
	case "prd":
		return PrdEnv, nil
	default:
		return "", errors.New("illegal env, should be dev or prd")
	}
}

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
	Env              ConfigurationEnv
}

type args struct {
	ConfigPath       string
	Env              ConfigurationEnv
	NetworkInterface string
	Port             int
}

const (
	//DefaultConfigFilepath = "./config.yml"
	DefaultConfigFilepath = "/Users/zhujianxing/GoLandProjects/code/MS_RemoteCode/config.yml"
)

func parse2(configFilepath string) Configuration {
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

func parse(args *args) *EachConfig {
	if args.Env == "" {
		panic("ENV not set, plz check your environs or cmd args")
	}
	if args.ConfigPath == "" {
		panic("CONFIG_PATH not set, plz check your environs or cmd args")
	}
	bs, err := ioutil.ReadFile(args.ConfigPath)
	if err != nil {
		log.Fatalf("ConfigForEnv parse failed, read file failed, err=[%v]", err)
	}
	configs := make(Configuration)
	err = yaml.Unmarshal(bs, &configs)
	if err != nil {
		log.Fatalf("ConfigForEnv parse failed, unmarshal config failed, err=[%v]", err)
	}

	conf := configs[args.Env]
	// reset settings by args
	conf.Env = args.Env
	if args.Port != 0 {
		conf.Port = args.Port
	}
	if args.NetworkInterface != "" {
		conf.NetworkInterface = args.NetworkInterface
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
