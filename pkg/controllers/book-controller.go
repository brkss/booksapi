package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/brkss/booksapi/pkg/models"
	"github.com/brkss/booksapi/pkg/utils"
	"github.com/gorilla/mux"
)

var book models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	book := models.GetAllBooks()
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing book's id ! ")
	}
	book, _ := models.GetBookById(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {

	book := &models.Book{}
	utils.ParseBody(r, book)
	b := book.CreateBook()

	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error accured whle parsing")
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}
	utils.ParseBody(r, book)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error accured while parsing !")
	}
	b, db := models.GetBookById(ID)
	if book.Name != "" {
		b.Name = book.Name
	}
	if book.Author != "" {
		b.Author = book.Author
	}
	if book.Publication != "" {
		b.Publication = book.Publication
	}
	db.Save(&b)
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/jsonm")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
