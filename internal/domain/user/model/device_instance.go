package model

import (
	"time"

	userType "github.com/rudianto-dev/gotemp-api-service/internal/domain/user/model/type"
	"github.com/rudianto-dev/gotemp-sdk/pkg/utils"
)

type DeviceInstance struct {
	ID         string
	DeviceID   string
	InstanceID string
	UserID     string
	CreatedAt  int64
	UpdatedAt  int64
}

func NewDeviceInstance(req userType.CreateDeviceInstance) (domain *DeviceInstance, err error) {
	time := time.Now().Unix()
	if req.UserID == "" {
		err = utils.ErrBadRequest
		return
	}
	domain = &DeviceInstance{
		ID:         utils.GenerateUUID(),
		DeviceID:   req.DeviceID,
		InstanceID: req.InstanceID,
		UserID:     req.UserID,
		CreatedAt:  time,
		UpdatedAt:  time,
	}
	return
}

func UpdateDeviceInstance(req userType.UpdateDeviceInstance) (domain *DeviceInstance, err error) {
	time := time.Now().Unix()
	if req.ID == "" {
		err = utils.ErrBadRequest
		return
	}
	if req.UserID == "" {
		err = utils.ErrBadRequest
		return
	}
	domain = &DeviceInstance{
		ID:         req.ID,
		DeviceID:   req.DeviceID,
		InstanceID: req.InstanceID,
		UserID:     req.UserID,
		UpdatedAt:  time,
	}
	return
}
