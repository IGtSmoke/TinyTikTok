package setup

import (
	"TinyTikTok/conf"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Mdb *gorm.DB

func Gorm() {
	var err error
	Mdb, err = gorm.Open(mysql.Open(conf.Conf.DSN), &gorm.Config{})
	if err != nil {
		log.Err(err)
	}
}
