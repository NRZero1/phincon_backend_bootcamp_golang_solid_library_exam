package usecase

import (
	"solid_library_exam/internal/domain"
	"solid_library_exam/internal/repository"
)

type BookUseCase struct {
	repo repository.BookRepositoryInterface
}

func NewBookUseCase(bookRepo repository.BookRepositoryInterface) BookUseCaseInterface {
	return BookUseCase{
		repo: bookRepo,
	}
}

func (uc BookUseCase) Save(book domain.Book) (domain.Book, error) {
	savedBook, err := uc.repo.Save(&book)

	if err != nil {
		return domain.Book{}, err
	}

	return savedBook, nil
}

func (uc BookUseCase) UpdateStatus(id int) (domain.Book, error) {
	book, err := uc.repo.UpdateStatus(id)

	if err != nil {
		return domain.Book{}, nil
	}

	return book, nil
}