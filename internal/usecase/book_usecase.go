package usecase

import (
	"solid_library_exam/internal/domain"
)

type BookUseCaseInterface interface {
	BookSave
	BookUpdateStatus
}

type BookSave interface {
	Save(domain.Book) (domain.Book, error)
}

type BookUpdateStatus interface {
	UpdateStatus(int) (domain.Book, error)
}