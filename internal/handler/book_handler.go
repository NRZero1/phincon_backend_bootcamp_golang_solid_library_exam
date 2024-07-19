package handler

import (
	"solid_library_exam/internal/domain"
	"solid_library_exam/internal/usecase"
)

type BookHandlerInterface interface {
	Save()
}

type BookHandler struct {
	BookUseCase usecase.BookUseCaseInterface
}

func NewBookHandler(bookUseCase usecase.BookUseCaseInterface) usecase.BookUseCaseInterface {
	return BookHandler{
		BookUseCase: bookUseCase,
	}
}

func (h BookHandler) Save(book domain.Book) error {
	err := h.BookUseCase.Save(book)

	if err != nil {
		return err
	}

	return nil
}