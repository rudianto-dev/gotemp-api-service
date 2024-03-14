package user

import (
	"time"

	userContract "github.com/rudianto-dev/gotemp-sdk/contract/user"
	"github.com/rudianto-dev/gotemp-sdk/pkg/utils"
)

// User is representing the user data struct
type User struct {
	ID        string
	Name      string
	CreatedAt int64
	UpdatedAt int64
}

func NewUser(req userContract.CreateUser) (domain *User, err error) {
	time := time.Now().Unix()
	domain = &User{
		ID:        utils.GenerateUUID(),
		Name:      req.Name,
		CreatedAt: time,
		UpdatedAt: time,
	}
	return
}

func UpdateUser(req userContract.EditUser) (domain *User, err error) {
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
