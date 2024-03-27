package model

import (
	"time"

	alertType "github.com/rudianto-dev/gotemp-api-service/internal/domain/alert/model/type"
	"github.com/rudianto-dev/gotemp-sdk/pkg/utils"
)

type Alert struct {
	ID        string
	Message   string
	Sender    string
	Receiver  string
	Channel   alertType.Channel
	CreatedAt int64
}

func New(req alertType.Create) (domain *Alert) {
	time := time.Now().Unix()
	domain = &Alert{
		ID:        utils.GenerateUUID(),
		Message:   req.Message,
		Sender:    req.Sender,
		Receiver:  req.Receiver,
		Channel:   req.Channel,
		CreatedAt: time,
	}
	return
}
