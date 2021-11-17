package main

import (
	"remote_code/config"
	"remote_code/server"
)

func main() {
	// config.InitConfig()
	config.InitConfigDefault()
	server.StartServe()
}
