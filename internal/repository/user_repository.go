package repository

import (
	"solid_library_exam/internal/domain"
)

type UserRepositoryInterface interface {
	UserSave
	// UserBorrowBook
	// UserReturnBook
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
