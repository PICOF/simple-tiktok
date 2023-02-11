package config

import (
	"testing"
	"time"
)

func TestConfig(t *testing.T) {
	v := GetConfig("feed")
	for {
		time.Sleep(time.Second * 5)
		println(v.GetString("server.serviceName"))
	}
}
