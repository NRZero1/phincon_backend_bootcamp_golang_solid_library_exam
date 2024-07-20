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