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

func (repo BookRepository) SetUserId(idBook int, idUser int) (domain.Book, error) {
	foundBook, err := repo.FindById(idBook)

	if err != nil {
		return domain.Book{}, err
	}

	foundBook.UserId = idUser
	repo.books[idBook] = foundBook
	return foundBook, nil
}

func (repo BookRepository) FindById(id int) (domain.Book, error) {
	if foundBook, exist := repo.books[id]; exist {
		return foundBook, nil
	}

	return domain.Book{}, fmt.Errorf("no book with ID %d exist", id)
}

func (repo BookRepository) FindByTitle(title string) ([]domain.Book, error) {
	var foundBooks []domain.Book
	for _, v := range repo.books {
		if v.Title == title {
			foundBooks = append(foundBooks, v)
		}
	}

	if len(foundBooks) == 0 {
		return nil, fmt.Errorf("no books with title '%s' exist", title)
	}

	return foundBooks, nil
}

func (repo BookRepository) GetAll() ([]domain.Book) {
	listOfBooks := make([]domain.Book, 0, len(repo.books))

	for _, v := range repo.books {
		listOfBooks = append(listOfBooks, v)
	}

	return listOfBooks
}