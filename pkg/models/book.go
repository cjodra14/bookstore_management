package models

import (
	"github.com/cjodra14/bookstore_management/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func dbInit() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (book *Book) CreateBook() *Book {
	db.NewRecord(book)
	db.Create(&book)
	return book
}

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBookByID(bookID int64) (*Book, *gorm.DB) {
	var getBook Book
	dataBase := db.Where("ID=?", bookID).Find(&getBook)
	return &getBook, dataBase
}

func DeleteBook(bookID int64) Book {
	var book Book
	db.Where("ID=?", bookID).Delete(book)
	return book
}
