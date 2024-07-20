package handler

import (
	"solid_library_exam/internal/domain"
	"solid_library_exam/internal/usecase"
)

type BookHandler struct {
	BookSave usecase.BookSave
}

func NewBookHandler(bookUseCase usecase.BookSave) usecase.BookSave {
	return BookHandler{
		BookSave: bookUseCase,
	}
}

func (h BookHandler) Save(book domain.Book) (domain.Book, error) {
	savedBook, err := h.BookSave.Save(book)

	if err != nil {
		return domain.Book{}, err
	}

	return savedBook, nil
}