package configuration

import (
	"log"

	"github.com/rudianto-dev/gotemp-sdk/pkg/utils"
)

type ConfigurationSchema struct {
	Host     Host     `json:"host"`
	JWT      JWT      `json:"jwt"`
	OTP      OTP      `json:"otp"`
	Database Database `json:"database"`
	Redis    Redis    `json:"redis"`
	Alert    Alert    `json:"alert"`
	Graceful Graceful `json:"graceful"`
}

type Host struct {
	Address string `json:"address"`
	Debug   bool   `json:"debug"`
}

type JWT struct {
	Public       string `json:"public"`
	Private      string `json:"private"`
	ExpireInHour int    `json:"expire_in_hour" mapstructure:"expire_in_hour"`
}

type OTP struct {
	ExpireInSecond             int `json:"expire_in_second" mapstructure:"expire_in_second"`
	VerificationExpireInSecond int `json:"verification_expire_in_second" mapstructure:"verification_expire_in_second"`
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

type Alert struct {
	Telegram Telegram `json:"telegram"`
}

type Telegram struct {
	URL       string `json:"url"`
	Token     string `json:"token"`
	ChannelID string `json:"channel_id" mapstructure:"channel_id"`
}

type Graceful struct {
	TimeoutInSecond int64 `json:"timeout_in_second" mapstructure:"timeout_in_second"`
}

func New() *ConfigurationSchema {
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
