package models

import (
	"OnlineJudge/db"
	"testing"
)

/*
import "time"

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
		Username:     "kevince",
		Password:     []byte("123456"),
		Email:        "abc@gmail.com",
		RegisterTime: time.Now(),
		Privilege:    "root",
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

func TestUpdatePassword(t *testing.T) {
	db, err := db.NewDB()
	if err != nil {
		t.Fatal(err)
	}
	tx, err := db.Beginx()
	if err != nil {
		t.Fatal(err)
	}
	defer tx.Rollback()
	um := NewUserModel()
	if err := um.UpdatePassword(tx, "kevince", []byte("abc")); err != nil {
		t.Fatal(err)
	}
	if err := tx.Commit(); err != nil {
		t.Fatal(err)
	}
}

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
	user, err := um.QueryByName(tx, "kevince", nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(user)
}

func TestAuth(t *testing.T) {
	db, err := db.NewDB()
	if err != nil {
		t.Fatal(err)
	}
	tx, err := db.Beginx()
	if err != nil {
		t.Fatal(err)
	}
	um := NewUserModel()
	r, err := um.Auth(tx, "kevince", []byte("abc"))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}
