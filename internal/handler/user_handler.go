package handler

import (
	"solid_library_exam/internal/domain"
)

type UserHandlerInterface interface {
	UserSave
}

type UserSave interface {
	Save(domain.User) (domain.User, error)
}