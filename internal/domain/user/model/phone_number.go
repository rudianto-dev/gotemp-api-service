package model

import (
	"time"

	userType "github.com/rudianto-dev/gotemp-api-service/internal/domain/user/model/type"
	"github.com/rudianto-dev/gotemp-sdk/pkg/utils"
)

type PhoneNumber struct {
	ID          string
	PhoneNumber string
	UserID      string
	CreatedAt   int64
	UpdatedAt   int64
}

func NewPhoneNumber(req userType.CreatePhoneNumber) (domain *PhoneNumber, err error) {
	time := time.Now().Unix()
	if req.UserID == "" {
		err = utils.ErrBadRequest
		return
	}
	domain = &PhoneNumber{
		ID:          utils.GenerateUUID(),
		PhoneNumber: req.PhoneNumber,
		UserID:      req.UserID,
		CreatedAt:   time,
		UpdatedAt:   time,
	}
	return
}

func UpdatePhoneNumber(req userType.UpdatePhoneNumber) (domain *PhoneNumber, err error) {
	time := time.Now().Unix()
	if req.ID == "" {
		err = utils.ErrBadRequest
		return
	}
	if req.UserID == "" {
		err = utils.ErrBadRequest
		return
	}
	domain = &PhoneNumber{
		ID:          req.ID,
		PhoneNumber: req.PhoneNumber,
		UpdatedAt:   time,
	}
	return
}
