package main

import (
	"remote_code/config"
	"remote_code/mongodb"
	"remote_code/server"
)

func main() {
	// config.InitConfig()
	config.InitConfigDefault()
	mongodb.InitEngine()
	server.StartServer()
}
