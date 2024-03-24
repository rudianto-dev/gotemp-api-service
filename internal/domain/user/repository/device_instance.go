package repository

import (
	userDomain "github.com/rudianto-dev/gotemp-api-service/internal/domain/user/model"
)

type DeviceInstanceEntity struct {
	ID         string `json:"id" db:"id"`
	UserID     string `json:"user_id" db:"user_id"`
	DeviceID   string `json:"device_id" db:"device_id"`
	InstanceID string `json:"instance_id" db:"instance_id"`
	CreatedAt  int64  `json:"created_at" db:"created_at"`
	UpdatedAt  int64  `json:"updated_at" db:"updated_at"`
}

func ToDeviceInstanceEntity(domain *userDomain.DeviceInstance) *DeviceInstanceEntity {
	user := &DeviceInstanceEntity{
		ID:         domain.ID,
		UserID:     domain.UserID,
		DeviceID:   domain.DeviceID,
		InstanceID: domain.InstanceID,
		CreatedAt:  domain.CreatedAt,
		UpdatedAt:  domain.UpdatedAt,
	}
	return user
}

func ToDeviceInstanceDomain(entity *DeviceInstanceEntity) *userDomain.DeviceInstance {
	return &userDomain.DeviceInstance{
		ID:         entity.ID,
		UserID:     entity.UserID,
		DeviceID:   entity.DeviceID,
		InstanceID: entity.InstanceID,
		CreatedAt:  entity.CreatedAt,
		UpdatedAt:  entity.UpdatedAt,
	}
}

func ToDeviceInstanceDomains(entities []*DeviceInstanceEntity) []*userDomain.DeviceInstance {
	domains := []*userDomain.DeviceInstance{}
	for _, entity := range entities {
		domains = append(domains, ToDeviceInstanceDomain(entity))
	}
	return domains
}
