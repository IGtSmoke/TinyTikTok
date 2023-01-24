package conf

import (
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Config struct {
	DSN                  string `yaml:"DSN"`
	SnowflakeId          int64  `yaml:"SnowflakeId"`
	BucketName           string `yaml:"BucketName"`
	MinioEndpoint        string `yaml:"MinioEndpoint"`
	MinioAccessKeyID     string `yaml:"MinioAccessKeyID"`
	MinioSecretAccessKey string `yaml:"MinioSecretAccessKey"`
	RedisAddr            string `yaml:"RedisAddr"`
	RedisPassword        string `yaml:"RedisPassword"`
	RedisDB              int    `yaml:"RedisDB"`
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
