package utils

import (
	"TinyTikTok/conf"
	"TinyTikTok/conf/setup"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestUploadFile(t *testing.T) {
	conf.LoadConfig()
	setup.Minio()
	file, _ := os.Open("..\\tmp\\test.mp4")
	defer file.Close()
	fi, _ := os.Stat("..\\tmp\\test.mp4")
	err := UploadFile("videos", "test.mp4", file, fi.Size())
	fmt.Println(err)
}

func TestGetFileUrl(t *testing.T) {
	conf.LoadConfig()
	setup.Minio()
	url, err := GetFileUrl("videos", "test.mp4", 0)
	fmt.Println(url, err, strings.Split(url.String(), "?")[0])
	fmt.Println(url.Path, url.RawPath)
}
