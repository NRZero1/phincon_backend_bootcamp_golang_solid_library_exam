package handler

import (
	"solid_library_exam/internal/domain"
)

type BookHandlerInterface interface {
	BookSave
	BookFindById
	BookFindByTitle
}

type BookSave interface {
	Save(domain.Book) (domain.Book, error)
}

type BookFindById interface {
	FindById(int) (domain.Book, error)
}

type BookFindByTitle interface {
	FindByTitle(string) ([]domain.Book, error)
}