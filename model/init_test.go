package model

import (
	"remote_code/config"
	"testing"
)

func TestInitGorm(t *testing.T) {
	config.InitConfigDefault()
	InitGorm()
}
