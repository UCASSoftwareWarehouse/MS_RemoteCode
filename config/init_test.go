package config

import "testing"

func TestInitConfigDefault(t *testing.T) {
	InitConfigDefault()
	t.Logf("Default Config=[%+v]", Conf)
	InitConfig(DefaultConfigFilepath, PrdEnv)
	t.Logf("Prd Config=[%+v]", Conf)
}
