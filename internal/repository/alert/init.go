package alert

import (
	alertInterface "github.com/rudianto-dev/gotemp-api-service/internal/domain/alert"
	"github.com/rudianto-dev/gotemp-sdk/pkg/logger"
)

type AlertRepository struct {
	logger   *logger.Logger
	telegram *TelegramClient
}

type TelegramClient struct {
	URL       string
	Token     string
	ChannelID string
}

func NewAlertRepository(logger *logger.Logger, telegram *TelegramClient) alertInterface.Repository {
	return &AlertRepository{
		logger:   logger,
		telegram: telegram,
	}
}
