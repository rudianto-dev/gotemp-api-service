package repository

import (
	alertDomain "github.com/rudianto-dev/gotemp-api-service/internal/domain/alert/model"
	alertType "github.com/rudianto-dev/gotemp-api-service/internal/domain/alert/model/type"
)

type AlertEntity struct {
	ID        string
	Message   string
	Sender    string
	Receiver  string
	Channel   alertType.Channel
	CreatedAt int64
}

func ToAlertEntity(domain *alertDomain.Alert) *AlertEntity {
	user := &AlertEntity{
		ID:        domain.ID,
		Message:   domain.Message,
		Sender:    domain.Sender,
		Receiver:  domain.Receiver,
		Channel:   domain.Channel,
		CreatedAt: domain.CreatedAt,
	}
	return user
}

func ToAlertDomain(entity *AlertEntity) *alertDomain.Alert {
	return &alertDomain.Alert{
		ID:        entity.ID,
		Message:   entity.Message,
		Sender:    entity.Sender,
		Receiver:  entity.Receiver,
		Channel:   entity.Channel,
		CreatedAt: entity.CreatedAt,
	}
}
