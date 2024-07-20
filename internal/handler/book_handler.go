package handler

import (
	"solid_library_exam/internal/domain"
)

/*seperating function into each of their own interface so struct won't be forced to implement
function it doesn't need (implementing I)*/
type BookHandlerInterface interface {
	BookSave
	BookSetUserId
	BookFindById
	BookFindByTitle
	BookGetAll
}

type BookSave interface {
	Save(domain.Book) (domain.Book, error)
}

type BookSetUserId interface {
	SetUserId(int, int) (domain.Book, error)
}

type BookFindById interface {
	FindById(int) (domain.Book, error)
}

type BookFindByTitle interface {
	FindByTitle(string) ([]domain.Book, error)
}

type BookGetAll interface {
	GetAll() ([]domain.Book)
}