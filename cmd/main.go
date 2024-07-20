package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	
	for {
		fmt.Println("Program Perpustakaan")
		fmt.Println("=========================")
		fmt.Println("Menu")
		fmt.Println("1. Pinjam Buku")
		fmt.Println("2. Pengembalian Buku")
		fmt.Println("3. Cek list buku yang dipinjam")
		fmt.Print("\nInput menu: ")
		receivedMenu, _ := reader.ReadString('\n')
		
		if reflect.ValueOf(receivedMenu).Kind() != reflect.Int {
			fmt.Println("Invalid input, please input a number")
			continue
		}

		menuInput, _ := strconv.Atoi(receivedMenu)

		switch menuInput {
		case 1:
			fmt.Print("Please enter username: ")
			receivedUsername, _ := reader.ReadString('\n')
			
			if reflect.ValueOf(receivedUsername).Kind() != reflect.String {
				fmt.Println("Invalid input, please input a string")
				continue
			}

			// call user use case to check if username inputted exist
			// if doesn't exist create a new user

			// call store to borrowed books struct and add the borrowed book to slice in user struct
		case 2:
			fmt.Print("Please enter username: ")
			receivedUsername, _ := reader.ReadString('\n')
			
			if reflect.ValueOf(receivedUsername).Kind() != reflect.String {
				fmt.Println("Invalid input, please input a string")
				continue
			}
			// call user use case to check if username inputted exist
			// if doesn't exist abort the command return book
			
			// call return the borrowed books struct and remove the borrowed book from slice in user struct
		case 3:
			// fetch map of borrowed books -> get the username -> get the book title
		default:
			fmt.Println("Invalid menu inputted, please enter only a number 1 or 2")
			fmt.Println()
		}
	}
}