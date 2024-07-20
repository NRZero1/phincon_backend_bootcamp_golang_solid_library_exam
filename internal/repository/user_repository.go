package repository

import (
	"solid_library_exam/internal/domain"
)

type UserRepositoryInterface interface {
	UserSave
	// UserBorrowBook
	// UserReturnBook
	UserFindByUsername
}

type UserSave interface {
	Save(*domain.User) (domain.User, error)
}

type UserReturnBook interface {
	ReturnBook(int) (domain.User, error)
}

type UserBorrowBook interface {
	BorrowBook(int) (domain.User, error)
}

type UserFindByUsername interface {
	FindByUsername(string) (domain.User, error)
}
