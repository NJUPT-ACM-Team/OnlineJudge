package db

func Init() {
	mysql_config := MySQLConfig{
		drivername: "mysql",
		username:   "test",
		dbname:     "ojtest",
		password:   "abc123",
	}
	var config Config
	config = &mysql_config
	UseConfig(config)
}
