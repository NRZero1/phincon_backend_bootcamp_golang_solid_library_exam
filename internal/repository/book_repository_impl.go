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
	foundBook, err := repo.FindById(id)

	if err != nil {
		return domain.Book{}, err
	}

	foundBook.Status = !foundBook.Status
	return foundBook, nil
}

func (repo BookRepository) FindById(id int) (domain.Book, error) {
	if foundBook, exist := repo.books[id]; exist {
		return foundBook, nil
	}

	return domain.Book{}, fmt.Errorf("no book with ID %d exist", id)
}

func (repo BookRepository) FindByTitle(title string) (domain.Book, error) {
	for k, v := range repo.books {
		if v.Title == title {
			return repo.books[k], nil
		}
	}

	return domain.Book{}, fmt.Errorf("no book with title '%s' exist(s)", title)
}