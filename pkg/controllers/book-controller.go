package controllers

import (
	"encoding/json"

	"log"
	"net/http"
	"strconv"

	"github.com/cjodra14/bookstore_management/pkg/models"
	"github.com/cjodra14/bookstore_management/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBook(writer http.ResponseWriter, request *http.Request) {
	newBooks := models.GetAllBooks()
	res, err := json.Marshal(newBooks)
	if err != nil {
		log.Fatal(err)
	}

	writer.Header().Set("Content-type", "pkglicaton/json")
	writer.WriteHeader(http.StatusOK)
	_, err = writer.Write(res)
	if err != nil {
		log.Fatal(err)
	}
}

func GetBookByID(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	bookID := vars["bookID"]
	ID, err := strconv.ParseInt(bookID, 0, 0)

	if err != nil {
		log.Println("error while parsing")
	}

	bookDetails, _ := models.GetBookByID(ID)

	res, _ := json.Marshal(bookDetails)
	writer.Header().Set("Content-type", "pkglicaton/json")
	writer.WriteHeader(http.StatusOK)
	_, err = writer.Write(res)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateBook(writer http.ResponseWriter, request *http.Request) {
	createBook := &models.Book{}
	utils.ParseBody(request, createBook)
	book := createBook.CreateBook()

	res, err := json.Marshal(book)
	if err != nil {
		log.Fatal(err)
	}

	writer.WriteHeader(http.StatusOK)
	_, err = writer.Write(res)
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteBook(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	bookID := vars["bookID"]

	ID, err := strconv.ParseInt(bookID, 0, 0)
	if err != nil {
		log.Fatal("error while parsing")
	}

	book := models.DeleteBook(ID)

	res, err := json.Marshal(book)
	if err != nil {
		log.Fatal(err)
	}

	writer.Header().Set("Content-type", "pkglicaton/json")
	writer.WriteHeader(http.StatusOK)

	_, err = writer.Write(res)
	if err != nil {
		log.Fatal(err)
	}
}

func UpdateBook(writer http.ResponseWriter, request *http.Request) {
	updateBook := &models.Book{}
	utils.ParseBody(request, updateBook)
	vars := mux.Vars(request)
	bookID := vars["bookID"]

	ID, err := strconv.ParseInt(bookID, 0, 0)
	if err != nil {
		log.Fatal("error while parsing")
	}

	bookDetails, db := models.GetBookByID(ID)

	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}

	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}

	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}

	db.Save(&bookDetails)

	res, err := json.Marshal(bookDetails)
	if err != nil {
		log.Fatal(err)
	}

	writer.Header().Set("Content-type", "pkglicaton/json")
	writer.WriteHeader(http.StatusOK)

	_, err = writer.Write(res)
	if err != nil {
		log.Fatal(err)
	}
}
