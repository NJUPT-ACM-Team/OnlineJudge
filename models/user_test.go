package models

import (
	"OnlineJudge/models/db"
	"testing"
)

/*
func TestUserInsert(t *testing.T) {
	db, err := db.NewDB()
	if err != nil {
		t.Fatal(err)
	}
	tx, err := db.Beginx()
	if err != nil {
		t.Fatal(err)
	}
	um := NewUserModel()
	user := &User{
		Username:     "Kevince",
		Password:     "123456",
		Email:        "abc@gmail.com",
		RegisterTime: time.Now(),
		Permission:   "root",
	}
	id, err := um.Insert(tx, user)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(id)
	if err := tx.Commit(); err != nil {
		t.Fatal(err)
	}
}
*/

func TestUserQueryByName(t *testing.T) {
	db, err := db.NewDB()
	if err != nil {
		t.Fatal(err)
	}
	tx, err := db.Beginx()
	if err != nil {
		t.Fatal(err)
	}
	um := NewUserModel()
	user, err := um.QueryByName(tx, "Kevince", nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(user)
}

func TestValidate(t *testing.T) {
	db, err := db.NewDB()
	if err != nil {
		t.Fatal(err)
	}
	tx, err := db.Beginx()
	if err != nil {
		t.Fatal(err)
	}
	um := NewUserModel()
	r, err := um.Validate(tx, "kevince", "123456")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}
