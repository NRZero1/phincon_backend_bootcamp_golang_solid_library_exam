package handler

import (
	"solid_library_exam/internal/domain"
)

type BookHandlerInterface interface {
	BookSave
}

type BookSave interface {
	Save(domain.Book) (domain.Book, error)
}