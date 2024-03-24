package repository

import (
	userDomain "github.com/rudianto-dev/gotemp-api-service/internal/domain/user/model"
	userType "github.com/rudianto-dev/gotemp-api-service/internal/domain/user/model/type"
)

type UserEntity struct {
	ID        string `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	Username  string `json:"username" db:"username"`
	Password  string `json:"password" db:"password"`
	Status    int8   `json:"status" db:"status"`
	CreatedAt int64  `json:"created_at" db:"created_at"`
	UpdatedAt int64  `json:"updated_at" db:"updated_at"`
}

func ToUserEntity(domain *userDomain.User) *UserEntity {
	user := &UserEntity{
		ID:        domain.ID,
		Name:      domain.Name,
		Status:    int8(domain.Status),
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
	if domain.Username != "" {
		user.Username = domain.Username
	}
	return user
}

func ToUserDomain(entity *UserEntity) *userDomain.User {
	return &userDomain.User{
		ID:        entity.ID,
		Name:      entity.Name,
		Username:  entity.Username,
		Password:  entity.Password,
		Status:    userType.Status(entity.Status),
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}

func ToUserDomains(entities []*UserEntity) []*userDomain.User {
	domains := []*userDomain.User{}
	for _, entity := range entities {
		domains = append(domains, ToUserDomain(entity))
	}
	return domains
}
