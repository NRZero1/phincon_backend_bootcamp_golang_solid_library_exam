package handler

import (
	"solid_library_exam/internal/domain"
	"solid_library_exam/internal/usecase"
)

type BookHandler struct {
	bookUseCase usecase.BookUseCaseInterface
}


// Implement D in Solid, using high level abstraction so it will not depend on specific Book Use Case
// All of the function below doesn't need to change anymore but open if book handler want to extend more function (implementing O)
// by using interface as a type it will also implementing L in solid because interface is like its parent class, so replacing book handler interface with its child book handler won't break the app
func NewBookHandler(bookUseCase usecase.BookUseCaseInterface) BookHandlerInterface {
	return BookHandler{
		bookUseCase: bookUseCase,
	}
}

// each of this function have single responsibility implementing S
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

func (h BookHandler) GetAll() ([]domain.Book) {
	return h.bookUseCase.GetAll()
}