package config

import (
	"OnlineJudge/db"
	"OnlineJudge/mq"

	"encoding/json"
	"io"
	"log"
	"os"
	"strings"
)

var test_config = `
{
	"mysql": {
		"username": "test",
		"dbname": "ojtest2",
		"password": "abc123",
		"address": "localhost",
		"protocol": "tcp",
		"port": 3306,
		"params": {"parseTime": "true", "charset": "utf8"}
	},
	"rabbitmq": {
		"username": "oj",
		"password": "ojtest",
		"address": "192.168.56.1:5672"
	},
	"irpc": {
		"server_bind_address": "localhost:9999",
		"client_connect_address": "192.168.56.1"
	}

}
`

type Config struct {
	MySQL struct {
		Username string
		DBname   string
		Password string
		Address  string
		Protocol string
		Params   map[string]string
	}
	RabbitMQ struct {
		Username string
		Password string
		Address  string
	}
	IRPC struct {
		ServerBindAddress    string
		ClientConnectAddress string
	}
}

func (this *Config) GetDBConfig() *db.MySQLConfig {
	c := &db.MySQLConfig{
		Username: this.MySQL.Username,
		DBname:   this.MySQL.DBname,
		Password: this.MySQL.Password,
		Address:  this.MySQL.Address,
		Protocol: this.MySQL.Protocol,
		Params:   this.MySQL.Params,
	}
	return c
}

func (this *Config) GetMQConfig() *mq.MQConfig {
	c := &mq.MQConfig{
		Username: this.RabbitMQ.Username,
		Password: this.RabbitMQ.Password,
		Address:  this.RabbitMQ.Address,
	}
	return c
}

// func (this *Config) GetIRPCConfig()

func Load(path string) *Config {
	var file io.Reader
	if path == "" {
		log.Println("using test config.")
		file = strings.NewReader(test_config)
	} else {
		var err error
		file, err = os.Open(path)
		if err != nil {
			panic(err)
		}
	}
	decoder := json.NewDecoder(file)
	config := &Config{}
	err := decoder.Decode(config)
	if err != nil {
		panic(err)
	}
	return config
}
