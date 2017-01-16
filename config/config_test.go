package config

import (
	"testing"
)

func TestLoad(t *testing.T) {
	file := "../etc/config.json"
	config := Load(file)
	t.Log(config)
	t.Log(config.GetDBConfig())
	t.Log(config.GetMQConfig())
}
