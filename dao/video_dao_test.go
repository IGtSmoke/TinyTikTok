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

func TestUnix(t *testing.T) {
	var timestamp int64 = 1675575960832
	fmt.Println(time.Unix(timestamp, 0))
}
