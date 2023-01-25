// Package conf contains global configuration variables / Load configuration file to global variables
package conf

import (
	"os"

	"github.com/rs/zerolog/log"
	yaml "gopkg.in/yaml.v3"
)

// Config Global configuration variables
type Config struct {
	DSN                  string `yaml:"dsn"`
	SnowflakeID          int64  `yaml:"snowflakeId"`
	BucketName           string `yaml:"bucketName"`
	GinPort              string `yaml:"ginPort"`
	MinioEndpoint        string `yaml:"minioEndpoint"`
	MinioAccessKeyID     string `yaml:"minioAccessKeyId"`
	MinioSecretAccessKey string `yaml:"minioSecretAccessKey"`
	RedisAddr            string `yaml:"redisAddr"`
	RedisPassword        string `yaml:"redisPassword"`
	RedisDB              int    `yaml:"redisDb"`
}

// Conf Global configuration variables
var Conf *Config

// LoadConfig Load configuration file to global variables
func LoadConfig() error {
	ymlFile, err := os.ReadFile("./conf/config.yaml")
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
