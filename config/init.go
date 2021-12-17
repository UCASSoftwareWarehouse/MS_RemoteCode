package config

import (
	"flag"
	"log"
	"os"
	"remote_code/utils"
	"runtime"
	"strings"
)

var Conf *EachConfig

func IsProd() bool {
	log.Println(runtime.GOOS)
	if runtime.GOOS == "linux" {
		return true
	}
	return false
}

//func InitConfig(configFilepath string, env ConfigurationEnv) {
//	c := parse(configFilepath)
//	Conf = c[env]
//}

func InitConfig() {
	args := &args{}
	initEnvironArgs(args)
	initCmdArgs(args)
	Conf = parse(args)
}

func InitConfigDefault() {
	//相对路径
	pwd, _ := os.Getwd()
	for !strings.HasSuffix(pwd, "MS_RemoteCode") {
		pwd = utils.GetParentDirectory(pwd)
	}
	pwd += "/config.yml"
	log.Printf("config file path:%+v", pwd)
	//c := parse(DefaultConfigFilepath)
	c := parse2(pwd)
	//判断mac or linux
	if IsProd() {
		Conf = c[PrdEnv]
	} else {
		Conf = c[DevEnv]
	}
	log.Printf("InitConfigDefault %+v", Conf)
}

func InitConfigWithFile(path string, env ConfigurationEnv) {
	args := &args{
		ConfigPath: path,
		Env:        env,
	}
	Conf = parse(args)
}

func initCmdArgs(args *args) {
	var configPath string
	var env string
	var port int
	flag.StringVar(&configPath, "config_path", "", "配置文件路径")
	flag.StringVar(&env, "env", "", "是否为测试测试环境，值为dev或prd")
	flag.IntVar(&port, "port", 0, "指定端口号，默认为配置文件中配置的端口号")
	flag.Parse()
	e, _ := convert2Env(env)
	if e != "" {
		args.Env = e
	}
	if port != 0 {
		args.Port = port
	}
	if configPath != "" {
		args.ConfigPath = configPath
	}
}

func initEnvironArgs(args *args) {
	configPath, ok := os.LookupEnv("CONFIG_PATH")
	if ok {
		args.ConfigPath = configPath
	}
	env, ok := os.LookupEnv("ENV")
	if ok {
		e, _ := convert2Env(env)
		if e != "" {
			args.Env = e
		}
	}
	networkInterface, ok := os.LookupEnv("NETWORK_INTERFACE")
	if ok {
		args.NetworkInterface = networkInterface
	}
}
