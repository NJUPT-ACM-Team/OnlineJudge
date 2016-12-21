package config

import (
	"testing"
)

func TestLoad(t *testing.T) {
	file := "../etc/config.json"
	config, err := Load(file)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(config)
	t.Log(config.GetDBConfig())
	t.Log(config.GetMQConfig())
}
