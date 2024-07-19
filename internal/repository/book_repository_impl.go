package repository

import (
	"fmt"
	"solid_library_exam/internal/domain"
)

type BookRepository struct {
	books map[int]domain.Book
}

func NewBookRepository() BookRepositoryInterface {
	return BookRepository {
		books: map[int]domain.Book{},
	}
}

func (repo BookRepository) Save(book *domain.Book) (domain.Book, error) {
	if _, exists := repo.books[book.ID]; exists {
		return domain.Book{}, fmt.Errorf("book with ID %d already exist", book.ID)
	}
	repo.books[book.ID] = *book
	return repo.books[book.ID], nil
}

func (repo BookRepository) UpdateStatus(id int) (domain.Book, error) {
	if foundBook, exists := repo.books[id]; exists {
		foundBook.Status = !foundBook.Status
		return foundBook, nil
	}
	return domain.Book{}, nil
}