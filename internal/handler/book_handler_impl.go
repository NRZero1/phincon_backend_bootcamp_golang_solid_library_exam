package handler

import (
	"solid_library_exam/internal/domain"
	"solid_library_exam/internal/usecase"
)

type BookHandler struct {
	bookUseCase usecase.BookUseCaseInterface
}

func NewBookHandler(bookUseCase usecase.BookUseCaseInterface) BookHandlerInterface {
	return BookHandler{
		bookUseCase: bookUseCase,
	}
}

func (h BookHandler) Save(book domain.Book) (domain.Book, error) {
	savedBook, err := h.bookUseCase.Save(book)

	if err != nil {
		return domain.Book{}, err
	}

	return savedBook, nil
}

func (h BookHandler) SetUserId(idBook int, idUser int) (domain.Book, error) {
	book, err := h.bookUseCase.SetUserId(idBook, idUser)

	if err != nil {
		return domain.Book{}, nil
	}

	return book, nil
}

func (h BookHandler) FindById(id int) (domain.Book, error) {
	foundBook, err := h.bookUseCase.FindById(id)

	if err != nil {
		return domain.Book{}, err
	}

	return foundBook, nil
}

func (h BookHandler) FindByTitle(title string) ([]domain.Book, error) {
	foundBook, err := h.bookUseCase.FindByTitle(title)

	if err != nil {
		return nil, err
	}

	return foundBook, nil
}