package usecase

import (
	"solid_library_exam/internal/domain"
	"solid_library_exam/internal/repository"
)

type BookUseCase struct {
	repository repository.BookRepository
}

type BookUseCaseInterface interface {
	Save
}

type Save interface {
	Save(domain.Book) error
}

type BookUseCaseImpl struct {
	bookRepo repository.BookRepositoryInterface
}

func NewBookUseCase(bookRepo repository.BookRepositoryInterface) BookUseCaseInterface {
	return BookUseCaseImpl{
		bookRepo: bookRepo,
	}
}

func (uc BookUseCaseImpl) Save(book domain.Book) error {
	err := uc.bookRepo.Save(&book)

	if err != nil {
		return err
	}

	return nil
}