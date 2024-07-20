package main

import (
	"bufio"
	"fmt"
	"os"
	"solid_library_exam/internal/domain"
	"solid_library_exam/internal/handler"
	"solid_library_exam/internal/repository"
	"solid_library_exam/internal/usecase"
	"strconv"
)

var hb handler.BookHandlerInterface
var hu handler.UserHandlerInterface

func init() {
	hb = initBookHandler()
	hu = initUserHandler()

	books := []domain.Book{
		{ID: 1, Title: "Abc", Author: "Bene", UserId: 0},
		{ID: 2, Title: "Hahaha", Author: "Niko", UserId: 0},
		{ID: 3, Title: "Hihihi", Author: "Raymond", UserId: 0},
		{ID: 4, Title: "Wkwkwk", Author: "Daniel", UserId: 0},
		{ID: 5, Title: "asdfghj", Author: "Ahmad", UserId: 0},
		{ID: 6, Title: "zxcvbn", Author: "Alexander", UserId: 0},
		{ID: 7, Title: "qwertyui", Author: "Nisa", UserId: 0},
		{ID: 8, Title: "12345678", Author: "Saputra", UserId: 0},
		{ID: 9, Title: "poiuytr", Author: "Nadya", UserId: 0},
		{ID: 10, Title: "lkjhgf", Author: "RRC", UserId: 0},
	}

	for _, v := range books {
		hb.Save(v)
	}

	printAllBooks()
}

func printAllBooks() {
	for _, v := range hb.GetAll() {
		var username string
		fmt.Println()
		fmt.Println("Book Id: ", v.ID)
		fmt.Println("Book Title: ", v.Title)
		fmt.Println("Book Author: ", v.Author)
		if v.UserId == 0 {
			username = "null"
		} else {
			user, _ := hu.FindById(v.UserId)
			username = user.Username
		}
		fmt.Println("Loaner: ", username)
	}
	fmt.Println()
}

func printListOfAvailableBooks() {
	for _, v := range hb.GetAll() {
		if v.UserId != 0 {
			continue
		}
		fmt.Println()
		fmt.Println("Book Id: ", v.ID)
		fmt.Println("Book Title: ", v.Title)
		fmt.Println("Book Author: ", v.Author)
		fmt.Println("Loaner: null")
	}
}

func printAllLoanedBooks() {
	for _, v := range hb.GetAll() {
		if v.UserId != 0 {
			var username string
			fmt.Println()
			fmt.Println("Book Id: ", v.ID)
			fmt.Println("Book Title: ", v.Title)
			fmt.Println("Book Author: ", v.Author)
			user, _ := hu.FindById(v.UserId)
			username = user.Username
			fmt.Println("Loaner: ", username)
		}
	}
}

func printAllUser() {
	for _, v := range hu.GetAll() {
		fmt.Println("User Id: ", v.ID)
		fmt.Println("User Loan Books: ", v.ListOfBorrowedBooks)
	}
}

func printListOfLoanedBooks(idUser int) {
	foundUser, _ := hu.FindById(idUser)

	fmt.Println("User Loaned Books: ")
	for _, v := range foundUser.ListOfBorrowedBooks {
		fmt.Println("Book Id: ", v.ID)
		fmt.Println("Book Title: ", v.Title)
		fmt.Println("Book Author: ", v.Author)
	}
}

func initBookHandler() handler.BookHandlerInterface {
	repo := repository.NewBookRepository()
	uc := usecase.NewBookUseCase(repo)
	h := handler.NewBookHandler(uc)
	return h
}

func initUserHandler() handler.UserHandlerInterface {
	repo := repository.NewUserRepository()
	uc := usecase.NewUserUseCase(repo)
	h := handler.NewUserHandler(uc)
	return h
}

func main() {
	userID := 1
	Scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Program Perpustakaan")
		fmt.Println("=========================")
		fmt.Println("Menu")
		fmt.Println("1. Pinjam Buku")
		fmt.Println("2. Pengembalian Buku")
		fmt.Println("3. Cek list buku yang dipinjam")
		fmt.Print("\nInput menu: ")
		Scanner.Scan()
		receivedMenu := Scanner.Text()
		fmt.Println(receivedMenu)

		menuInput, err := strconv.Atoi(receivedMenu)

		if err != nil {
			fmt.Println("Invald input, please input a number")
			continue
		}

		switch menuInput {
		case 1:
			fmt.Print("Please enter username: ")
			Scanner.Scan()
			receivedUsername := Scanner.Text()

			foundUser, err := hu.FindByUsername(receivedUsername)

			if err != nil {
				for {
					fmt.Print("No username detected, create account? (y/n) ")
					Scanner.Scan()
					createAccountInput := Scanner.Text()

					if createAccountInput != "y" && createAccountInput != "n" {
						fmt.Println("Invalid input, please enter only y or n")
						continue
					}
					foundUser, _ = hu.Save(receivedUsername, userID)
					userID++
					break
				}
			}

			for {
				printListOfAvailableBooks()
				fmt.Print("Enter book Id to borrow: ")
				Scanner.Scan()
				bookIdInput := Scanner.Text()

				bookId, errBookId := strconv.Atoi(bookIdInput)

				if errBookId != nil {
					fmt.Println("Invalid input, please enter a number")
					continue
				}

				fmt.Println(bookId)

				foundBook, errFoundBook := hb.FindById(bookId)

				if errFoundBook != nil {
					fmt.Printf("No book with ID %d found\n", bookId)
					continue
				}

				if foundBook.UserId != 0 {
					loanerUser, _ := hu.FindById(foundBook.UserId)
					fmt.Printf("Sorry, Book with Id %d and Title %s still being loaned by %s\n", foundBook.ID, foundBook.Title, loanerUser.Username)
					break
				}
				hu.BorrowBook(foundUser.ID, foundBook)
				hb.SetUserId(foundBook.ID, foundUser.ID)
				printAllUser()
				printAllBooks()
				break
			}
		case 2:
			for {
				fmt.Print("Please enter username: ")
				Scanner.Scan()
				receivedUsername := Scanner.Text()

				result, err := hu.FindByUsername(receivedUsername)

				if err != nil {
					fmt.Println("Invalid username, please try again")
					continue
				}
				printListOfLoanedBooks(result.ID)

				fmt.Println()
			
				for {
					fmt.Printf("Select Book Id you want to return: ")
					Scanner.Scan()
					receivedId := Scanner.Text()

					bookId, errBookId := strconv.Atoi(receivedId)

					if errBookId != nil {
						fmt.Println("Invalid Input, please input a number")
						continue
					}

					_, err = hu.ReturnBook(result.ID, bookId)
					if err != nil {
						fmt.Println("Error returning book:", err)
						continue
					}

					// Reset the book's UserId
					_, err = hb.SetUserId(bookId, 0)
					if err != nil {
						fmt.Println("Error updating book record:", err)
					}
					break
				}
				break
			}

		case 3:
			printAllLoanedBooks()
		default:
			fmt.Println("Invalid menu inputted, please enter only a number 1 or 2")
			fmt.Println()
		}
	}
}
