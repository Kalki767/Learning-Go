package controllers

import (
	"fmt"
	"Console_Based_Library_Management_System/services"
	"Console_Based_Library_Management_System/models"
	"strings"
	"os"
	"bufio"
	"strconv"
)

func stringInput() string{
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	return strings.TrimSpace(input)
}
func DisplayInterface(){
	fmt.Println("Welcome to the Library Management System")
	fmt.Println("Please select an option")
	fmt.Println("1. Add Book")
	fmt.Println("2. Remove Book")
	fmt.Println("3. Borrow Book")
	fmt.Println("4. Return Book")
	fmt.Println("5. List Available Books")
	fmt.Println("6. List Borrowed Books")
	optionstr := stringInput()
	option, err := strconv.Atoi(optionstr)
	for err != nil{
		fmt.Println("Please enter a valid option")
		optionstr = stringInput()
		option, err = strconv.Atoi(optionstr)
	}
	switch option {
	case 1:
		AddBook()
	case 2:
		RemoveBook()
	case 3:
		BorrowBook()
	case 4:
		ReturnBook()
	case 5:
		ListAvailableBooks()
	case 6:
		ListBorrowedBooks()
	default:
		fmt.Println("Invalid option. Please Enter a valid option")
		DisplayInterface()
}
}
var library = services.NewLibrary()
func AddBook() {
	fmt.Println("Enter the book title")
	title := stringInput()
	fmt.Println("Enter the author name")
	author := stringInput()
	fmt.Println("Enter the book id")
	idStr := stringInput()
	id, err := strconv.Atoi(idStr)
	for err != nil{
		fmt.Println("Please enter a valid id")
		idStr = stringInput()
		id, err = strconv.Atoi(idStr)
	}
	fmt.Println("Enter the book status : 1 for Available and 2 for Borrowed")
	choice := stringInput()
	status := ""
	for choice != "1" && choice != "2"{
		fmt.Println("Please enter a valid choice")
		choice = stringInput()
	}
	if choice == "1"{
		status = "Available"
	}else if choice == "2"{
		status = "Borrowed"
	}
	book := models.Book{
		Title: title,
		Author: author,
		Id: id,
		Status: status,
	}
	err = library.AddBook(book)
	if err == nil {
		fmt.Println("Book added successfully")
	} else {
		fmt.Println(err)
	}
}

func RemoveBook(){
	fmt.Println("Enter the book id")
	idStr := stringInput()
	id, err := strconv.Atoi(idStr)
	for err != nil{
		fmt.Println("Please enter a valid id")
		idStr = stringInput()
		id, err = strconv.Atoi(idStr)
	}
	err = library.RemoveBook(id)
	if err == nil {
		fmt.Println("Book removed successfully")
	} else {
		fmt.Println(err)
	}
}

func BorrowBook(){
	fmt.Println("Enter the book id")
	idStr := stringInput()
	id, err := strconv.Atoi(idStr)
	for err != nil{
		fmt.Println("Please enter a valid book id")
		idStr = stringInput()
		id, err = strconv.Atoi(idStr)
	}
	fmt.Println("Enter the member id")
	memberIDStr := stringInput()
	memberID, err := strconv.Atoi(memberIDStr)
	for err != nil{
		fmt.Println("Please enter a valid member id")
		memberIDStr = stringInput()
		memberID, err = strconv.Atoi(memberIDStr)
	}
	err = library.BorrowBook(id, memberID)
	if err == nil {
		fmt.Println("Book borrowed successfully")
	} else {
		fmt.Println(err)
	}
}

func ReturnBook(){
	fmt.Println("Enter the book id")
	idStr := stringInput()
	id, err := strconv.Atoi(idStr)
	for err != nil{
		fmt.Println("Please enter a valid book id")
		idStr = stringInput()
		id, err = strconv.Atoi(idStr)
	}
	fmt.Println("Enter the member id")
	memberIDStr := stringInput()
	memberID, err := strconv.Atoi(memberIDStr)
	for err != nil{
		fmt.Println("Please enter a valid member id")
		memberIDStr = stringInput()
		memberID, err = strconv.Atoi(memberIDStr)
	}
	err = library.ReturnBook(id, memberID)
	if err == nil {
		fmt.Println("Book returned successfully")
	} else {
		fmt.Println(err)
	}
}

func ListAvailableBooks(){
	books := library.ListAvailableBooks()
	if len(books) == 0 {
		fmt.Println("No books available")
	} else {
		for _, book := range books {
			fmt.Println(book)
		}
	}
}

func ListBorrowedBooks(){
	fmt.Println("Enter the member id")
	memberIDStr := stringInput()
	memberID, err := strconv.Atoi(memberIDStr)
	for err != nil{
		fmt.Println("Please enter a valid member id")
		memberIDStr = stringInput()
		memberID, err = strconv.Atoi(memberIDStr)
	}
	books := library.ListBorrowedBooks(memberID)
	if len(books) == 0 {
		fmt.Println("No books borrowed")
	} else {
		for _, book := range books {
			fmt.Println(book)
		}
	}
}