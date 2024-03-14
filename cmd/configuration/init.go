package configuration

import (
	"log"

	"github.com/spf13/viper"
)

type ConfigurationSchema struct {
	Host     Host     `json:"host"`
	Database Database `json:"database"`
	Redis    Redis    `json:"redis"`
	Graceful Graceful `json:"graceful"`
}

type Host struct {
	Address string `json:"address"`
	Debug   bool   `json:"debug"`
}

type Database struct {
	Master                  string `json:"master"`
	Slave                   string `json:"slave"`
	RetryBackoff            int    `json:"retry_backoff" mapstructure:"retry_backoff"`
	IntervalCheckConnection int    `json:"interval_check_connection" mapstructure:"interval_check_connection"`
}

type Redis struct {
	Address  string `json:"address"`
	Password string `json:"password"`
	DB       int    `json:"db"`
	MaxRetry int    `json:"max_retry" mapstructure:"max_retry"`
}

type Graceful struct {
	TimeoutInSecond int64 `json:"timeout_in_second" mapstructure:"timeout_in_second"`
}

func NewConfiguration() *ConfigurationSchema {
	v := viper.New()
	v.SetConfigFile("/home/rats/project/go-template/gotemp-api-service/sample.config.json")

	config := ConfigurationSchema{}
	err := v.ReadInConfig()
	if err != nil {
		log.Panic(err)
	}
	err = v.UnmarshalExact(&config)
	if err != nil {
		log.Panic(err)
	}
	return &config
}
