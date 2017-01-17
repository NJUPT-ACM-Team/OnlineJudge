package db

import (
	"errors"
)

type MySQLConfig struct {
	Drivername string
	Username   string
	Password   string
	DBname     string
	Address    string
	Protocol   string
	Params     map[string]string
}

func (this *MySQLConfig) GetDriverName() string {
	if this.Drivername == "" {
		return "mysql"
	}
	return this.Drivername
}

func (this *MySQLConfig) GetDataSourceName() (string, error) {
	if this.Username == "" {
		return "", errors.New("username not provided.")
	}
	if this.DBname == "" {
		return "", errors.New("dbname not provided.")
	}

	dsn := this.Username
	if this.Password != "" {
		dsn += ":" + this.Password
	}

	dsn += "@"

	if this.Protocol != "" {
		dsn += this.Protocol
		if this.Address != "" {
			dsn += "(" + this.Address + ")"
		}
	}

	dsn += "/" + this.DBname

	if len(this.Params) != 0 {
		dsn += "?"
		first := true
		for k, v := range this.Params {
			if first {
				first = false
			} else {
				dsn += "&"
			}
			dsn += k + "=" + v
		}
	}
	return dsn, nil
}
