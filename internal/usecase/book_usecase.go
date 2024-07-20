package usecase

import (
	"solid_library_exam/internal/domain"
)

type BookUseCaseInterface interface {
	BookSave
	BookUpdateStatus
	BookFindById
	BookFindByTitle
}

type BookSave interface {
	Save(domain.Book) (domain.Book, error)
}

type BookUpdateStatus interface {
	UpdateStatus(int) (domain.Book, error)
}

type BookFindById interface {
	FindById(int) (domain.Book, error)
}

type BookFindByTitle interface {
	FindByTitle(string) (domain.Book, error)
}
