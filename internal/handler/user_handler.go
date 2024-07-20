package handler

import (
	"solid_library_exam/internal/domain"
)

type UserHandlerInterface interface {
	UserSave
	UserFindByUsername
}

type UserSave interface {
	Save(domain.User) (domain.User, error)
}

type UserFindByUsername interface {
	FindByUsername(string) (domain.User, error)
}