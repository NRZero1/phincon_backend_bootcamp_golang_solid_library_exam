package handler

import (
	"solid_library_exam/internal/domain"
	"solid_library_exam/internal/usecase"
)

type BookHandler struct {
	BookSave usecase.BookSave
	BookFindByTitle usecase.BookFindByTitle
	BookFindById usecase.BookFindById
}

func NewBookHandler(bookUseCaseSave usecase.BookSave, bookUseCaseFindById usecase.BookFindById, bookUseCaseFindByTitle usecase.BookFindByTitle) BookHandlerInterface {
	return BookHandler{
		BookSave: bookUseCaseSave,
		BookFindByTitle: bookUseCaseFindByTitle,
		BookFindById: bookUseCaseFindById,
	}
}

func (h BookHandler) Save(book domain.Book) (domain.Book, error) {
	savedBook, err := h.BookSave.Save(book)

	if err != nil {
		return domain.Book{}, err
	}

	return savedBook, nil
}

func (h BookHandler) FindById(id int) (domain.Book, error) {
	foundBook, err := h.BookFindById.FindById(id)

	if err != nil {
		return domain.Book{}, err
	}

	return foundBook, nil
}

func (h BookHandler) FindByTitle(title string) (domain.Book, error) {
	foundBook, err := h.BookFindByTitle.FindByTitle(title)

	if err != nil {
		return domain.Book{}, err
	}

	return foundBook, nil
}