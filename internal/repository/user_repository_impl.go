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