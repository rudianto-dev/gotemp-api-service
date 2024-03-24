package model

import (
	"time"

	userType "github.com/rudianto-dev/gotemp-api-service/internal/domain/user/model/type"
	"github.com/rudianto-dev/gotemp-sdk/pkg/utils"
)

type CIF struct {
	ID          string
	UserID      string
	ReferenceID string
	CreatedAt   int64
	UpdatedAt   int64
}

func NewCIF(req userType.CreateCIF) (domain *CIF, err error) {
	time := time.Now().Unix()
	if req.UserID == "" || req.ReferenceID == "" {
		err = utils.ErrBadRequest
		return
	}
	domain = &CIF{
		ID:          utils.GenerateUUID(),
		UserID:      req.UserID,
		ReferenceID: req.ReferenceID,
		CreatedAt:   time,
		UpdatedAt:   time,
	}
	return
}

func UpdateCIF(req userType.UpdateCIF) (domain *CIF, err error) {
	time := time.Now().Unix()
	if req.UserID == "" || req.ReferenceID == "" {
		err = utils.ErrBadRequest
		return
	}
	domain = &CIF{
		ID:          req.ID,
		UserID:      req.UserID,
		ReferenceID: req.ReferenceID,
		UpdatedAt:   time,
	}
	return
}
