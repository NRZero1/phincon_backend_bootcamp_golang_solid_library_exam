package repository

import (
	"errors"
	"fmt"
	"solid_library_exam/internal/domain"
)

type BookRepositoryInterface interface {
	BookSave
}

type BookSave interface {
	Save(*domain.Book) error
}

type BookRepository struct {
	books map[int]domain.Book
}

func NewBookRepository() BookRepositoryInterface {
	return BookRepository {
		books: map[int]domain.Book{},
	}
}

func (repo BookRepository) Save(book *domain.Book) error {
	if _, exists := repo.books[book.ID]; exists {
		return errors.New(fmt.Sprintf("Book with ID %s already exist", book.ID))
	}
	repo.books[book.ID] = *book
	fmt.Println("Saved data: ", repo.books)
	return nil
}