package base

import (
	"testing"
	"time"
)

func TestGetDefaultTime(t *testing.T) {
	t.Log(GetDefaultTime())
}

func TestMarshalTime(t *testing.T) {
	tm := time.Now()
	t.Log(MarshalTime(tm))
}

func TestUnmarshalTime(t *testing.T) {
	// tm, err := UnmarshalTime("Thu May 25 2017 19:16:30 GMT+0800 (CST)")
	tm, err := UnmarshalTime("Thu, 25 May 2017 11:15:58 GMT")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tm.UTC())
}
