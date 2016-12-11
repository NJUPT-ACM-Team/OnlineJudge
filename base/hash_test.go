package base

import (
	"testing"
)

func TestMD5Hash(t *testing.T) {
	msg := "Hello World"
	t.Log(MD5Hash([]byte(msg)))
}
