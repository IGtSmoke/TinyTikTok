package setup

import (
	"TinyTikTok/conf"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	gorm "gorm.io/gorm"
)

// Mdb Export global MySQL client
var Mdb *gorm.DB

// Gorm Initialize MySQL client
func Gorm() {
	var err error
	Mdb, err = gorm.Open(mysql.Open(conf.Conf.DSN), &gorm.Config{})
	if err != nil {
		log.Err(err)
	}
}
