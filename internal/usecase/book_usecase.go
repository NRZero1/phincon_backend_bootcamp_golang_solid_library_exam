package usecase

import (
	"solid_library_exam/internal/domain"
	"solid_library_exam/internal/repository"
)

type BookUseCase struct {
	repository repository.BookRepository
}

type BookUseCaseInterface interface {
	AddBook
}

type AddBook interface {
	AddBook(domain.Book) error
}

type BookUseCaseImpl struct {
	bookRepo repository.BookRepositoryInterface
}

func NewBookUseCase(bookRepo repository.BookRepositoryInterface) BookUseCaseInterface {
	return BookUseCaseImpl{
		bookRepo: bookRepo,
	}
}

func (uc BookUseCaseImpl) AddBook(book domain.Book) error {
	err := uc.bookRepo.Save(&book)

	if err != nil {
		return err
	}

	return nil
}