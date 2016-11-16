package db

import (
	"errors"
)

type MySQLConfig struct {
	drivername string
	username   string
	password   string
	dbname     string
	address    string
	protocol   string
	params     map[string]string
}

func (this *MySQLConfig) GetDriverName() string {
	if this.drivername == "" {
		return "mysql"
	}
	return this.drivername
}

func (this *MySQLConfig) GetDataSourceName() (string, error) {
	if this.username == "" {
		return "", errors.New("username not provided.")
	}
	if this.dbname == "" {
		return "", errors.New("dbname not provided.")
	}

	dsn := this.username
	if this.password != "" {
		dsn += ":" + this.password
	}

	dsn += "@"

	if this.protocol != "" {
		dsn += this.protocol
		if this.address != "" {
			dsn += "(" + this.address + ")"
		}
	}

	dsn += "/" + this.dbname

	if len(this.params) != 0 {
		dsn += "?"
		first := true
		for k, v := range this.params {
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
