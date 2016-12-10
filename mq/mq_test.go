package mq

import (
	"testing"
)

func TestDeclare(t *testing.T) {
	Init()
	mq := New()
	if err := mq.Connect(); err != nil {
		t.Fatal(err)
	}
	defer mq.Disconnect()

	if err := mq.DeclareLJ(); err != nil {
		t.Fatal(err)
	}
}
