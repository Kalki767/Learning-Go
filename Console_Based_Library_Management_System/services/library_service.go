package services

import (
	"errors"
	"Console_Based_Library_Management_System/models"
)

type Library struct {
	books map[int]models.Book
	members map[int]models.Member
}

func NewLibrary() *Library {
	return &Library{
		books: make(map[int]models.Book),
		members: make(map[int]models.Member),
	}
}

type LibraryManager interface {
	AddBook(book models.Book) error
	RemoveBook(bookID int) error
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book 
	ListBorrowedBooks(memberID int) []models.Book
}

func (l *Library) AddBook(book models.Book) error {
	_, ok := l.books[book.Id]
	if ok {
		return errors.New("book already exists")
	}
	l.books[book.Id] = book
	return nil
}

func (l *Library) RemoveBook(bookID int) error {
	_, ok := l.books[bookID]
	if ok {
		delete(l.books, bookID)
		return nil
	}
	return errors.New("book not found")
}

func (l *Library) BorrowBook(bookID int, memberID int) error {
	book, ok := l.books[bookID]
	if !ok {
		return errors.New("book not found")
	}
	if book.Status == "Borrowed" {
		return errors.New("book already borrowed")
	}
	book.Status = "Borrowed"
	l.books[bookID] = book
	return nil
}

func (l *Library) ReturnBook(bookID int, memberID int) error {
	book, ok := l.books[bookID]
	if !ok {
		return errors.New("book not found")
	}
	if book.Status == "Available" {
		return errors.New("book already available")
	}
	book.Status = "Available"
	l.books[bookID] = book
	return nil
}

func (l *Library) ListAvailableBooks() []models.Book {
	var books []models.Book
	for _, book := range l.books {
		if book.Status == "Available" {
			books = append(books, book)
		}
	}
	return books
}

func (l *Library) ListBorrowedBooks(memberID int) []models.Book {
	var books []models.Book
	for _, book := range l.books {
		if book.Status == "Borrowed" {
			books = append(books, book)
		}
	}
	return books
}

