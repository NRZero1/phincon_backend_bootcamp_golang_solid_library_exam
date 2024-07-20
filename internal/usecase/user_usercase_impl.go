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