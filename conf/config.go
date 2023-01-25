// Package conf contains global configuration variables / Load configuration file to global variables
package conf

import (
	"io/ioutil"

	"github.com/rs/zerolog/log"
	yaml "gopkg.in/yaml.v3"
)

// Config Global configuration variables
type Config struct {
	DSN                  string `yaml:"dsn"`
	SnowflakeID          int64  `yaml:"snowflakeID"`
	BucketName           string `yaml:"bucketName"`
	MinioEndpoint        string `yaml:"minioEndpoint"`
	MinioAccessKeyID     string `yaml:"minioAccessKeyID"`
	MinioSecretAccessKey string `yaml:"minioSecretAccessKey"`
	RedisAddr            string `yaml:"redisAddr"`
	RedisPassword        string `yaml:"redisPassword"`
	RedisDB              int    `yaml:"redisDb"`
}

var Conf *Config

func LoadConfig() error {
	ymlFile, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		log.Error().Err(err).Msg("Failed to read config file")

		return err
	}

	if err = yaml.Unmarshal(ymlFile, &Conf); err != nil {
		log.Error().Err(err).Msg("Failed to unmarshal config file")

		return err
	}
	return nil
}
