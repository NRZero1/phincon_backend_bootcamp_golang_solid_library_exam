package domain

type User struct {
	ID                 int
	Username           string
	Name               string
	MapOfBorrowedBooks map[int]Book
}