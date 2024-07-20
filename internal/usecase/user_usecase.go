package usecase

import (
	"solid_library_exam/internal/domain"
)

type UserUseCaseInterface interface {
	UserSave
}

type UserSave interface {
	Save(domain.User) (domain.User, error)
}
