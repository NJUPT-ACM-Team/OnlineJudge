package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Config interface {
	GetDriverName() string
	GetDataSourceName() (string, error)
}

var config Config

func UseConfig(cfg Config) {
	config = cfg
}

func New() *sqlx.DB {
	db, err := NewDB()
	if err != nil {
		panic(err)
	}
	return db
}

type DBUtil struct {
	db *sqlx.DB
	Tx *sqlx.Tx
}

func NewDBU(d *sqlx.DB) *DBUtil {
	return &DBUtil{db: d}
}

func (this *DBUtil) MustCommit() {
	if err := this.Tx.Commit(); err != nil {
		panic(err)
	}
}

func (this *DBUtil) MustBegin() *sqlx.Tx {
	this.Tx = this.db.MustBegin()
	return this.Tx
}

func (this *DBUtil) Close() {
	this.db.Close()
}

func NewDB() (*sqlx.DB, error) {
	InitTest()
	dn := config.GetDriverName()
	dsn, err := config.GetDataSourceName()
	if err != nil {
		return nil, err
	}
	db, err := sqlx.Open(dn, dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func GetConfig() Config {
	return config
}
