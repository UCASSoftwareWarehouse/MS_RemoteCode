package mongodb

import (
	"log"
	"path/filepath"
	"remote_code/config"
	"testing"
)

func TestInitEngine(t *testing.T) {
	config.InitConfigDefault()
	InitEngine()
}

func Test2(t *testing.T) {
	path, _ := filepath.Abs("")
	log.Println(path)
}
