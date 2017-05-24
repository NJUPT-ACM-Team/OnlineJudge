package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"log"
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

func (this *DBUtil) Rollback() {
	this.Tx.Rollback()
}

func (this *DBUtil) MustCommit() {
	if err := this.Tx.Commit(); err != nil {
		log.Println("failed to commit:" + err.Error())
		panic(err)
	}
}

func (this *DBUtil) Commit() error {
	return this.Tx.Commit()
}

func (this *DBUtil) MustBegin() *sqlx.Tx {
	if this.Tx != nil {
		this.Tx.Rollback()
	}
	this.Tx = this.db.MustBegin()
	return this.Tx
}

func (this *DBUtil) Close() {
	this.db.Close()
}

var dbM *sqlx.DB

func MustGetDB() *sqlx.DB {
	if dbM == nil {
		db, err := NewDB()
		if err != nil {
			panic(err)
		}
		dbM = db
	}
	return dbM
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
