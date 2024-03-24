package model

import (
	"time"

	userType "github.com/rudianto-dev/gotemp-api-service/internal/domain/user/model/type"
	"github.com/rudianto-dev/gotemp-sdk/pkg/utils"
)

type Email struct {
	ID        string
	Email     string
	UserID    string
	CreatedAt int64
	UpdatedAt int64
}

func NewEmail(req userType.CreateEmail) (domain *Email, err error) {
	time := time.Now().Unix()
	if req.UserID == "" {
		err = utils.ErrBadRequest
		return
	}
	domain = &Email{
		ID:        utils.GenerateUUID(),
		Email:     req.Email,
		UserID:    req.UserID,
		CreatedAt: time,
		UpdatedAt: time,
	}
	return
}

func UpdateEmail(req userType.UpdateEmail) (domain *Email, err error) {
	time := time.Now().Unix()
	if req.ID == "" {
		err = utils.ErrBadRequest
		return
	}
	if req.UserID == "" {
		err = utils.ErrBadRequest
		return
	}
	domain = &Email{
		ID:        req.ID,
		Email:     req.Email,
		UpdatedAt: time,
	}
	return
}
