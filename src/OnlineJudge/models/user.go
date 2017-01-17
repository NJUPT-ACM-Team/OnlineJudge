package models

import (
	"OnlineJudge/base"

	"github.com/jmoiron/sqlx"

	"errors"
	"fmt"
	"time"
)

type User struct {
	UserId   int64 `db:"user_id"`
	Username string
	Password []byte
	Email    string
	Phone    string
	School   string
	Motto    string

	TotalLocalSubmit int `db:"total_local_submit"`
	TotalLocalAC     int `db:"total_local_ac"`
	TotalSubmit      int `db:"total_submit"`
	TotalAC          int `db:"total_ac"`

	RegisterTime  time.Time `db:"register_time"`
	LastLoginTime time.Time `db:"last_login_time"`
	LoginIPAddr   string    `db:"login_ip_addr"`
	Privilege     string
	LockStatus    int `db:"lock_status"`
}

type UserModel struct {
	Model
}

func NewUserModel() *UserModel {
	return &UserModel{Model{Table: "Users"}}
}

func hashPassword(passwd []byte) ([]byte, error) {
	if passwd == nil {
		return nil, errors.New("Empty password field")
	}
	return base.GenHash(passwd)
}

// Hash password before insert
func (this *UserModel) Insert(tx *sqlx.Tx, user *User) (int64, error) {
	var err error
	if user.Password, err = hashPassword(user.Password); err != nil {
		return 0, err
	}
	last_insert_id, err := this.InlineInsert(tx, user, nil, []string{"user_id"})
	if err != nil {
		return 0, err
	}
	return last_insert_id, nil
}

func (this *UserModel) Update(tx *sqlx.Tx, user *User, pk string, required []string, excepts []string) error {
	var err error
	if base.ArrayContains(required, "password") {
		if user.Password, err = hashPassword(user.Password); err != nil {
			return err
		}
	}
	if base.IsNilOrZero(required) && !base.ArrayContains(excepts, "password") {
		if user.Password, err = hashPassword(user.Password); err != nil {
			return err
		}
	}
	if pk == "" {
		pk = "user_id"
	}
	if err := this.InlineUpdate(tx, user, pk, required, excepts); err != nil {
		return err
	}
	return nil
}

func (this *UserModel) QueryById(tx *sqlx.Tx, id int, required []string, excepts []string) (*User, error) {
	user := User{}
	str_fields, err := GenerateSelectSQL(user, required, excepts)
	if err != nil {
		return nil, err
	}
	if err := tx.Get(&user, fmt.Sprintf("SELECT %s FROM %s WHERE user_id=?", str_fields, this.Table), id); err != nil {
		return nil, err
	}
	return &user, nil
}

func (this *UserModel) QueryByName(tx *sqlx.Tx, name string, required []string, excepts []string) (*User, error) {

	user := User{}
	str_fields, err := GenerateSelectSQL(user, required, excepts)
	if err != nil {
		return nil, err
	}
	if err := tx.Get(&user, fmt.Sprintf("SELECT %s FROM %s WHERE username=? LIMIT 1", str_fields, this.Table), name); err != nil {
		return nil, err
	}
	return &user, nil
}

func (this *UserModel) QueryIdByName(tx *sqlx.Tx, name string) (int64, error) {
	user, err := this.QueryByName(tx, name, []string{"user_id"}, nil)
	if err != nil {
		return 0, err
	}
	if user.UserId == 0 {
		return 0, errors.New("Failed to get user_id")
	}
	return user.UserId, nil
}

func (this *UserModel) Auth(tx *sqlx.Tx, name string, password []byte) (bool, error) {
	// log.Println(password)
	user, err := this.QueryByName(tx, name, []string{"password"}, nil)
	if err != nil {
		return false, err
	}
	return base.MatchHash(user.Password, password), nil
}

func (this *UserModel) UpdatePassword(tx *sqlx.Tx, name string, passwd []byte) error {
	user := &User{
		Username: name,
		Password: passwd,
	}
	return this.Update(tx, user, "username", []string{"password"}, nil)
}

func (this *UserModel) UpdateIPAddr(tx *sqlx.Tx, name string, ip string) error {
	user := &User{
		Username:    name,
		LoginIPAddr: ip,
	}
	return this.Update(tx, user, "username", []string{"login_ip_addr"}, nil)
}

func (this *UserModel) UpdateLastLoginTime(tx *sqlx.Tx, name string, t time.Time) error {
	user := &User{
		Username:      name,
		LastLoginTime: t,
	}
	return this.Update(tx, user, "username", []string{"last_login_time"}, nil)
}
