package model

import (
	"time"

	userType "github.com/rudianto-dev/gotemp-api-service/internal/domain/user/model/type"
	"github.com/rudianto-dev/gotemp-sdk/pkg/utils"
)

// User is representing the user data struct
type User struct {
	ID        string
	Name      string
	Username  string
	Password  string
	Status    userType.Status
	CreatedAt int64
	UpdatedAt int64
}

func New(req userType.Create) (domain *User, err error) {
	time := time.Now().Unix()
	domain = &User{
		ID:        utils.GenerateUUID(),
		Name:      req.Name,
		Username:  req.Username,
		Status:    req.Status,
		CreatedAt: time,
		UpdatedAt: time,
	}
	return
}

func Update(req userType.Edit) (domain *User, err error) {
	time := time.Now().Unix()
	if req.ID == "" {
		err = utils.ErrBadRequest
		return
	}
	domain = &User{
		ID:        req.ID,
		Name:      req.Name,
		UpdatedAt: time,
	}
	return
}
