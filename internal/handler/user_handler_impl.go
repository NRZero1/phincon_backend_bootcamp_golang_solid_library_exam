package handler

import (
	"solid_library_exam/internal/domain"
	"solid_library_exam/internal/usecase"
)

type UserHandler struct {
	UserUseCase usecase.UserUseCaseInterface
}

func NewUserHandler(userUseCase usecase.UserUseCaseInterface) UserHandlerInterface {
	return UserHandler{
		UserUseCase: userUseCase,
	}
}

func (h UserHandler) Save(user domain.User) (domain.User, error) {
	savedUser, err := h.UserUseCase.Save(user)

	if err != nil {
		return domain.User{}, err
	}

	return savedUser, nil
}

func (h UserHandler) FindByUsername(username string) (domain.User, error) {
	foundUser, err := h.UserUseCase.FindByUsername(username)

	if err != nil {
		return domain.User{}, err
	}

	return foundUser, nil
}