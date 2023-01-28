package utils

import (
	"TinyTikTok/conf"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	err := conf.LoadConfig()
	if err != nil {
		return
	}
	t.Log(conf.Conf)
}
