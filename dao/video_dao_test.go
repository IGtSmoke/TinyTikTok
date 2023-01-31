package dao

import (
	"TinyTikTok/conf"
	"TinyTikTok/conf/setup"
	"fmt"
	"testing"
	"time"
)

func TestGetVideosAndNextTimeByLastTime(t *testing.T) {
	conf.LoadConfig()
	setup.Gorm()
	now := time.Now()
	fmt.Println(now)
	Videos, nextTime := GetVideosAndNextTimeByLastTime(now)
	fmt.Println(Videos, nextTime)
}
