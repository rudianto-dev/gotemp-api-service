package configuration

import (
	"log"

	"github.com/rudianto-dev/gotemp-sdk/pkg/utils"
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
	c, err := utils.LoadConfiguration()
	if err != nil {
		log.Panic(err)
	}

	config := ConfigurationSchema{}
	err = c.ReadInConfig()
	if err != nil {
		log.Panic(err)
	}
	err = c.UnmarshalExact(&config)
	if err != nil {
		log.Panic(err)
	}
	return &config
}
