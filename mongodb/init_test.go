package mongodb

import (
	"remote_code/config"
	"testing"
)

func TestInitEngine(t *testing.T) {
	config.InitConfigDefault()
	InitEngine()
}
