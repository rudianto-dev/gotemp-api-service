package user

import (
	userDomain "github.com/rudianto-dev/gotemp-api-service/internal/domain/user"
)

type UserEntity struct {
	ID        string `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	CreatedAt int64  `json:"created_at" db:"created_at"`
	UpdatedAt int64  `json:"updated_at" db:"updated_at"`
}

func toUserDomain(entity *UserEntity) *userDomain.User {
	return &userDomain.User{
		ID:        entity.ID,
		Name:      entity.Name,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}

func toUserDomains(entities []*UserEntity) []*userDomain.User {
	domains := []*userDomain.User{}
	for _, entity := range entities {
		domains = append(domains, toUserDomain(entity))
	}
	return domains
}

func toUserEntity(domain *userDomain.User) *UserEntity {
	return &UserEntity{
		ID:        domain.ID,
		Name:      domain.Name,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func toUserEntities(domains []*userDomain.User) []*UserEntity {
	entities := []*UserEntity{}
	for _, domain := range domains {
		entities = append(entities, toUserEntity(domain))
	}
	return entities
}
