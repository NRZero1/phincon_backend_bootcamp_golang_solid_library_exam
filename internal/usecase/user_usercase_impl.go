package usecase

import (
	"solid_library_exam/internal/domain"
	"solid_library_exam/internal/repository"
)

type UserUseCase struct {
	repo repository.UserRepositoryInterface
}

func NewUserUseCase(repo repository.UserRepositoryInterface) UserUseCaseInterface {
	return UserUseCase{
		repo: repo,
	}
}

func (uc UserUseCase) Save(user domain.User) (domain.User, error) {
	savedUser, err := uc.repo.Save(&user)

	if err != nil {
		return domain.User{}, err
	}

	return savedUser, nil
}

func (uc UserUseCase) FindByUsername(username string) (domain.User, error) {
	foundUser, err := uc.repo.FindByUsername(username)

	if err != nil {
		return domain.User{}, err
	}

	return foundUser, nil
}

func (uc UserUseCase) FindById(id int) (domain.User, error) {
	foundUser, err := uc.repo.FindById(id)

	if err != nil {
		return domain.User{}, err
	}

	return foundUser, nil
}

func (uc UserUseCase) BorrowBook(idUser int, book domain.Book) (domain.User) {
	updatedUser := uc.repo.BorrowBook(idUser, book)
	return updatedUser
}

func (uc UserUseCase) GetAll() ([]domain.User) {
	return uc.repo.GetAll()
}

func (uc UserUseCase) ReturnBook(userId int, bookId int) (domain.User, error) {
	updatedUser, err := uc.repo.ReturnBook(userId, bookId)
	
	if err != nil {
		return domain.User{}, err
	}

	return updatedUser, nil
}