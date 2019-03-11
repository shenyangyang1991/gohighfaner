package gredis

import (
	"testing"
	"time"
)

func TestSetup(t *testing.T) {
	Setup()
	err := Set("first", "hello world", 10*time.Second)
	if err != nil {
		t.Error(err)
	}
	obj := ""
	err = Get("first", &obj)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(obj)
	}
}
