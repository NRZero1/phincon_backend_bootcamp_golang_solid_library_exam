package repository

import "solid_library_exam/internal/domain"

type BookRepositoryInterface interface {
	BookSave
	BookSetUserId
	BookFindById
	BookFindByTitle
}

type BookSave interface {
	Save(*domain.Book) (domain.Book, error)
}

type BookFindById interface {
	FindById(int) (domain.Book, error)
}

type BookFindByTitle interface {
	FindByTitle(string) ([]domain.Book, error)
}

type BookSetUserId interface {
	SetUserId(int, int) (domain.Book, error)
}