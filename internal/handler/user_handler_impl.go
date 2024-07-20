package handler

import (
	"bufio"
	"fmt"
	"os"
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

func (h UserHandler) Save(username string, userId int) (domain.User, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Please enter your name: ")
	inputtedName, _ := reader.ReadString('\n')

	user := domain.User{
		ID: userId,
		Username: username,
		Name: inputtedName,
	}
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

func (h UserHandler) FindById(id int) (domain.User, error) {
	foundUser, err := h.UserUseCase.FindById(id)

	if err != nil {
		return domain.User{}, err
	}

	return foundUser, nil
}

func (h UserHandler) BorrowBook(idUser int, book domain.Book) (domain.User) {
	updatedUser := h.UserUseCase.BorrowBook(idUser, book)
	return updatedUser
}

func (h UserHandler) GetAll() ([]domain.User) {
	return h.UserUseCase.GetAll()
}

func (h UserHandler) ReturnBook(userID int, bookID int) (domain.User, error) {
	return h.UserUseCase.ReturnBook(userID, bookID)
}