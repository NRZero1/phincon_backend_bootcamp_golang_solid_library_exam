package domain

type User struct {
	ID                  int
	Username            string
	Name                string
	ListOfBorrowedBooks []Book
}