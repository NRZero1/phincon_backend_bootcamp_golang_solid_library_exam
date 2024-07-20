package repository

import (
	"errors"
	"fmt"
	"solid_library_exam/internal/domain"
)

type UserRepository struct {
	users map[int]domain.User
}

func NewUserRepository() UserRepositoryInterface {
	return UserRepository{
		users: map[int]domain.User{},
	}
}

func (repo UserRepository) Save(user *domain.User) (domain.User, error) {
	if _, exists := repo.users[user.ID]; exists {
		return domain.User{}, errors.New("user already exists")
	}
	repo.users[user.ID] = *user
	fmt.Println("Saved data: ", repo.users[user.ID])
	return repo.users[user.ID], nil
}

func (repo UserRepository) FindByUsername(username string) (domain.User, error) {
	for k, v := range repo.users {
		if v.Username == username {
			return repo.users[k], nil
		}
	}

	return domain.User{}, fmt.Errorf("user with username %s doesn't exist", username)
}

func (repo UserRepository) FindById(id int) (domain.User, error) {
	if foundUser, exists := repo.users[id]; exists {
		return foundUser, nil
	}

	return domain.User{}, fmt.Errorf("no user found with Id %d", id)
}

func (repo UserRepository) BorrowBook(idUser int, book domain.Book) domain.User {
    foundUser, _ := repo.FindById(idUser)
    foundUser.ListOfBorrowedBooks = append(foundUser.ListOfBorrowedBooks, book)
    
    repo.users[idUser] = foundUser
    
    return repo.users[idUser]
}

func (repo UserRepository) GetAll() ([]domain.User) {
	listOfUsers := make([]domain.User, 0, len(repo.users))

	for _, v := range repo.users {
		listOfUsers = append(listOfUsers, v)
	}

	return listOfUsers
}

func (repo UserRepository) ReturnBook(userId int, bookId int) (domain.User, error) {
	user, err := repo.FindById(userId)

	if err != nil {
		return domain.User{}, err
	}

	var updatedList []domain.Book
	bookFound := false
	for _, b := range user.ListOfBorrowedBooks {
		if b.ID == bookId {
			bookFound = true
			continue
		}
		updatedList = append(updatedList, b)
	}

	if !bookFound {
		return domain.User{}, errors.New("book not found in user's borrowed list")
	}

	user.ListOfBorrowedBooks = updatedList

	return user, nil
}