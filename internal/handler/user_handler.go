package handler

import (
	"solid_library_exam/internal/domain"
)

type UserHandlerInterface interface {
	UserSave
	UserFindByUsername
	UserFindById
	UserBorrowBook
	UserGetAll
	UserReturnBook
}

type UserSave interface {
	Save(string, int) (domain.User, error)
}

type UserFindByUsername interface {
	FindByUsername(string) (domain.User, error)
}

type UserFindById interface {
	FindById(int) (domain.User, error)
}

type UserBorrowBook interface {
	BorrowBook(int, domain.Book) (domain.User)
}

type UserGetAll interface {
	GetAll() ([]domain.User)
}

type UserReturnBook interface {
	ReturnBook(int, int) (domain.User, error)
}